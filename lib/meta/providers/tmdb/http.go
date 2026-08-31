package themoviedb

import (
	"animeta/lib/core/fetch"
	"io"
	"net/http"
)

func (c *Client) AuthenticatedFetch(method, path string, body io.Reader) (*http.Response, error) {
	req, err := fetch.Request(method, path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("accept", "application/json")

	resp, err := fetch.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
