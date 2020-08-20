package rest

var (
	pageNotFoundError = responseError{
		Code:    100,
		Message: "page does not exist",
	}
)

type errorResponse struct {
	Errors []responseError `json:"errors"`
}

type responseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
