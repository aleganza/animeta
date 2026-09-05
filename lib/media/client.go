package media

import (
	"animeta/lib/data_sources/anime_meta/providers/tvdb"
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
