package api

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse[T any] struct {
	Status  int `json:"status"`
	Message T   `json:"data"`
}
