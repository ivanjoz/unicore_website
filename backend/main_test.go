package main

import (
	"testing"

	"app/core"
)

// The routes map is keyed on the normalised path, so a mismatch between the two is a 404 that
// looks exactly like a broken deployment. This is cheap insurance against that.
func TestEveryRouteIsReachableUnderItsOwnKey(t *testing.T) {
	for key := range routes {
		method, path, found := splitRouteKey(key)
		if !found {
			t.Fatalf("la clave %q no tiene la forma \"MÉTODO /ruta\"", key)
		}
		if normalized := normalizeRoute(path); normalized != path {
			t.Fatalf("la ruta %q no está normalizada; debería declararse como %q", path, normalized)
		}
		if _, reachable := routes[method+" "+normalizeRoute(path)]; !reachable {
			t.Fatalf("la ruta %q no es alcanzable con su propia clave", key)
		}
	}
}

func TestNormalizeRoute(t *testing.T) {
	cases := map[string]string{
		"/p-contact-message":  "/p-contact-message",
		"p-contact-message":   "/p-contact-message",
		"/p-contact-message/": "/p-contact-message",
		" /p-contact-message": "/p-contact-message",
		"":                    "/",
		"/":                   "/",
	}
	for input, expected := range cases {
		if actual := normalizeRoute(input); actual != expected {
			t.Fatalf("normalizeRoute(%q) = %q, se esperaba %q", input, actual, expected)
		}
	}
}

func TestUnknownRouteIs404(t *testing.T) {
	body := "{}"
	response := route(&core.HandlerArgs{
		Method: "POST", Route: "/no-existe", Body: &body, ClientIP: "203.0.113.7",
	})
	if response.StatusCode != 404 {
		t.Fatalf("se esperaba 404, llegó %v", response.StatusCode)
	}
}

// A handler that panics must become a 500, not take the process down: under Lambda a panic
// poisons a warm container that AWS will keep sending traffic to.
func TestPanicBecomesA500(t *testing.T) {
	routes["POST /p-panic-de-prueba"] = func(req *core.HandlerArgs) core.HandlerResponse {
		panic("estalla")
	}
	defer delete(routes, "POST /p-panic-de-prueba")

	body := "{}"
	response := route(&core.HandlerArgs{
		Method: "POST", Route: "/p-panic-de-prueba", Body: &body, ClientIP: "203.0.113.7",
	})
	if response.StatusCode != 500 {
		t.Fatalf("se esperaba 500, llegó %v", response.StatusCode)
	}
}

func splitRouteKey(key string) (method string, path string, found bool) {
	for index := 0; index < len(key); index++ {
		if key[index] == ' ' {
			return key[:index], key[index+1:], true
		}
	}
	return "", "", false
}
