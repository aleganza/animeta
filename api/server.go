package main

import (
	"animeta/api/endpoint/health"
	endpoint_root "animeta/api/endpoint/root"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", endpoint_root.Handler)
	http.HandleFunc("/health", endpoint_health.Handler)

	log.Println("Server in ascolto su :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
