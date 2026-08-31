package media

import (
	"animeta/lib/meta/providers/tvdb"

	"golang.org/x/sync/errgroup"
)


func FetchSeries(id string) ([]Media, error) {
	if err := clientGate(); err != nil {
		return nil, err
	}


	var (
		tvdbExtended tvdb.TvdbSeriesExtendedResponse
		tvdbEpisodes tvdb.TvdbSeriesEpisodesResponse
	)

	g := errgroup.Group

	g.Go(func() error {
		var err error
		tvdbExtended, err = client.tvdb.FetchSeriesExtended(id)
		return err
	})

	g.Go(func() error {
		var err error
		tvdbEpisodes, err = client.tvdb.FetchSeriesEpisodes(id)
		return err
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	var series Media
	
	series.Titles = tvdbExtended.Data.Translations.NameTranslations

	for _, episode := range tvdbEpisodes.Data.Episodes {
		
	}

	return nil, nil
}
