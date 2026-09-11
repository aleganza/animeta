package endpoint_meta

import (
	"animeta/lib/data_sources/anime_mappings"
	"animeta/lib/media"
)

var providerAdapters = map[media.Provider]ProviderAdapter{
	media.ProviderAniList: anime_mappings.GetMappingsFromAniListId,
	media.ProviderMAL:     anime_mappings.GetMappingsFromMALId,
}

// IMDb ids are used to resolve the TVDB movie id, since mappings only hold the series id.
func GetTvdbData(tvdbId int, tvdbSeasonNumber int, imdbIs []string) (media.Media, error) {
	isMovie := tvdbSeasonNumber == 0

	var mediaData media.Media

	if isMovie {
		// map tvdb "season 0 series id" to tvdb movie id
		tvdbMovieID, err := media.ResolveMovieId(imdbIs)

		movie, err := media.FetchMovie(tvdbMovieID)
		if err != nil {
			return media.Media{}, err
		}

		mediaData = movie
	} else {
		series, err := media.FetchSeries(tvdbId, tvdbSeasonNumber)
		if err != nil {
			return media.Media{}, err
		}

		mediaData = series
	}

	return mediaData, nil
}
