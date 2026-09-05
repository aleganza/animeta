package api

import (
	"encoding/json"
	"net/http"
)

func WriteError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorResponse{
		Status:  status,
		Message: message,
	})
}
