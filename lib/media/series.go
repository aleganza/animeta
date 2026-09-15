package media

import (
	"animeta/lib/data_sources/anime_meta/providers/tvdb"
	"golang.org/x/sync/errgroup"
)

// TODO: title is still anime title, not season title
// how to: in tvdbTranslations go to data.seasons, get id from number, then create new endpoint /seasons/id
func FetchSeries(tvdbId int, tvdbSeasonNumber int) (Media, error) {
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
			Name:     nameTranslation.Name,
			Language: nameTranslation.Language,
		})
	}

	series.Episodes = make([]Episode, 0, len(tvdbEpisodes.Data.Episodes))

	for _, episode := range tvdbEpisodes.Data.Episodes {
		// fetch-all: tvdbSeasonNumber == 0 means the whole series,
		// skipping track-season (0) specials
		if tvdbSeasonNumber > 0 && episode.SeasonNumber != tvdbSeasonNumber {
			continue
		}

		if tvdbSeasonNumber == 0 && episode.SeasonNumber == 0 {
			continue
		}

		num := episode.Number
		if tvdbSeasonNumber == 0 {
			// whole-series numbering matches anidb/tenrai global episode numbers
			num = episode.AbsoluteNumber
		}

		series.Episodes = append(series.Episodes, Episode{
			TvdbId:    episode.ID,
			Number:    num,
			Thumbnail: tvdb.CDNUrl + episode.Image,
			Titles: []Title{
				{
					Name:     episode.Name,
					Language: "eng",
				},
			},
			Overview: episode.Overview,
			Aired:    episode.Aired,
			Runtime:  episode.Runtime,
			Year:     episode.Year,
		})
	}

	series.Artworks = resolveTvdbArtworks(tvdbTranslations.Data.Artworks)
	series.EpisodeCount = len(series.Episodes)

	return series, nil
}
