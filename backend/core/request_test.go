package core

import (
	"testing"
	"time"
)

// The limiter is only as good as this function: two requests that should share a budget must
// produce the same key, and two that should not must not.
func TestClientIPKeyIsStableAndPrefixedForIPv6(t *testing.T) {
	keyOf := func(address string) (int64, bool) {
		return (&HandlerArgs{ClientIP: address}).ClientIPKey()
	}

	ipv4Key, ok := keyOf("203.0.113.7")
	if !ok {
		t.Fatal("una IPv4 válida debería producir una clave")
	}
	if ipv4Key != 203<<24|0<<16|113<<8|7 {
		t.Fatalf("clave IPv4 inesperada: %v", ipv4Key)
	}

	// Whitespace around the value is what a header carries in practice, and it must not change
	// the key — otherwise the same visitor gets a fresh budget depending on the proxy.
	if paddedKey, _ := keyOf("  203.0.113.7 "); paddedKey != ipv4Key {
		t.Fatalf("los espacios cambiaron la clave: %v != %v", paddedKey, ipv4Key)
	}

	// Two addresses inside one customer's /64 must collapse to one key, or the limit is free to
	// bypass by picking another address from the same block.
	firstInBlock, _ := keyOf("2001:db8:abcd:1234::1")
	secondInBlock, _ := keyOf("2001:db8:abcd:1234:ffff:ffff:ffff:ffff")
	if firstInBlock != secondInBlock {
		t.Fatalf("dos direcciones del mismo /64 dieron claves distintas: %v != %v",
			firstInBlock, secondInBlock)
	}

	// A different /63 must not. Bit 63 of the prefix is the one the shift drops, so this pair is
	// the closest two blocks that are still required to be distinct.
	otherBlock, _ := keyOf("2001:db8:abcd:1236::1")
	if otherBlock == firstInBlock {
		t.Fatal("dos bloques /63 distintos colapsaron en la misma clave")
	}

	// IPv6 keys must stay positive and clear of the IPv4 range, so the two families can never
	// collide inside one partition key.
	if firstInBlock <= 1<<32 {
		t.Fatalf("la clave IPv6 cayó dentro del rango IPv4: %v", firstInBlock)
	}

	if _, ok := keyOf("no-es-una-ip"); ok {
		t.Fatal("una dirección inválida no debería producir clave")
	}
}

func TestWeekCodeAt(t *testing.T) {
	// 2026-08-06 is a Thursday in ISO week 32, so the code is 2026*100 + 32 - 200000.
	moment := time.Date(2026, 8, 6, 12, 0, 0, 0, time.UTC)
	if code := WeekCodeAt(moment); code != 2632 {
		t.Fatalf("código de semana inesperado: %v", code)
	}
}

func TestCorsHeadersOnlyEchoAnAllowedOrigin(t *testing.T) {
	Env.ALLOWED_ORIGINS = []string{"https://un.pe"}

	if origin := CorsHeaders("https://un.pe")["Access-Control-Allow-Origin"]; origin != "https://un.pe" {
		t.Fatalf("el origen permitido no fue devuelto: %q", origin)
	}
	// An origin nobody listed gets no header at all, which is what makes the browser block the
	// read. Echoing it back would make the allowlist decorative.
	if origin, present := CorsHeaders("https://evil.example")["Access-Control-Allow-Origin"]; present {
		t.Fatalf("un origen no permitido fue devuelto: %q", origin)
	}

	Env.ALLOWED_ORIGINS = []string{"*"}
	if origin := CorsHeaders("https://evil.example")["Access-Control-Allow-Origin"]; origin != "*" {
		t.Fatalf(`el comodín debería responder "*", respondió %q`, origin)
	}
}
