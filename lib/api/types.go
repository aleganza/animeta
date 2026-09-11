package api

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SuccessResponse[T any] struct {
	Success bool `json:"success"`
	Message T    `json:"data"`
}
