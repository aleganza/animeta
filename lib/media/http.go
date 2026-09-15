package media

import (
	"animeta/lib/data_sources/anime_mappings"
	"fmt"
)

type ProviderAdapter func(string) (anime_mappings.AnimeListFullData, error)

var providerAdapters = map[Provider]ProviderAdapter{
	ProviderAniList: anime_mappings.GetMappingsFromAniListId,
	ProviderMAL:     anime_mappings.GetMappingsFromMALId,
}

type FetchResult struct {
	Media
	Mappings anime_mappings.AnimeListFullData `json:"mappings"`
}

func Fetch(provider Provider, id string) (FetchResult, error) {
	adapter, ok := providerAdapters[provider]
	if !ok {
		return FetchResult{}, fmt.Errorf(`provider not supported. available providers: %s`, GetProvidersPretty())
	}

	mappings, err := adapter(id)
	if err != nil {
		return FetchResult{}, err
	}

	tvdbId := mappings.TVDBID

	// standalone movies have no series/season mapping, so Season is nil.
	tvdbSeasonNumber := 0
	if mappings.Season != nil {
		tvdbSeasonNumber = mappings.Season.TVDB
	}

	mediaData, err := fetchTvdbData(tvdbId, tvdbSeasonNumber, mappings.IMDbID)
	if err != nil {
		return FetchResult{}, err
	}

	return FetchResult{
		Media:    mediaData,
		Mappings: mappings,
	}, nil
}

// IMDb ids are used to resolve the TVDB movie id, since mappings only hold the series id.
func fetchTvdbData(tvdbId int, tvdbSeasonNumber int, imdbIds []string) (Media, error) {
	isMovie := tvdbSeasonNumber == 0

	if isMovie {
		// map tvdb "season 0 series id" to tvdb movie id
		tvdbMovieID, err := ResolveMovieId(imdbIds)
		if err != nil {
			return Media{}, err
		}

		movie, err := FetchMovie(tvdbMovieID)
		if err != nil {
			return Media{}, err
		}

		return movie, nil
	}

	series, err := FetchSeries(tvdbId, tvdbSeasonNumber)
	if err != nil {
		return Media{}, err
	}

	return series, nil
}
