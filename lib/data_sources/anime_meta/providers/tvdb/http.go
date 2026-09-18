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
	opts := &fetch.RequestOptions{
		Body: body,
		Headers: http.Header{
			"Authorization": {"Bearer " + c.Token},
		},
	}

	return fetch.MakeCachedRequest(method, path, opts)
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
		BaseURL+"/series/"+strconv.Itoa(id)+"/extended?meta=translations&short=false",
	)
}

func (c *Client) FetchSeriesEpisodes(id int, page int) (TvdbSeriesEpisodesResponse, error) {
	return fetchWrapper[TvdbSeriesEpisodesData](
		c,
		BaseURL+"/series/"+strconv.Itoa(id)+"/episodes/default/eng?page="+strconv.Itoa(page),
	)
}

func (c *Client) FetchAllSeriesEpisodes(id int) (TvdbSeriesEpisodesResponse, error) {
	var out TvdbSeriesEpisodesResponse

	page := 0

	for {
		resp, err := c.FetchSeriesEpisodes(id, page)
		if err != nil {
			return out, err
		}

		out.Data.Episodes = append(out.Data.Episodes, resp.Data.Episodes...)

		if resp.Links.Next == nil {
			break
		}

		page++
	}

	return out, nil
}

// TODO: fetch season here

func (c *Client) FetchMovieExtended(id int) (TvdbMovieResponse, error) {
	return fetchWrapper[TvdbMovieData](c, BaseURL+"/movies/"+strconv.Itoa(id)+"/extended?meta=translations&short=false")
}

func (c *Client) FetchRemoteId(id string) (TvdbRemoteIdResponse, error) {
	return fetchWrapper[TvdbRemoteIdData](c, BaseURL+"/search/remoteid/"+id)
}
