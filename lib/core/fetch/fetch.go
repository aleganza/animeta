package fetch

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

func do(req *http.Request) (*http.Response, error) {
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

func request(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func MakeRequest(method, path string, opts *RequestOptions) (*http.Response, error) {
	var body io.Reader

	if opts != nil {
		body = opts.Body
	}

	req, err := request(method, path, body)
	if err != nil {
		return nil, err
	}

	if opts != nil && opts.Headers != nil {
		req.Header = opts.Headers
	}

	return do(req)
}

func ExtractResponseJsonBody(resp *http.Response, dst any) error {
	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

func ExtractResponseXmlBody(resp *http.Response, dst any) error {
	defer resp.Body.Close()

	err := xml.NewDecoder(resp.Body).Decode(dst)
	if err != nil {
		return err
	}

	return nil
}
