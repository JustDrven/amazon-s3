package api

type APIErrorResponse struct {
	Code      string `xml:"Code"`
	Message   string `xml:"Message"`
	RequestId string `xml:"RequestId"`
	HostId    string `xml:"HostId"`
}
