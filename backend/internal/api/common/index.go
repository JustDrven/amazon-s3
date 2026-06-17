package common

import "net/http"

func setContentType(response http.ResponseWriter, types string) {
	response.Header().Set("Content-Type", types)
}

func SetJsonType(response http.ResponseWriter) {
	setContentType(response, "application/json")
}

func SetXmlType(response http.ResponseWriter) {
	setContentType(response, "application/xml")
}
