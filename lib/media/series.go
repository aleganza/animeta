package media

import (
	"animeta/lib/data_sources/anime_meta/providers/tvdb"
	"golang.org/x/sync/errgroup"
)

func FetchSeries(tvdbId string, tvdbSeasonId string) (Media, error) {
	if err := clientGate(); err != nil {
		return Media{}, err
	}

	var (
		tvdbTranslations tvdb.TvdbSeriesTranslationsResponse
		tvdbEpisodes     tvdb.TvdbSeriesEpisodesResponse
	)

	var g errgroup.Group

	g.Go(func() error {
		var err error
		tvdbTranslations, err = client.tvdb.FetchSeriesTranslations(tvdbId)
		return err
	})

	g.Go(func() error {
		var err error
		tvdbEpisodes, err = client.tvdb.FetchSeriesEpisodes(tvdbId)
		return err
	})

	if err := g.Wait(); err != nil {
		return Media{}, err
	}

	var series Media

	for _, nameTranslation := range tvdbTranslations.Data.Translations.NameTranslations {
		series.Titles = append(series.Titles, Title{
			Name:      nameTranslation.Name,
			Language:  nameTranslation.Language,
			IsPrimary: nameTranslation.IsPrimary,
			IsAlias:   nameTranslation.IsAlias,
		})
	}

	series.Episodes = make([]Episode, 0, len(tvdbEpisodes.Data.Episodes))

	for _, episode := range tvdbEpisodes.Data.Episodes {
		series.Episodes = append(series.Episodes, Episode{
			TvdbId:       episode.ID,
			SeasonNumber: episode.SeasonNumber,
			Number:       episode.Number,
			Thumbnail:    episode.Image,
			Title:        episode.Name,
			Overview:     episode.Overview,
			Aired:        episode.Aired,
			Runtime:      episode.Runtime,
			Year:         episode.Year,
		})
	}

	return series, nil
}
