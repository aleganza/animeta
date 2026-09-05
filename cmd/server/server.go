package main

import (
	"animeta/lib/api/endpoint/health"
	"animeta/lib/api/endpoint/meta"
	"animeta/lib/api/endpoint/root"
	"animeta/lib/core/env"
	"fmt"
	"log"
	"net/http"
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

	handler := stripTrailingSlash(mux)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		handler,
	))
}
