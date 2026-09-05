package endpoint_health

import (
	"encoding/json"
	"net/http"
	"time"
)

var startTime = time.Now()

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Uptime    string `json:"uptime"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Uptime:    time.Since(startTime).String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
