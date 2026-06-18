package common

import (
	"errors"
	"net/http"
	"strings"
)

const (
	OAUTH_SCHEMA_KEY  = "OAuth "
	BEARER_SCHEMA_KEY = "Bearer "

	AUTH_HEADER_KEY = "Authorization"

	INVALID_SCHEMA_VALUE = "Invalid schema"
)

func setContentType(response http.ResponseWriter, types string) {
	response.Header().Set("Content-Type", types)
}

func SetJsonType(response http.ResponseWriter) {
	setContentType(response, "application/json")
}

func SetXmlType(response http.ResponseWriter) {
	setContentType(response, "application/xml")
}

func GetAuthoritation(endUserLogin bool, req *http.Request) (string, error) {
	auth := req.Header.Get(AUTH_HEADER_KEY)

	if endUserLogin {
		if !strings.HasPrefix(auth, OAUTH_SCHEMA_KEY) {
			goto invalidSchema
		}
		token, _ := strings.CutPrefix(auth, OAUTH_SCHEMA_KEY)

		return token, nil
	} else {
		if !strings.HasPrefix(auth, BEARER_SCHEMA_KEY) {
			goto invalidSchema
		}

		token, _ := strings.CutPrefix(auth, BEARER_SCHEMA_KEY)

		return token, nil

	}

invalidSchema:
	return "", errors.New(INVALID_SCHEMA_VALUE)

}
