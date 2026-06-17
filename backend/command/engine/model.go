package engine

type NotFoundResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
