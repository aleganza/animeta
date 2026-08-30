package tvdb

import (
	"animeta/lib/core/fetch"
)

func fetchWrapper[T any](c *Client, url string) (T, error) {
	var out T

	resp, err := c.TvdbAuthenticatedFetch("GET", url, nil)
	if err != nil {
		return out, err
	}

	if err := fetch.ExtractResponseJsonBody(resp, &out); err != nil {
		return out, err
	}

	return out, nil
}

func (c *Client) FetchSeriesExtended(id string) (TvdbSeriesExtendedResponse, error) {
	return fetchWrapper[TvdbSeriesExtendedResponse](c, BaseURL+"/series/"+id+"/extended?meta=translations&short=false")
}

func (c *Client) FetchSeriesEpisodes(id string) (TvdbSeriesEpisodesResponse, error) {
	return fetchWrapper[TvdbSeriesEpisodesResponse](c, BaseURL+"/series/"+id+"/episodes/default/eng")
}

func (c *Client) FetchMovieExtended(id string) (TvdbMovieResponse, error) {
	return fetchWrapper[TvdbMovieResponse](c, BaseURL+"/movies/"+id+"/extended?meta=translations&short=false")
}
