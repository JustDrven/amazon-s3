package common

type APIErrorResponse struct {
	Code    int    `xml:"Code" json:"code"`
	Message string `xml:"Message" json:"message"`
}
