package tvdb

import (
	"animeta/lib/core/fetch"
)

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

func (c *Client) FetchSeriesExtended(id string) (TvdbSeriesExtendedResponse, error) {
	return fetchWrapper[TvdbSeriesExtendedData](c, BaseURL+"/series/"+id+"/extended?meta=translations&short=false")
}

func (c *Client) FetchSeriesEpisodes(id string) (TvdbSeriesEpisodesResponse, error) {
	return fetchWrapper[TvdbSeriesEpisodesData](c, BaseURL+"/series/"+id+"/episodes/default/eng")
}

func (c *Client) FetchMovieExtended(id string) (TvdbMovieResponse, error) {
	return fetchWrapper[TvdbMovieData](c, BaseURL+"/movies/"+id+"/extended?meta=translations&short=false")
}
