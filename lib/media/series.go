package media

import (
	"animeta/lib/data_sources/anime_meta/providers/tvdb"

	"golang.org/x/sync/errgroup"
)

func FetchSeries(id string) ([]Media, error) {
	if err := clientGate(); err != nil {
		return nil, err
	}

	var (
		tvdbTranslations tvdb.TvdbSeriesTranslationsResponse
		tvdbEpisodes     tvdb.TvdbSeriesEpisodesResponse
	)

	var g errgroup.Group

	g.Go(func() error {
		var err error
		tvdbTranslations, err = client.tvdb.FetchSeriesTranslations(id)
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

	// series.Titles = tvdbTranslations.Data.Translations.NameTranslations

	// for _, episode := range tvdbEpisodes.Data.Episodes {

	// }

	return nil, nil
}
