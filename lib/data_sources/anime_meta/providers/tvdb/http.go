package tvdb

import (
	"animeta/lib/core/fetch"
	"io"
	"net/http"
	"strconv"
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

func fetchWrapper[T any](c *Client, url string) (TvdbResponse[T], error) {
	var out TvdbResponse[T]

	resp, err := c.TvdbAuthenticatedFetch("GET", url, nil)
	if err != nil {
		return out, err
	}

	err = fetch.ExtractResponseJsonBody(resp, &out)
	if err != nil {
		return out, err
	}

	if out.Status != "success" {
		return out, err
	}

	return out, nil
}

func (c *Client) FetchSeriesTranslations(id int) (TvdbSeriesTranslationsResponse, error) {
	return fetchWrapper[TvdbSeriesTranslationsData](
		c,
		BaseURL+"/series/"+strconv.Itoa(id)+"/extended?meta=translations&short=true",
	)
}

func (c *Client) FetchSeriesEpisodes(id int) (TvdbSeriesEpisodesResponse, error) {
	return fetchWrapper[TvdbSeriesEpisodesData](c, BaseURL+"/series/"+strconv.Itoa(id)+"/episodes/default/eng")
}

// TODO: fetch season here

func (c *Client) FetchMovieExtended(id int) (TvdbMovieResponse, error) {
	return fetchWrapper[TvdbMovieData](c, BaseURL+"/movies/"+strconv.Itoa(id)+"/extended?meta=translations&short=false")
}

func (c *Client) FetchRemoteId(id string) (TvdbRemoteIdResponse, error) {
	return fetchWrapper[TvdbRemoteIdData](c, BaseURL+"/search/remoteid/"+id)
}
