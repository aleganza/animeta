package tenrai

import (
	"animeta/lib/core/env"
	"animeta/lib/core/fetch"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func NewClient() (Client, error) {
	serverKey, _ := env.GetVar("TENRAI_SERVER_KEY")

	return Client{
		ServerKey: serverKey,
	}, nil
}

func (c *Client) fetch(path string, dst any) error {
	req, err := fetch.Request("GET", BaseURL+path, nil)
	if err != nil {
		return err
	}

	if c.ServerKey != "" {
		req.Header.Set("X-Server-Key", c.ServerKey)
	}

	resp, err := fetch.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		resp.Body.Close()
		retry := resp.Header.Get("Retry-After")
		if seconds, err := strconv.Atoi(retry); err == nil {
			time.Sleep(time.Duration(seconds) * time.Second)
			return c.fetch(path, dst)
		}
		return fmt.Errorf("tenrai rate limit exceeded, retry after %q", retry)
	}

	return fetch.ExtractResponseJsonBody(resp, dst)
}

func (c *Client) FetchAnimeEpisodes(malId int, page int) (AnimeEpisodesResponse, error) {
	var out AnimeEpisodesResponse

	err := c.fetch("/anime/"+strconv.Itoa(malId)+"/episodes?page="+strconv.Itoa(page), &out)
	if err != nil {
		return out, err
	}

	return out, nil
}
