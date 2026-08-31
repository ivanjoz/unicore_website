package core

import (
	"os"
	"strconv"
	"strings"
)

// Every knob this backend has, read once from the process environment. There is no config file:
// the only deployment target is Lambda, where CloudFormation owns the environment block
// (cloud/template.yml), and a file would be a second source of truth for the same values.
type EnvVariables struct {
	// APP_NAME prefixes the physical resource names and is what the notification email signs off
	// with. APP_URL is only used in that email, to say where the message came from.
	APP_NAME string
	APP_URL  string

	// DYNAMO_TABLE is the single table everything lives in. DYNAMO_ENDPOINT points at a local
	// DynamoDB for tests and development; empty means the real service.
	DYNAMO_TABLE    string
	DYNAMO_ENDPOINT string

	// ALLOWED_ORIGINS is the comma-separated CORS allowlist. The frontend is a static site on
	// another origin, so without this the browser never delivers the response. A single "*" is
	// accepted for development; in production it should name the site, because this endpoint
	// costs money to invoke and an allowlist is the cheapest brake on being embedded elsewhere.
	ALLOWED_ORIGINS []string

	// CONTACT_EMAIL is the inbox the form delivers to. Empty disables the endpoint outright: a
	// message with no destination is only a row nobody will ever read.
	CONTACT_EMAIL string
	// CONTACT_MAX_MESSAGES_PER_IP and CONTACT_WINDOW_MINUTES are the rate limit. Defaults below.
	CONTACT_MAX_MESSAGES_PER_IP int64
	CONTACT_WINDOW_MINUTES      int64
	// CONTACT_TTL_DAYS expires stored messages through DynamoDB's TTL. Zero keeps them forever,
	// which is the default: a contact message is the whole point of the table.
	CONTACT_TTL_DAYS int64

	// The mailer picks SES when SES_FROM_EMAIL is set, SMTP when SMTP_HOST is, and refuses to
	// send otherwise. SES is the default because it needs no secret in the environment — the
	// Lambda role carries the permission — and this is an AWS-only deployment.
	SES_FROM_EMAIL string
	SMTP_HOST      string
	SMTP_PORT      int
	SMTP_USER      string
	SMTP_PASSWORD  string
	SMTP_FROM      string

	// LOGS_FULL prints every log line. Off in Lambda, where only errors are worth the ingestion
	// cost of CloudWatch, and on in the local server.
	LOGS_FULL bool
}

var Env = EnvVariables{}

// LoadEnv fills Env from the process environment. Called once, before the first request.
func LoadEnv() {
	Env.APP_NAME = envString("APP_NAME", "unicore")
	Env.APP_URL = envString("APP_URL", "https://un.pe")
	Env.DYNAMO_TABLE = envString("DYNAMO_TABLE", Env.APP_NAME+"-db")
	Env.DYNAMO_ENDPOINT = envString("DYNAMO_ENDPOINT", "")

	Env.ALLOWED_ORIGINS = []string{}
	for _, origin := range strings.Split(envString("ALLOWED_ORIGINS", "*"), ",") {
		if trimmed := strings.TrimSpace(origin); len(trimmed) > 0 {
			Env.ALLOWED_ORIGINS = append(Env.ALLOWED_ORIGINS, strings.ToLower(trimmed))
		}
	}

	Env.CONTACT_EMAIL = strings.ToLower(envString("CONTACT_EMAIL", ""))
	Env.CONTACT_MAX_MESSAGES_PER_IP = envInt("CONTACT_MAX_MESSAGES_PER_IP", 3)
	Env.CONTACT_WINDOW_MINUTES = envInt("CONTACT_WINDOW_MINUTES", 10)
	Env.CONTACT_TTL_DAYS = envInt("CONTACT_TTL_DAYS", 0)

	Env.SES_FROM_EMAIL = strings.TrimSpace(envString("SES_FROM_EMAIL", ""))
	Env.SMTP_HOST = envString("SMTP_HOST", "")
	Env.SMTP_PORT = int(envInt("SMTP_PORT", 587))
	Env.SMTP_USER = envString("SMTP_USER", "")
	Env.SMTP_PASSWORD = envString("SMTP_PASSWORD", "")
	Env.SMTP_FROM = envString("SMTP_FROM", Env.SMTP_USER)

	Env.LOGS_FULL = envString("LOGS_FULL", "") == "1"

	// A limit of zero would let the endpoint refuse everything, and a window of zero would make
	// every request its own window — both are silent ways of breaking the form, so a bad value
	// falls back to the default instead of being honoured.
	if Env.CONTACT_MAX_MESSAGES_PER_IP <= 0 {
		Env.CONTACT_MAX_MESSAGES_PER_IP = 3
	}
	if Env.CONTACT_WINDOW_MINUTES <= 0 {
		Env.CONTACT_WINDOW_MINUTES = 10
	}
}

func envString(name string, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(name)); len(value) > 0 {
		return value
	}
	return fallback
}

func envInt(name string, fallback int64) int64 {
	value := strings.TrimSpace(os.Getenv(name))
	if len(value) == 0 {
		return fallback
	}
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		Log("variable de entorno ignorada, no es un número::", name, value)
		return fallback
	}
	return parsed
}
