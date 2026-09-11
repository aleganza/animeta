package media

func FetchMovie(tvdbId int) (Media, error) {
	if err := clientGate(); err != nil {
		return Media{}, err
	}

	tvdbMovie, err := client.tvdb.FetchMovieExtended(tvdbId)
	if err != nil {
		return Media{}, err
	}

	var media Media

	for _, nameTranslation := range tvdbMovie.Data.Translations.NameTranslations {
		media.Titles = append(media.Titles, Title{
			Name:     nameTranslation.Name,
			Language: nameTranslation.Language,
		})
	}

	title := ""
	overview := ""

	for _, nameTranslation := range tvdbMovie.Data.Translations.NameTranslations {
		if nameTranslation.Language == "eng" {
			title = nameTranslation.Name
			break
		}
	}

	for _, overviewTranslation := range tvdbMovie.Data.Translations.OverviewTranslations {
		if overviewTranslation.Language == "eng" {
			overview = overviewTranslation.Overview
			break
		}
	}

	media.Episodes = []Episode{
		{
			TvdbId:       tvdbMovie.Data.ID,
			SeasonNumber: 0,
			Number:       1,
			Thumbnail:    tvdbMovie.Data.Image,
			Title:        title,
			Overview:     overview,
			Aired:        tvdbMovie.Data.FirstRelease.Date,
			Runtime:      tvdbMovie.Data.Runtime,
			Year:         tvdbMovie.Data.Year,
		},
	}

	return media, nil
}
