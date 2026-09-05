package anime_mappings

// anime-list-full.json

type AnimeListFullData struct {
	Type               AnimeType         `json:"type"`
	AniDBID            int               `json:"anidb_id"`
	AniListID          int               `json:"anilist_id"`
	AnimeCountdownID   int               `json:"animecountdown_id"`
	AnimeNewsNetworkID int               `json:"animenewsnetwork_id"`
	AnimePlanetID      string            `json:"anime-planet_id"`
	AniSearchID        int               `json:"anisearch_id"`
	IMDbID             []string          `json:"imdb_id"`
	KitsuID            int               `json:"kitsu_id"`
	LiveChartID        int               `json:"livechart_id"`
	MALID              int               `json:"mal_id"`
	SimklID            int               `json:"simkl_id"`
	TheMovieDBID       TheMovieDBID      `json:"themoviedb_id"`
	TVDBID             int               `json:"tvdb_id"`
	Season             *AnimeListSeason  `json:"season"`
	EpisodeOffset      *EpisodeOffset    `json:"episode_offset"`
}

type TheMovieDBID struct {
	TV    int   `json:"tv"`
	Movie []int `json:"movie"`
}

type AnimeListSeason struct {
	TVDB int `json:"tvdb"`
	TMDB int `json:"tmdb"`
}

type EpisodeOffset struct {
	TVDB int `json:"tvdb"`
	TMDB int `json:"tmdb"`
}

// generic

type AnimeType string

const (
	AnimeTypeTV      AnimeType = "TV"
	AnimeTypeOVA     AnimeType = "OVA"
	AnimeTypeSpecial AnimeType = "SPECIAL"
	AnimeTypeMovie   AnimeType = "MOVIE"
)

// *_index.json

type AnimeListIndex map[string]AnimeListIndexEntry

type AnimeListIndexEntry struct {
	AnimeList  []int `json:"anime-list"`
	Collection []int `json:"collection"`
}

// mapping providers

type MappingProvider string

const (
	MappingProviderAniDB            MappingProvider = "anidb"
	MappingProviderAniList          MappingProvider = "anilist"
	MappingProviderAnimePlanet      MappingProvider = "anime-planet"
	MappingProviderAnimeCountdown   MappingProvider = "animecountdown"
	MappingProviderAnimeNewsNetwork MappingProvider = "animenewsnetwork"
	MappingProviderAniSearch        MappingProvider = "anisearch"
	MappingProviderIMDb             MappingProvider = "imdb"
	MappingProviderKitsu            MappingProvider = "kitsu"
	MappingProviderLiveChart        MappingProvider = "livechart"
	MappingProviderMAL              MappingProvider = "mal"
	MappingProviderSimkl            MappingProvider = "simkl"
	MappingProviderTMDB             MappingProvider = "themoviedb"
	MappingProviderTVDB             MappingProvider = "tvdb"
)
