package fetch

import (
	"io"
	"net/http"
)

type Json map[string]any

type RequestOptions struct {
	Body    io.Reader
	Headers http.Header
}
