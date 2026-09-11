package endpoint_meta

import (
	"animeta/lib/data_sources/anime_mappings"
	"animeta/lib/media"
	"fmt"
)

var providerAdapters = map[media.Provider]ProviderAdapter{
	media.ProviderAniList: anime_mappings.GetMappingsFromAniListId,
	media.ProviderMAL:     anime_mappings.GetMappingsFromMALId,
}

func GetTvdbData(tvdbId int, tvdbSeasonNumber int) (media.Media, error) {
	isMovie := tvdbSeasonNumber == 0

	var mediaData media.Media

	if isMovie {
		return media.Media{}, fmt.Errorf("movies data are not supported yet")

		// movie, err := media.FetchMovie(tvdbId)
		// if err != nil {
		// 	return media.Media{}, err
		// }

		// mediaData = movie
	} else {
		series, err := media.FetchSeries(tvdbId, tvdbSeasonNumber)
		if err != nil {
			return media.Media{}, err
		}

		mediaData = series
	}

	return mediaData, nil
}
