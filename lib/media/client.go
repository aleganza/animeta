package media

import (
	"animeta/lib/data_sources/anime_meta/providers/anidb"
	"animeta/lib/data_sources/anime_meta/providers/tvdb"
	"fmt"
)

var client *Client

func NewClient(tvdbClient *tvdb.Client, anidbClient *anidb.Client) error {
	client = &Client{
		tvdb:  tvdbClient,
		anidb: anidbClient,
	}

	return nil
}

func clientGate() error {
	if client == nil {
		return fmt.Errorf("NewClient must be called first")
	}

	return nil
}
