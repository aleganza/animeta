package mappings

// anime-list-full.json

type AnimeListFullData struct {
	Type               AnimeType
	AniDBID            int
	AniListID          int
	AnimeCountdownID   int
	AnimeNewsNetworkID int
	AnimePlanetID      string
	AniSearchID        int
	IMDbID             []string
	KitsuID            int
	LiveChartID        int
	MALID              int
	TheMovieDBID       TheMovieDBID
	TVDBID             int
	Season             *AnimeListSeason
	EpisodeOffset      *EpisodeOffset
}

type TheMovieDBID struct {
	TV    int
	Movie []int
}

type AnimeListSeason struct {
	TVDB int
	TMDB int
}

type EpisodeOffset struct {
	TVDB int
	TMDB int
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
	AnimeList  []int
	Collection []int
}
