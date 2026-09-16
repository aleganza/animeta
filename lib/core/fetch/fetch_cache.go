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

	entryKey := resolveCacheKey(method, path, requestBody)

	if value, found := cache.Get(entryKey); found {
		log.Printf("CACHED request to: %s %s", method, path)
		
		return &http.Response{
			StatusCode:    value.StatusCode,
			Status:        value.Status,
			Header:        value.Header,
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
		Header:     resp.Header,
		Body:       body,
	}

	cache.Set(entryKey, value, int64(len(body)), ttl)

	resp.Body = io.NopCloser(bytes.NewReader(body))
	resp.ContentLength = int64(len(body))

	return resp, nil
}

func resolveCacheKey(method, path string, body []byte) string {
	hash := sha256.New()

	hash.Write([]byte(method))
	hash.Write([]byte{0})
	hash.Write([]byte(path))
	hash.Write([]byte{0})
	hash.Write(body)

	return hex.EncodeToString(hash.Sum(nil))
}