package router

import "net/http"

type Endpoint struct {
	Method EndpointMethod
	Path   string

	Handler func(http.ResponseWriter, *http.Request)
}
