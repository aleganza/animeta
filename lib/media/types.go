package media

import "animeta/lib/meta/providers/tvdb"

type Client struct {
	tvdb *tvdb.Client
}

type Media struct {
	Titles Titles
	// Episodes    map[string]Episode
	// EpisodeCount int
	// SpecialCount int
	// Images      []Image
	// Mappings    Mappings
}

type Titles map[string]string

type Episode struct {
	SeasonNumber          int
	EpisodeNumber         int
	AbsoluteEpisodeNumber int

	Title map[string]string

	AirDate    string
	AirDateUTC string

	Runtime  int
	Overview string
	Image    string

	Episode string

	AnidbEID int
	Length   int
	Airdate  string
	Rating   string
	Summary  string

	FinaleType string
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
