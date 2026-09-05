package main

import (
	"net/http"
	"strings"
)

// stripTrailingSlash removes a trailing slash before routing.
// Example: "/meta/tvdb/123/" becomes "/meta/tvdb/123".
func stripTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r2 := r.Clone(r.Context())
			r2.URL.Path = strings.TrimRight(r.URL.Path, "/")
			next.ServeHTTP(w, r2)
			return
		}

		next.ServeHTTP(w, r)
	})
}
