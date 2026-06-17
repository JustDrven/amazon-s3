package router

import "net/http"

type Endpoint struct {
	Method string
	Path   string

	Handler func(http.ResponseWriter, *http.Request)
}
