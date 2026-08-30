package media

import (
	"animeta/lib/core/threads"
	"animeta/lib/meta/providers/tvdb"
	"fmt"
)

var client *Client

func NewClient(tvdbClient *tvdb.Client) error {
	client = &Client{
		tvdb: tvdbClient,
	}

	return nil
}

func clientGate() error {
	if client == nil {
		return fmt.Errorf("NewClient must be called first")
	}

	return nil
}

func FetchSeries(id string) ([]Media, error) {
	if err := clientGate(); err != nil {
		return nil, err
	}

	results := threads.RunParallel(
    threads.ParallelFetchFunc{
        Name: "tvdb_extended",
        Func: func() (any, error) {
            return client.tvdb.FetchSeriesExtended(id)
        },
    },
    threads.ParallelFetchFunc{
        Name: "tvdb_episodes",
        Func: func() (any, error) {
            return client.tvdb.FetchSeriesEpisodes(id)
        },
    },
)
	var series []Media

	for _, result := range results {
		if result.Name == "tvdb_extended" {
			a := result.Body.
		}
	}

	return "s", nil
}
