package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"app/core"
	"app/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Every route this backend serves, keyed by "METHOD /path". A map and not a mux because there is
// one endpoint and no path parameters; the day either changes, this becomes an http.ServeMux and
// the two entrypoints below keep working unchanged.
//
// The "p-" prefix marks a public route, as it does in the genix backend this is modelled on: no
// token, no account, and therefore the per-IP limiter is the only thing standing in front of it.
var routes = map[string]core.HandlerFunc{
	"POST /p-contact-message": handlers.PostContactMessage,
}

func main() {
	core.LoadEnv()

	// AWS_LAMBDA_RUNTIME_API is set by the Lambda runtime and by nothing else, which makes it the
	// one reliable way to tell the two modes apart without a build tag or a flag somebody has to
	// remember to pass.
	if len(os.Getenv("AWS_LAMBDA_RUNTIME_API")) > 0 {
		lambda.Start(handleLambdaRequest)
		return
	}

	core.Env.LOGS_FULL = true
	port := strings.TrimSpace(os.Getenv("PORT"))
	if len(port) == 0 {
		port = "3333"
	}
	core.Log("*Servidor local escuchando en http://localhost:" + port)
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           http.HandlerFunc(handleLocalRequest),
		ReadHeaderTimeout: 10 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		core.Log("error al levantar el servidor local::", err)
		os.Exit(1)
	}
}

// route dispatches one already-normalised request and is where a panic stops.
//
// A handler that panics must answer 500 rather than take the process down: under the local server
// that would end every in-flight request, and under Lambda it poisons a warm container that AWS
// will keep sending traffic to. The stack goes to the log, never to the caller.
func route(req *core.HandlerArgs) (response core.HandlerResponse) {
	defer func() {
		if recovered := recover(); recovered != nil {
			core.Log("error:: pánico en", req.Method, req.Route, recovered, string(debug.Stack()))
			response = req.MakeErrCode("Error interno.", 500)
		}
	}()

	handler, found := routes[req.Method+" "+req.Route]
	if !found {
		return req.MakeErrCode("La ruta solicitada no existe.", 404)
	}
	return handler(req)
}

// normalizeRoute is what the routes map is keyed on. Exactly one leading slash and no trailing
// one, so "/p-contact-message" and "p-contact-message/" reach the same handler instead of the
// second one being a 404 that looks like a deployment problem.
func normalizeRoute(path string) string {
	trimmed := strings.Trim(strings.TrimSpace(path), "/")
	return "/" + trimmed
}

// finish merges the CORS decision into the handler's own headers and guarantees a content type.
// Both entrypoints go through it so neither can forget one of the two.
func finish(response core.HandlerResponse, origin string) core.HandlerResponse {
	if response.Headers == nil {
		response.Headers = map[string]string{}
	}
	for name, value := range core.CorsHeaders(origin) {
		response.Headers[name] = value
	}
	if _, hasContentType := response.Headers["Content-Type"]; !hasContentType {
		response.Headers["Content-Type"] = "application/json; charset=utf-8"
	}
	return response
}

// ---------------------------------------------------------------------------- Lambda entrypoint

// The event type is the API Gateway v2 shape, which a Lambda Function URL also sends. Using the
// shared type means the same binary serves either front door without a second handler.
func handleLambdaRequest(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	method := strings.ToUpper(request.RequestContext.HTTP.Method)
	origin := headerValue(request.Headers, "origin")

	// The preflight is answered here rather than by the Function URL's own CORS block, so the
	// allowlist has exactly one definition — ALLOWED_ORIGINS — instead of one in the environment
	// and a second one in CloudFormation that would silently disagree with it.
	if method == "OPTIONS" {
		response := finish(core.HandlerResponse{StatusCode: 204}, origin)
		return events.APIGatewayV2HTTPResponse{
			StatusCode: response.StatusCode, Headers: response.Headers,
		}, nil
	}

	body := request.Body
	if request.IsBase64Encoded {
		decoded, err := decodeBase64Body(request.Body)
		if err != nil {
			core.Log("error:: no se pudo decodificar el cuerpo en base64", err)
			decoded = ""
		}
		body = decoded
	}

	req := &core.HandlerArgs{
		Method:    method,
		Route:     normalizeRoute(request.RawPath),
		Body:      &body,
		ClientIP:  lambdaClientIP(&request),
		Headers:   request.Headers,
		StartTime: time.Now(),
	}

	response := finish(route(req), origin)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: response.StatusCode,
		Headers:    response.Headers,
		Body:       string(response.Body),
	}, nil
}

// lambdaClientIP resolves the address the rate limiter keys on, and is the one place in this
// backend where getting it wrong costs the whole limit.
//
// SourceIP is written by AWS from the TCP peer of the Function URL and cannot be forged. Headers
// can: X-Forwarded-For is appended to, so a client that sends its own lands first in the list, and
// anything limited by that value would be bypassable with one curl flag. It is therefore never
// read.
//
// CLIENT_IP_HEADER is the deliberate exception, for the case where a CDN sits in front and
// SourceIP is the CDN's own address — every visitor would then share one key and the first three
// of them would exhaust the limit for everybody. Setting it is only safe when the Function URL
// cannot be reached directly, because from that moment the header is as trustworthy as whoever
// can reach the origin.
func lambdaClientIP(request *events.APIGatewayV2HTTPRequest) string {
	if headerName := strings.TrimSpace(os.Getenv("CLIENT_IP_HEADER")); len(headerName) > 0 {
		if value := headerValue(request.Headers, headerName); len(value) > 0 {
			// CloudFront-Viewer-Address carries "ip:port"; a plain viewer-address header does not.
			// Splitting only when it parses leaves a bare IPv6 address intact.
			if host, _, err := net.SplitHostPort(value); err == nil {
				return host
			}
			return value
		}
	}
	return request.RequestContext.HTTP.SourceIP
}

// API Gateway lowercases header names, but a Function URL invoked directly does not always, and
// nothing in the contract promises either. Matching case-insensitively costs one comparison.
func headerValue(headers map[string]string, name string) string {
	if value, found := headers[name]; found {
		return value
	}
	lowered := strings.ToLower(name)
	for key, value := range headers {
		if strings.ToLower(key) == lowered {
			return value
		}
	}
	return ""
}

// ----------------------------------------------------------------------- local server entrypoint

// The local server exists so the handler can be exercised end to end against a DynamoDB Local
// container, without deploying. It is not a production target: there is no TLS here and no
// throttle in front of it.
func handleLocalRequest(writer http.ResponseWriter, request *http.Request) {
	origin := request.Header.Get("Origin")

	if request.Method == http.MethodOptions {
		response := finish(core.HandlerResponse{StatusCode: 204}, origin)
		writeLocalResponse(writer, response)
		return
	}

	// Bounded before it is read, not after: an unbounded ReadAll on a public endpoint is a way to
	// spend all the memory the process has. The ceiling is well above the 4000-rune message the
	// form allows, even with every rune at 4 bytes and JSON escaping on top.
	bodyBytes, err := io.ReadAll(io.LimitReader(request.Body, 64*1024))
	if err != nil {
		writeLocalResponse(writer, finish(core.HandlerResponse{
			StatusCode: 400, Body: []byte(`{"error":"No se pudo leer el cuerpo."}`),
		}, origin))
		return
	}
	body := string(bodyBytes)

	headers := map[string]string{}
	for name := range request.Header {
		headers[name] = request.Header.Get(name)
	}

	req := &core.HandlerArgs{
		Method:    strings.ToUpper(request.Method),
		Route:     normalizeRoute(request.URL.Path),
		Body:      &body,
		ClientIP:  localClientIP(request),
		Headers:   headers,
		StartTime: time.Now(),
	}

	writeLocalResponse(writer, finish(route(req), origin))
}

func localClientIP(request *http.Request) string {
	if realIP := strings.TrimSpace(request.Header.Get("X-Real-IP")); len(realIP) > 0 {
		return realIP
	}
	host, _, err := net.SplitHostPort(request.RemoteAddr)
	if err != nil {
		return strings.TrimSpace(request.RemoteAddr)
	}
	return host
}

func writeLocalResponse(writer http.ResponseWriter, response core.HandlerResponse) {
	for name, value := range response.Headers {
		writer.Header().Set(name, value)
	}
	writer.WriteHeader(response.StatusCode)
	if len(response.Body) > 0 {
		if _, err := writer.Write(response.Body); err != nil {
			core.Log("error al escribir la respuesta::", err)
		}
	}
}

// A Function URL base64-encodes the body whenever it decides the content type is binary. The form
// posts JSON and normally arrives as text, so this is the path nothing takes until a client omits
// its Content-Type — at which point a body silently read as base64 would be an unexplainable
// parse error.
func decodeBase64Body(body string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		return "", fmt.Errorf("cuerpo en base64 inválido: %w", err)
	}
	return string(decoded), nil
}
