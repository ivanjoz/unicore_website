package core

import (
	"fmt"
	"log"
	"strings"
)

// Log is the one way this backend writes to CloudWatch.
//
// Outside LOGS_FULL only lines that carry an error or a warning survive. CloudWatch bills by
// ingested byte and a contact form is invoked by strangers, so a chatty info line is a cost an
// attacker controls; an error line is the one thing worth paying for every time.
func Log(values ...any) {
	message := makeLogLine(values...)
	if !Env.LOGS_FULL && !isNoteworthy(message) {
		return
	}
	log.Println(message)
}

func makeLogLine(values ...any) string {
	parts := make([]string, 0, len(values))
	for _, value := range values {
		parts = append(parts, fmt.Sprint(value))
	}
	return strings.Join(parts, " ")
}

func isNoteworthy(message string) bool {
	lowered := strings.ToLower(message)
	return strings.Contains(lowered, "error") || strings.Contains(lowered, "warn") ||
		strings.Contains(lowered, "límite") || strings.Contains(lowered, "no se pudo")
}

// StrCut bounds a value before it reaches a log line. Bodies here are attacker-controlled and up
// to CONTACT_MAX_MESSAGE_RUNES long; logging one whole is how a form becomes a way to write into
// somebody else's bill.
func StrCut(value string, max int) string {
	runes := []rune(value)
	if len(runes) <= max {
		return value
	}
	return string(runes[:max]) + "…"
}
