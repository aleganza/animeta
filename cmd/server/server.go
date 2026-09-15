package main

import (
	endpoint_health "animeta/lib/api/endpoint/health"
	endpoint_meta "animeta/lib/api/endpoint/meta"
	endpoint_root "animeta/lib/api/endpoint/root"

	"animeta/lib/api/middleware"
	"animeta/lib/core/env"
	"animeta/lib/core/rate_limiter"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	bootstrap()

	port, err := env.GetIntVar("SERVER_PORT")
	if err != nil {
		fmt.Printf(`SERVER_PORT env variable not found, using default server port`)
		port = 8080
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", endpoint_root.Handler)
	mux.HandleFunc("/health", endpoint_health.Handler)

	mux.HandleFunc("GET /meta", endpoint_meta.Handler)
	mux.HandleFunc("GET /meta/{provider}", endpoint_meta.Handler)
	mux.HandleFunc("GET /meta/{provider}/{id}", endpoint_meta.Handler)

	log.Printf("Server listening on :%d...", port)

	// rate limiter: es. 100 richieste al minuto per IP
	rl := rate_limiter.New(5, time.Second, rate_limiter.WithTTL(1*time.Minute))
	rl.StartCleanup(time.Minute)
	defer rl.Stop()

	// catena: mux -> stripTrailingSlash -> rate limit
	handler := stripTrailingSlash(mux)
	handler = middleware.RateLimit(rl, middleware.KeyByIP)(handler)

	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handler,
	))
}
