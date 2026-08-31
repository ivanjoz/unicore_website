package core

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"mime"
	"net"
	"net/smtp"
	"strconv"
	"time"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	sestypes "github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

// ErrMailerNotConfigured is the absence of any way to send. It is a distinct error because the
// contact handler treats it as a delivery failure and not as a rejected request: the message is
// already stored by then, and losing what a visitor wrote because an operator has not verified an
// SES identity yet would be the worse outcome.
var ErrMailerNotConfigured = errors.New("el envío de correo no está configurado")

// SendEmail delivers one HTML message.
//
// Two backends, chosen by which one is configured. SES is preferred and needs no secret in the
// environment — the Lambda execution role carries the permission — which is why it is the default
// for a deployment that is AWS-only anyway. SMTP stays as the escape hatch for an operator who
// already has a mailbox somewhere and does not want to verify a domain with SES.
func SendEmail(toAddress string, subject string, htmlBody string) error {
	if len(toAddress) == 0 {
		return errors.New("no se indicó el destinatario del correo")
	}
	if len(Env.SES_FROM_EMAIL) > 0 {
		return sendEmailWithSES(toAddress, subject, htmlBody)
	}
	if len(Env.SMTP_HOST) > 0 {
		return sendEmailWithSMTP(toAddress, subject, htmlBody)
	}
	return ErrMailerNotConfigured
}

// The whole send has to finish well inside the Lambda timeout, and it runs on the request path
// with a visitor waiting, so a mail server that accepts the connection and then stalls must not
// be able to hold the invocation open.
const mailerTimeout = 8 * time.Second

func sendEmailWithSES(toAddress string, subject string, htmlBody string) error {
	ctx, cancel := context.WithTimeout(context.Background(), mailerTimeout)
	defer cancel()

	awsCfg, err := awsconfig.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("error al cargar la configuración de AWS: %w", err)
	}

	utf8 := "UTF-8"
	_, err = sesv2.NewFromConfig(awsCfg).SendEmail(ctx, &sesv2.SendEmailInput{
		FromEmailAddress: &Env.SES_FROM_EMAIL,
		Destination:      &sestypes.Destination{ToAddresses: []string{toAddress}},
		Content: &sestypes.EmailContent{
			Simple: &sestypes.Message{
				Subject: &sestypes.Content{Data: &subject, Charset: &utf8},
				Body:    &sestypes.Body{Html: &sestypes.Content{Data: &htmlBody, Charset: &utf8}},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("error al enviar el correo por SES: %w", err)
	}
	Log("SendEmail:: correo enviado por SES a", toAddress)
	return nil
}

func sendEmailWithSMTP(toAddress string, subject string, htmlBody string) error {
	fromAddress := Env.SMTP_FROM
	if len(fromAddress) == 0 {
		return errors.New("SMTP_FROM está vacío y SMTP_USER no lo suple")
	}

	address := net.JoinHostPort(Env.SMTP_HOST, strconv.Itoa(Env.SMTP_PORT))
	connection, err := net.DialTimeout("tcp", address, mailerTimeout)
	if err != nil {
		return fmt.Errorf("error al conectar con el servidor SMTP: %w", err)
	}
	// The deadline covers the whole conversation, not just the dial: STARTTLS and DATA are where
	// a hung server actually parks a caller.
	if err := connection.SetDeadline(time.Now().Add(mailerTimeout)); err != nil {
		connection.Close()
		return fmt.Errorf("error al fijar el tiempo límite del SMTP: %w", err)
	}

	client, err := smtp.NewClient(connection, Env.SMTP_HOST)
	if err != nil {
		connection.Close()
		return fmt.Errorf("error al iniciar la sesión SMTP: %w", err)
	}
	defer client.Close()

	if ok, _ := client.Extension("STARTTLS"); ok {
		if err := client.StartTLS(&tls.Config{ServerName: Env.SMTP_HOST}); err != nil {
			return fmt.Errorf("error al negociar STARTTLS: %w", err)
		}
	}
	if len(Env.SMTP_PASSWORD) > 0 {
		auth := smtp.PlainAuth("", Env.SMTP_USER, Env.SMTP_PASSWORD, Env.SMTP_HOST)
		if err := client.Auth(auth); err != nil {
			return fmt.Errorf("error al autenticar contra el SMTP: %w", err)
		}
	}
	if err := client.Mail(fromAddress); err != nil {
		return fmt.Errorf("error al declarar el remitente: %w", err)
	}
	if err := client.Rcpt(toAddress); err != nil {
		return fmt.Errorf("error al declarar el destinatario: %w", err)
	}

	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("error al abrir el cuerpo del correo: %w", err)
	}
	// Headers carry accented Spanish and an em dash, which are not ASCII and so cannot travel
	// raw: an unencoded byte above 0x7f is what turns a subject line into mojibake in half the
	// mail clients that read it.
	message := fmt.Sprintf("From: %v <%v>\r\nTo: %v\r\nSubject: %v\r\n"+
		"MIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%v",
		mime.QEncoding.Encode("UTF-8", Env.APP_NAME), fromAddress, toAddress,
		mime.QEncoding.Encode("UTF-8", subject), htmlBody)
	if _, err := writer.Write([]byte(message)); err != nil {
		return fmt.Errorf("error al escribir el cuerpo del correo: %w", err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("error al cerrar el cuerpo del correo: %w", err)
	}
	if err := client.Quit(); err != nil {
		return fmt.Errorf("error al cerrar la sesión SMTP: %w", err)
	}

	Log("SendEmail:: correo enviado por SMTP a", toAddress)
	return nil
}
