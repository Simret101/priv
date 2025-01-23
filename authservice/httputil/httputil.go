package httputil

// HTTPError represents an error that occurred while handling a request.
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Invalid request"`
}
