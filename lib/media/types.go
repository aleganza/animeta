package media

import "animeta/lib/data_sources/anime_meta/providers/tvdb"

type Client struct {
	tvdb *tvdb.Client
}

type Media struct {
	// Titles   ?????????
	Episodes map[string]Episode
	// EpisodeCount int
	// SpecialCount int
	// Images      []Image
	// Mappings    Mappings
}

// type Titles map[string]string

// type Title struct {
// 	Name      string
// 	Language  string
// 	IsPrimary bool
// 	IsAlias   bool
// }

type Episode struct {
	TvdbId int

	SeasonNumber int
	// Girls und Panzer has 5.5
	Number       int

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
