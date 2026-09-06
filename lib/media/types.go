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
	Titles   []Title
	Episodes map[string]Episode
	// EpisodeCount int
	// SpecialCount int
	// Images      []Image
	// Mappings    Mappings
}

type Title struct {
	Name      string
	Language  string
	IsPrimary bool
	IsAlias   bool
}

type Episode struct {
	TvdbId int

	SeasonNumber int
	// Girls und Panzer has 5.5
	Number int

	Thumbnail string
	Title     string
	Overview  string
	Aired     string
	Runtime   int
	Year      string
}

type Image struct {
	CoverType string
	URL       string
}

type Mappings struct {
	AnimePlanetID string
	KitsuID       int
	MalID         int
	Type          string

	AnilistID   int
	AnisearchID int
	AnidbID     int

	NotifymoeID *int

	LivechartID int

	TheTVDBID    int
	IMDBID       string
	TheMovieDBID string
}
