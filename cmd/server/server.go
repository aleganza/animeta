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

	http.HandleFunc("/", endpoint_root.Handler)
	http.HandleFunc("/health", endpoint_health.Handler)

	http.HandleFunc("GET /meta/", endpoint_meta.Handler)
	http.HandleFunc("GET /meta/{provider}", endpoint_meta.Handler)
	http.HandleFunc("GET /meta/{provider}/{id}", endpoint_meta.Handler)

	log.Printf("Server listening on :%d...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
