package api

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorResponse{
		Success: false,
		Message: message,
	})
}

func WriteSuccess[T any](w http.ResponseWriter, status int, message T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(SuccessResponse[T]{
		Success: true,
		Message: message,
	})
}
