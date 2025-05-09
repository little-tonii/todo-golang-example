package godoc

type ErrorResponse struct {
	Message string `json:"message"`
}

type ErrorsResponse struct {
	Messages []string `json:"messages"`
}
