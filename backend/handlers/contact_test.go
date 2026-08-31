package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"app/core"
	"app/db"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func postContact(clientIP string, body map[string]string) core.HandlerResponse {
	encoded, _ := json.Marshal(body)
	asString := string(encoded)
	return PostContactMessage(&core.HandlerArgs{
		Method: "POST", Route: "p-contact-message",
		Body: &asString, ClientIP: clientIP, StartTime: time.Now(),
	})
}

func errorMessage(t *testing.T, response core.HandlerResponse) string {
	t.Helper()
	payload := map[string]string{}
	if err := json.Unmarshal(response.Body, &payload); err != nil {
		t.Fatalf("la respuesta no es JSON: %v", string(response.Body))
	}
	return payload["error"]
}

func validBody() map[string]string {
	return map[string]string{
		"Name":    "Ada Lovelace",
		"Email":   "ada@example.com",
		"Company": "Analytical Engines",
		"Message": "Quisiera conversar sobre una iniciativa de código abierto.",
	}
}

// Everything below rejects before the limiter is consulted, so none of it needs a database — which
// is also the point: a malformed body must never reach a write.
func TestContactRejectsBadInputWithoutTouchingTheDatabase(t *testing.T) {
	core.LoadEnv()
	core.Env.CONTACT_EMAIL = "contacto@un.pe"
	// Deliberately unset, so any test below that reached DynamoDB would fail loudly instead of
	// quietly talking to whatever account the developer happens to be logged into.
	core.Env.DYNAMO_TABLE = ""

	cases := []struct {
		name     string
		mutate   func(body map[string]string)
		expected string
	}{
		{"nombre vacío", func(b map[string]string) { b["Name"] = " " }, "El nombre es necesario."},
		{"nombre de una letra", func(b map[string]string) { b["Name"] = "A" }, "El nombre es necesario."},
		{"nombre larguísimo", func(b map[string]string) { b["Name"] = strings.Repeat("á", 121) },
			"El nombre no puede exceder 120 caracteres."},
		{"correo sin arroba", func(b map[string]string) { b["Email"] = "ada.example.com" },
			"El correo electrónico no posee un formato válido."},
		{"correo sin dominio", func(b map[string]string) { b["Email"] = "ada@example" },
			"El correo electrónico no posee un formato válido."},
		{"empresa larguísima", func(b map[string]string) { b["Company"] = strings.Repeat("e", 161) },
			"La empresa no puede exceder 160 caracteres."},
		{"mensaje corto", func(b map[string]string) { b["Message"] = "hola" },
			"El mensaje debe poseer al menos 10 caracteres."},
		{"mensaje larguísimo", func(b map[string]string) { b["Message"] = strings.Repeat("m", 4001) },
			"El mensaje no puede exceder 4000 caracteres."},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			body := validBody()
			testCase.mutate(body)
			response := postContact("203.0.113.7", body)
			if response.StatusCode != 400 {
				t.Fatalf("se esperaba 400, llegó %v", response.StatusCode)
			}
			if message := errorMessage(t, response); message != testCase.expected {
				t.Fatalf("mensaje inesperado: %q", message)
			}
		})
	}

	// A message whose length is only reached by counting bytes instead of runes must be accepted:
	// 4000 accented characters are 8000 bytes, and cutting the visitor off at half the limit the
	// form advertises is the bug the rune count exists to prevent.
	t.Run("acepta acentos hasta el límite en runas", func(t *testing.T) {
		body := validBody()
		body["Message"] = strings.Repeat("á", contactMaxMessageRunes)
		if response := postContact("203.0.113.7", body); response.StatusCode == 400 {
			t.Fatalf("un mensaje de 4000 runas fue rechazado: %v", errorMessage(t, response))
		}
	})
}

func TestContactRefusesWhenNoInboxIsConfigured(t *testing.T) {
	core.LoadEnv()
	core.Env.CONTACT_EMAIL = ""

	response := postContact("203.0.113.7", validBody())
	if response.StatusCode != 503 {
		t.Fatalf("se esperaba 503, llegó %v", response.StatusCode)
	}
}

func TestContactRefusesAnUnresolvableClient(t *testing.T) {
	core.LoadEnv()
	core.Env.CONTACT_EMAIL = "contacto@un.pe"

	// No address means no key, and no key means no limit. Serving the request anyway would be an
	// unmetered path into the inbox for anyone who can suppress the header.
	response := postContact("", validBody())
	if response.StatusCode != 400 {
		t.Fatalf("se esperaba 400, llegó %v", response.StatusCode)
	}
}

// The end-to-end path, against a real DynamoDB. See db/ratelimit_test.go for how to start one.
func TestContactStoresAndThenThrottles(t *testing.T) {
	if len(os.Getenv("DYNAMO_ENDPOINT")) == 0 {
		t.Skip("DYNAMO_ENDPOINT no está definido; se omite la prueba de integración")
	}

	core.LoadEnv()
	core.Env.CONTACT_EMAIL = "contacto@un.pe"
	core.Env.CONTACT_MAX_MESSAGES_PER_IP = 2
	core.Env.CONTACT_WINDOW_MINUTES = 5
	core.Env.DYNAMO_TABLE = "unicore-test-db"
	// No mailer is configured, so every send fails — which is the interesting half of this test:
	// the message must still be stored and the visitor must still be told it arrived.
	core.Env.SES_FROM_EMAIL = ""
	core.Env.SMTP_HOST = ""
	createTestTable(t)

	// A fresh address per run, so a rerun inside the same window does not start out throttled.
	clientIP := fmt.Sprintf("198.51.%d.%d", time.Now().Unix()%250, (time.Now().UnixNano()/1000)%250)

	for attempt := 1; attempt <= 2; attempt++ {
		response := postContact(clientIP, validBody())
		if response.StatusCode != 200 {
			t.Fatalf("el mensaje %v fue rechazado con %v: %v",
				attempt, response.StatusCode, errorMessage(t, response))
		}
		payload := map[string]bool{}
		if err := json.Unmarshal(response.Body, &payload); err != nil {
			t.Fatalf("respuesta ilegible: %v", string(response.Body))
		}
		if !payload["Received"] {
			t.Fatal("el mensaje debería declararse recibido")
		}
		if payload["Notified"] {
			t.Fatal("sin mailer configurado, Notified debería ser falso")
		}
	}

	response := postContact(clientIP, validBody())
	if response.StatusCode != 429 {
		t.Fatalf("el tercer mensaje debería dar 429, dio %v", response.StatusCode)
	}
	if retryAfter := response.Headers["Retry-After"]; len(retryAfter) == 0 {
		t.Fatal("un 429 sin Retry-After no le dice al visitante cuándo volver")
	}
	if message := errorMessage(t, response); !strings.Contains(message, "máximo de 2 mensajes") {
		t.Fatalf("el mensaje de rechazo no nombra el límite: %q", message)
	}

	// Both stored messages must survive as two rows: the sort key carries the limiter's counter
	// precisely so two submissions inside one second cannot overwrite each other.
	if stored := countStoredMessages(t, clientIP); stored != 2 {
		t.Fatalf("se almacenaron %v mensajes, se esperaban 2", stored)
	}
}

func testClient(t *testing.T) *dynamodb.Client {
	t.Helper()
	client, err := db.Client()
	if err != nil {
		t.Fatalf("no se pudo crear el cliente: %v", err)
	}
	return client
}

func createTestTable(t *testing.T) {
	t.Helper()
	_, err := testClient(t).CreateTable(context.Background(), &dynamodb.CreateTableInput{
		TableName:   aws.String(core.Env.DYNAMO_TABLE),
		BillingMode: types.BillingModePayPerRequest,
		AttributeDefinitions: []types.AttributeDefinition{
			{AttributeName: aws.String("pk"), AttributeType: types.ScalarAttributeTypeS},
			{AttributeName: aws.String("sk"), AttributeType: types.ScalarAttributeTypeS},
		},
		KeySchema: []types.KeySchemaElement{
			{AttributeName: aws.String("pk"), KeyType: types.KeyTypeHash},
			{AttributeName: aws.String("sk"), KeyType: types.KeyTypeRange},
		},
	})
	var alreadyExists *types.ResourceInUseException
	if err != nil && !errors.As(err, &alreadyExists) {
		t.Fatalf("no se pudo crear la tabla: %v", err)
	}
}

// Reads back through the sort-key prefix, which is the query the schema was shaped for: "what did
// this address send this week" is a range over one partition, not a scan.
func countStoredMessages(t *testing.T, clientIP string) int {
	t.Helper()
	ipKey, ok := (&core.HandlerArgs{ClientIP: clientIP}).ClientIPKey()
	if !ok {
		t.Fatalf("no se pudo calcular la clave de %v", clientIP)
	}

	output, err := testClient(t).Query(context.Background(), &dynamodb.QueryInput{
		TableName:              aws.String(core.Env.DYNAMO_TABLE),
		KeyConditionExpression: aws.String("pk = :pk AND begins_with(sk, :ip)"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":pk": &types.AttributeValueMemberS{
				Value: fmt.Sprintf("MSG#%d", core.WeekCodeAt(time.Now()))},
			":ip": &types.AttributeValueMemberS{Value: fmt.Sprintf("%019d#", ipKey)},
		},
	})
	if err != nil {
		t.Fatalf("no se pudieron leer los mensajes: %v", err)
	}
	return len(output.Items)
}
