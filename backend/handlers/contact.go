package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"app/core"
	"app/db"
)

// The public contact form posts a name, an address and a message; we store it and mail it to the
// inbox configured in CONTACT_EMAIL. The endpoint is unauthenticated, so the per-IP window in
// db.CheckAndCount is the only brake there is — there is no token to charge, no account to
// suspend, and the caller is a stranger by definition.

// Field ceilings, counted in runes so an accented Spanish message is not cut short of what the
// form's own limit implies. The row is built from an unauthenticated body, so the sizes are
// enforced here rather than trusted from the browser.
const (
	contactMaxNameRunes    = 120
	contactMaxEmailRunes   = 160
	contactMaxCompanyRunes = 160
	contactMaxMessageRunes = 4000
	// Below this a submission carries nothing to answer, and it is the length a bot filling every
	// field with one character trips over.
	contactMinMessageRunes = 10
	contactMinNameRunes    = 2
)

// Deliberately loose. A stricter pattern rejects addresses that are legal and deliverable, and
// the only thing that actually proves an address works is mail arriving at it — which is not a
// check this form performs. This exists to catch a typo, not to validate RFC 5322.
var emailPattern = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]{2,}$`)

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// PostContactMessage stores one contact-form submission and notifies the configured inbox.
func PostContactMessage(req *core.HandlerArgs) core.HandlerResponse {
	if req.Body == nil {
		return req.MakeErr("La solicitud no trae cuerpo.")
	}

	body := struct {
		Name    string
		Email   string
		Company string
		Message string
	}{}
	if err := json.Unmarshal([]byte(*req.Body), &body); err != nil {
		return req.MakeErr("Error al deserializar el cuerpo de la solicitud.", err)
	}

	// Refused before anything is written: a message with no destination is only a row nobody will
	// ever read, and an operator who has not set this yet should find out from the first test
	// submission rather than from an empty inbox weeks later.
	if len(core.Env.CONTACT_EMAIL) == 0 {
		return req.MakeErrCode(
			"El formulario de contacto no está configurado: falta CONTACT_EMAIL.", 503)
	}

	name := strings.TrimSpace(body.Name)
	email := normalizeEmail(body.Email)
	company := strings.TrimSpace(body.Company)
	message := strings.TrimSpace(body.Message)

	if utf8.RuneCountInString(name) < contactMinNameRunes {
		return req.MakeErr("El nombre es necesario.")
	}
	if utf8.RuneCountInString(name) > contactMaxNameRunes {
		return req.MakeErr(fmt.Sprintf("El nombre no puede exceder %v caracteres.", contactMaxNameRunes))
	}
	if utf8.RuneCountInString(email) > contactMaxEmailRunes || !emailPattern.MatchString(email) {
		return req.MakeErr("El correo electrónico no posee un formato válido.")
	}
	if utf8.RuneCountInString(company) > contactMaxCompanyRunes {
		return req.MakeErr(fmt.Sprintf("La empresa no puede exceder %v caracteres.", contactMaxCompanyRunes))
	}
	if utf8.RuneCountInString(message) < contactMinMessageRunes {
		return req.MakeErr(fmt.Sprintf("El mensaje debe poseer al menos %v caracteres.", contactMinMessageRunes))
	}
	if utf8.RuneCountInString(message) > contactMaxMessageRunes {
		return req.MakeErr(fmt.Sprintf("El mensaje no puede exceder %v caracteres.", contactMaxMessageRunes))
	}

	ipKey, hasClientIP := req.ClientIPKey()
	if !hasClientIP {
		return req.MakeErrCode("No se pudo determinar el origen de la solicitud.", 400)
	}

	// The limit is spent here and not before the validation above, so a visitor who mistypes their
	// address does not burn part of their budget on a submission that was never going to be
	// stored. The trade is that a malformed body is free to send; bounding that is the job of the
	// throttle on the Function URL, not of a counter that costs a write to consult.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	window := time.Duration(core.Env.CONTACT_WINDOW_MINUTES) * time.Minute
	decision, err := db.CheckAndCount(
		ctx, core.ActionContactByIP, ipKey, core.Env.CONTACT_MAX_MESSAGES_PER_IP, window)

	if err != nil {
		// Fail closed. With the limiter unreachable this endpoint is an open relay into somebody's
		// inbox, and a contact form is on nobody's critical path: a visitor who has to come back in
		// a minute is a far smaller cost than an inbox nobody can use again.
		core.Log("PostContactMessage:: error, límite de tasa no disponible::", err)
		return req.MakeErrCode("El servicio no está disponible. Intente nuevamente.", 503)
	}
	if !decision.Allowed {
		core.Log("PostContactMessage:: límite por IP alcanzado", ipKey)
		response := req.MakeErrCode(fmt.Sprintf(
			"Se alcanzó el máximo de %d mensajes por %d minutos desde esta conexión.",
			core.Env.CONTACT_MAX_MESSAGES_PER_IP, core.Env.CONTACT_WINDOW_MINUTES), 429)
		response.Headers["Retry-After"] = fmt.Sprint(int(decision.RetryAfter.Seconds()) + 1)
		return response
	}

	now := time.Now()
	storedMessage := db.ContactMessage{
		WeekCode: core.WeekCodeAt(now),
		IP:       ipKey,
		Created:  now.Unix(),
		Sequence: decision.Hits,
		Name:     name,
		Email:    email,
		Company:  company,
		Message:  message,
		Status:   db.ContactStatusDelivered,
		Updated:  now.Unix(),
	}

	// Persist first, deliver second. The row IS the message, and losing it to an SMTP hiccup would
	// lose what the visitor wrote — they are not coming back to type it again. Writing first also
	// means a failed send still consumes the sender's budget, which is what stops a broken mail
	// server from turning this into an unmetered endpoint.
	if err := db.InsertContactMessage(ctx, &storedMessage); err != nil {
		return req.MakeErrCode("Error al registrar el mensaje de contacto.", 500, err)
	}

	subject := fmt.Sprintf("%v — nuevo mensaje de contacto", core.Env.APP_NAME)
	if err := core.SendEmail(core.Env.CONTACT_EMAIL, subject, makeContactEmailBody(&storedMessage)); err != nil {
		// The visitor is told the message arrived, because it did: it is stored, and status 2 is
		// what marks it as still needing to reach the inbox.
		core.Log("PostContactMessage:: no se pudo notificar el mensaje", storedMessage.Created, err)
		if err := db.MarkContactMessageUndelivered(ctx, &storedMessage); err != nil {
			core.Log("PostContactMessage:: no se pudo marcar el mensaje como no entregado", err)
		}
		return req.MakeResponse(map[string]any{"Received": true, "Notified": false})
	}

	core.Log("PostContactMessage:: mensaje recibido de", email)
	return req.MakeResponse(map[string]any{"Received": true, "Notified": true})
}

// makeContactEmailBody escapes every field: the whole body is attacker-controlled text, and it is
// read in a mail client that renders HTML.
func makeContactEmailBody(message *db.ContactMessage) string {
	companyLine := ""
	if len(message.Company) > 0 {
		companyLine = fmt.Sprintf(`<p><b>Empresa:</b> %v</p>`, html.EscapeString(message.Company))
	}

	return fmt.Sprintf(`<html>
	<head><meta http-equiv="Content-Type" content="text/html; charset=utf-8" /></head>
	<body style="font-family:system-ui,sans-serif;color:#1e293b;line-height:1.55">
		<h2 style="color:#4646ee">Nuevo mensaje de contacto</h2>
		<p><b>Nombre:</b> %v</p>
		<p><b>Correo:</b> <a href="mailto:%v">%v</a></p>
		%v
		<p><b>Mensaje:</b></p>
		<p style="white-space:pre-wrap;border-left:3px solid #c7d2fe;padding-left:12px">%v</p>
		<p style="color:#64748b;font-size:13px">Enviado desde el formulario de contacto de %v.</p>
	</body>
</html>`,
		html.EscapeString(message.Name),
		html.EscapeString(message.Email), html.EscapeString(message.Email),
		companyLine,
		html.EscapeString(message.Message),
		html.EscapeString(core.Env.APP_URL))
}
