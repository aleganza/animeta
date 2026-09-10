package media

import "animeta/lib/data_sources/anime_meta/providers/tvdb"

// client

type Client struct {
	tvdb *tvdb.Client
}

// media providers

type Provider string

const (
	ProviderAniList Provider = "anilist"
	ProviderMAL     Provider = "mal"
)

// needed to list them
var Providers = []Provider{
	ProviderAniList,
	ProviderMAL,
}

// media

type Media struct {
	Titles   []Title            `json:"titles"`
	Episodes []Episode `json:"episodes"`
	// EpisodeCount int
	// SpecialCount int
	// Images      []Image
	// Mappings    Mappings
}

type Title struct {
	Name      string `json:"name"`
	Language  string `json:"language"`
	IsPrimary bool   `json:"isPrimary"`
	IsAlias   bool   `json:"isAlias"`
}

type Episode struct {
	TvdbId int `json:"tvdbId"`

	SeasonNumber int `json:"seasonNumber"`
	// Girls und Panzer has 5.5
	Number int `json:"number"`

	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
	Overview  string `json:"overview"`
	Aired     string `json:"aired"`
	Runtime   int    `json:"runtime"`
	Year      string `json:"year"`
}

type Image struct {
	CoverType string `json:"coverType"`
	URL       string `json:"url"`
}

type Mappings struct {
	AnimePlanetID string `json:"animePlanetId"`
	KitsuID       int    `json:"kitsuId"`
	MalID         int    `json:"malId"`
	Type          string `json:"type"`
	AnilistID     int    `json:"anilistId"`
	AnisearchID   int    `json:"anisearchId"`
	AnidbID       int    `json:"anidbId"`
	NotifymoeID   *int   `json:"notifymoeId"`
	LivechartID   int    `json:"livechartId"`
	TheTVDBID     int    `json:"theTvdbId"`
	IMDBID        string `json:"imdbId"`
	TheMovieDBID  string `json:"theMovieDbId"`
}
