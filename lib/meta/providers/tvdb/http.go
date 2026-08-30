package tvdb

import (
	"animeta/lib/core/fetch"
	"io"
	"net/http"
)

// Source - https://stackoverflow.com/a/54088988
// Posted by Jonathan Hall, modified by community. See post 'Timeline' for change history
// Retrieved 2026-07-28, License - CC BY-SA 4.0

func (c *Client) TvdbAuthenticatedFetch(method, path string, body io.Reader) (*http.Response, error) {
	req, err := fetch.Request(method, path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.Token)

	resp, err := fetch.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
