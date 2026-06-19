package cmd

type NotFoundResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
