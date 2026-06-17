package router

import (
	"net/http"

	"github.com/charmbracelet/log"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func (endpoint Endpoint) Register() {
	var finalPath string

	method := endpoint.Method
	path := endpoint.Path

	if len(method) == 0 {
		finalPath = path
	} else {
		finalPath = method + " " + path
	}

	http.HandleFunc(finalPath, endpoint.Handler)
	if method != "" {
		log.Info("Register new router", "METHOD", method, "PATH", path)
	} else {
		log.Info("Register new router", "PATH", path)
	}
}
