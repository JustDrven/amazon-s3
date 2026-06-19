package router

import (
	"net/http"

	"github.com/charmbracelet/log"
	"justdrven.dev/storage/shared/src/security"
)

type EndpointMethod int

const (
	GET EndpointMethod = iota
	POST
	PUT
	DELETE
)

func getMethod(method EndpointMethod) string {
	switch method {
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case DELETE:
		return "DELETE"
	case GET:
	default:
		return "GET"
	}

	return "GET"
}

func (endpoint Endpoint) Register() {
	var finalPath string

	method := getMethod(endpoint.Method)
	path := endpoint.Path

	if method == "" {
		finalPath = path
	} else {
		finalPath = method + " " + path
	}

	domain := http.HandlerFunc(endpoint.Handler)
	http.Handle(finalPath, security.CommonMiddleware(domain))

	if method != "" {
		log.Info("Register new router", "METHOD", method, "PATH", path)
	} else {
		log.Info("Register new router", "PATH", path)
	}
}
