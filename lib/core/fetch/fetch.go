package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Do(req *http.Request) (*http.Response, error) {
	log.Printf("%s %s", req.Method, req.URL)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		return nil, fmt.Errorf("request to %s failed: %s", req.URL, resp.Status)
	}
	return resp, nil
}

func Request(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// can't remember why I have Do and Request separated. using ExecuteRequest until I don't find out
func ExecuteRequest(method, path string, opts *RequestOptions) (*http.Response, error) {
	var body io.Reader

	if opts != nil {
			body = opts.Body
	}

	req, err := Request(method, path, body)
	if err != nil {
			return nil, err
	}

	if opts != nil && opts.Headers != nil {
			req.Header = opts.Headers
	}

	return Do(req)
}

func ExtractResponseJsonBody(resp *http.Response, dst any) error {
	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(dst)
	if err != nil {
		return err
	}

	return nil
}
