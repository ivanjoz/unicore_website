package core

import (
	"encoding/binary"
	"encoding/json"
	"net"
	"strings"
	"time"
)

// HandlerArgs is one request, already normalised out of whichever of the two entrypoints
// delivered it: the Lambda event or the local http server. Handlers see only this, which is what
// lets the same handler be exercised by `go test` without an AWS account.
type HandlerArgs struct {
	Method string
	Route  string
	Body   *string
	// ClientIP is resolved by the caller, never by a handler — see the note on the Lambda
	// entrypoint about why the source of this value is a security property and not a detail.
	ClientIP  string
	Headers   map[string]string
	StartTime time.Time
}

// HandlerResponse is what an entrypoint turns back into its own response shape.
type HandlerResponse struct {
	StatusCode int
	Body       []byte
	Headers    map[string]string
}

type HandlerFunc func(req *HandlerArgs) HandlerResponse

// MakeResponse answers 200 with the JSON encoding of payload.
func (req *HandlerArgs) MakeResponse(payload any) HandlerResponse {
	body, err := json.Marshal(payload)
	if err != nil {
		return req.MakeErrCode("Error al serializar la respuesta.", 500)
	}
	return HandlerResponse{StatusCode: 200, Body: body, Headers: map[string]string{}}
}

// MakeErr is the 400 every rejected body gets. The extra values are logged, never sent: they are
// where the underlying error goes, and an unauthenticated caller has no business reading it.
func (req *HandlerArgs) MakeErr(message string, logValues ...any) HandlerResponse {
	return req.MakeErrCode(message, 400, logValues...)
}

func (req *HandlerArgs) MakeErrCode(message string, code int, logValues ...any) HandlerResponse {
	if len(logValues) > 0 {
		Log(append([]any{"error::", req.Route, message}, logValues...)...)
	}
	body, _ := json.Marshal(map[string]string{"error": message})
	return HandlerResponse{StatusCode: code, Body: body, Headers: map[string]string{}}
}

// ClientIPKey packs the caller's address into the int64 the rate limiter is keyed by.
//
// IPv6 is keyed by prefix rather than by address: a single residential customer is handed a whole
// /64, often a /56, so limiting per address would be free to bypass — one visitor with a /64 owns
// 18 quintillion of them. The prefix is shifted one bit to stay in positive int64 range, which
// keys the /63, still far narrower than any block a customer receives. Real IPv6 prefixes start
// at 2000::/3, so the result cannot collide with the IPv4 range that sits below 2^32.
func (req *HandlerArgs) ClientIPKey() (int64, bool) {
	parsed := net.ParseIP(strings.TrimSpace(req.ClientIP))
	if parsed == nil {
		return 0, false
	}
	if asIPv4 := parsed.To4(); asIPv4 != nil {
		return int64(binary.BigEndian.Uint32(asIPv4)), true
	}
	asIPv6 := parsed.To16()
	if asIPv6 == nil {
		return 0, false
	}
	return int64(binary.BigEndian.Uint64(asIPv6[:8]) >> 1), true
}

// WeekCodeAt is the [year][week] partition code for a moment in time: year*100 + isoWeek - 200000,
// so 2026-W32 is 2632.
//
// Messages are partitioned by the ISO week they arrived in because they belong to no account —
// whoever writes one does not have one. That keeps a table nobody reads in bulk from growing a
// single unbounded partition, and makes purging old messages a matter of dropping whole weeks.
func WeekCodeAt(moment time.Time) int32 {
	year, week := moment.ISOWeek()
	return int32(year*100 + week - 200000)
}

// CorsHeaders is the allowlist decision for one request's Origin.
//
// The allowlist is matched rather than echoed blindly: this endpoint writes to a database and
// sends mail on someone else's bill, and reflecting any Origin would let any page on the internet
// spend both. "*" is honoured for development, and cannot carry credentials — which is fine,
// because the form sends none.
func CorsHeaders(origin string) map[string]string {
	headers := map[string]string{
		"Access-Control-Allow-Methods": "POST, OPTIONS",
		"Access-Control-Allow-Headers": "Content-Type",
		"Access-Control-Max-Age":       "86400",
		"Vary":                         "Origin",
	}
	for _, allowed := range Env.ALLOWED_ORIGINS {
		if allowed == "*" {
			headers["Access-Control-Allow-Origin"] = "*"
			return headers
		}
		if allowed == strings.ToLower(strings.TrimSpace(origin)) {
			headers["Access-Control-Allow-Origin"] = origin
			return headers
		}
	}
	// No Allow-Origin header at all: the browser blocks the read, which is the intended answer
	// for an origin nobody put on the list.
	return headers
}
