package middleware

import (
	"encoding/json"
	"net/http"
	"strconv"

	"animeta/lib/core/rate_limiter"
)

func RateLimit(rl *rate_limiter.Limiter, keyFunc func(*http.Request) string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := keyFunc(r)

			if !rl.Allow(key) {
				w.Header().Set("X-RateLimit-Remaining", "0")
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "rate limit exceeded",
				})
				return
			}

			w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(rl.Remaining(key)))
			next.ServeHTTP(w, r)
		})
	}
}

// KeyByIP usa l'IP del client come chiave. Se sei dietro un proxy/LB,
// controlla X-Forwarded-For prima di usare RemoteAddr.
func KeyByIP(r *http.Request) string {
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	return r.RemoteAddr
}
