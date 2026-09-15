package media

import (
	"animeta/lib/data_sources/anime_mappings"
	"animeta/lib/data_sources/anime_meta/providers/anidb"
	"fmt"
	"strconv"
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

	enrichWithAnidb(&mediaData, mappings.AniDBID)

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

func enrichWithAnidb(mediaData *Media, anidbId int) {
	if client.anidb == nil {
		return
	}

	if anidbId == 0 {
		return
	}

	anidbAnime, err := client.anidb.FetchAnime(anidbId)
	if err != nil {
		// keep the tvdb titles as fallback when anidb is unavailable
		return
	}

	mediaData.Titles = mapAnidbTitles(anidbAnime.Titles)

	for i, episode := range mediaData.Episodes {
		if anidbEpisode, ok := findAnidbEpisodeByNumber(anidbAnime.Episodes, episode.Number); ok {
			mediaData.Episodes[i].Titles = mapAnidbEpTitles(anidbEpisode.Titles)
			mediaData.Episodes[i].AnidbId = anidbEpisode.ID
		}
	}
}

func mapAnidbTitles(titles []anidb.AnidbTitle) []Title {
	out := make([]Title, 0, len(titles))

	for _, title := range titles {
		out = append(out, Title{
			Name:     title.Name,
			Language: title.Lang,
		})
	}

	return out
}

func mapAnidbEpTitles(titles []anidb.AnidbEpTitle) []Title {
	out := make([]Title, 0, len(titles))

	for _, title := range titles {
		out = append(out, Title{
			Name:     title.Name,
			Language: title.Lang,
		})
	}

	return out
}

func findAnidbEpisodeByNumber(episodes []anidb.AnidbEpisode, number int) (anidb.AnidbEpisode, bool) {
	target := strconv.Itoa(number)

	for _, episode := range episodes {
		if episode.EpNo == target {
			return episode, true
		}
	}

	return anidb.AnidbEpisode{}, false
}
