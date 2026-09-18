package fetch

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	cache_ristretto "animeta/lib/core/cache/providers/ristretto"
)

type CachedResponse struct {
	StatusCode int
	Status     string
	Header     http.Header
	Body       []byte
}

var cache = cache_ristretto.New[string, CachedResponse]()

var ttl = 24 * time.Hour

// headers that influence the response and must be part of the cache key
var cacheKeyHeaders = []string{"Authorization", "X-Server-Key"}

func MakeCachedRequest(
	method string,
	path string,
	opts *RequestOptions,
) (*http.Response, error) {

	var requestBody []byte

	if opts != nil && opts.Body != nil {
		var err error

		requestBody, err = io.ReadAll(opts.Body)
		if err != nil {
			return nil, fmt.Errorf("could not read request body: %w", err)
		}

		opts.Body = bytes.NewReader(requestBody)
	}

	var headers http.Header
	if opts != nil {
		headers = opts.Headers
	}

	entryKey := resolveCacheKey(method, path, requestBody, headers)

	if value, found := cache.Get(entryKey); found {
		log.Printf("CACHED request to: %s %s", method, path)

		// expose a copy so callers cannot mutate the cached header map
		return &http.Response{
			StatusCode:    value.StatusCode,
			Status:        value.Status,
			Header:        value.Header.Clone(),
			Body:          io.NopCloser(bytes.NewReader(value.Body)),
			ContentLength: int64(len(value.Body)),
		}, nil
	}

	resp, err := MakeRequest(method, path, opts)
	if err != nil {
		return nil, fmt.Errorf("request to %q failed: %w", path, err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
	}

	value := CachedResponse{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		// store a copy so the live response headers cannot mutate the entry
		Header: resp.Header.Clone(),
		Body:   body,
	}

	cache.Set(entryKey, value, int64(len(body)), ttl)

	resp.Body = io.NopCloser(bytes.NewReader(body))
	resp.ContentLength = int64(len(body))

	return resp, nil
}

func resolveCacheKey(method, path string, body []byte, headers http.Header) string {
	hash := sha256.New()

	hash.Write([]byte(method))
	hash.Write([]byte{0})
	hash.Write([]byte(path))
	hash.Write([]byte{0})

	// only include headers that affect the response, so changing
	// unrelated headers does not invalidate every cache entry
	for _, name := range cacheKeyHeaders {
		hash.Write([]byte(name))
		hash.Write([]byte{0})
		hash.Write([]byte(headers.Get(name)))
		hash.Write([]byte{0})
	}

	hash.Write(body)

	return hex.EncodeToString(hash.Sum(nil))
}
