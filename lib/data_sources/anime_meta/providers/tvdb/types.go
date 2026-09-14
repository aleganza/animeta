package tvdb

// === AUTH ===

type Client struct {
	Token string
}

type LoginResponse struct {
	Status string `json:"status"`
	Data   struct {
		Token string `json:"token"`
	} `json:"data"`
}

// === MEDIA ===

type TvdbResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type TvdbSeriesTranslationsResponse = TvdbResponse[TvdbSeriesTranslationsData]
type TvdbSeriesEpisodesResponse = TvdbResponse[TvdbSeriesEpisodesData]
type TvdbMovieResponse = TvdbResponse[TvdbMovieData]
type TvdbRemoteIdResponse = TvdbResponse[TvdbRemoteIdData]

type TvdbSeriesTranslationsData struct {
	ID                   int              `json:"id"`
	Name                 string           `json:"name"`
	Slug                 string           `json:"slug"`
	Image                string           `json:"image"`
	NameTranslations     []string         `json:"nameTranslations"`
	OverviewTranslations []string         `json:"overviewTranslations"`
	Aliases              []TvdbAlias      `json:"aliases"`
	FirstAired           string           `json:"firstAired"`
	LastAired            string           `json:"lastAired"`
	NextAired            string           `json:"nextAired"`
	Score                int              `json:"score"`
	Status               TvdbStatus       `json:"status"`
	OriginalCountry      string           `json:"originalCountry"`
	OriginalLanguage     string           `json:"originalLanguage"`
	DefaultSeasonType    int              `json:"defaultSeasonType"`
	IsOrderRandomized    bool             `json:"isOrderRandomized"`
	LastUpdated          string           `json:"lastUpdated"`
	AverageRuntime       int              `json:"averageRuntime"`
	Episodes             []TvdbEpisode    `json:"episodes"`
	Overview             string           `json:"overview"`
	Year                 string           `json:"year"`
	Translations         TvdbTranslations `json:"translations"`
	Artworks             []TvdbArtwork    `json:"artworks"`
}

type TvdbSeriesEpisodesData struct {
	ID                   int           `json:"id"`
	Name                 string        `json:"name"`
	Slug                 string        `json:"slug"`
	Image                string        `json:"image"`
	NameTranslations     []string      `json:"nameTranslations"`
	OverviewTranslations []string      `json:"overviewTranslations"`
	Aliases              []TvdbAlias   `json:"aliases"`
	FirstAired           string        `json:"firstAired"`
	LastAired            string        `json:"lastAired"`
	NextAired            string        `json:"nextAired"`
	Score                int           `json:"score"`
	Status               TvdbStatus    `json:"status"`
	OriginalCountry      string        `json:"originalCountry"`
	OriginalLanguage     string        `json:"originalLanguage"`
	DefaultSeasonType    int           `json:"defaultSeasonType"`
	IsOrderRandomized    bool          `json:"isOrderRandomized"`
	LastUpdated          string        `json:"lastUpdated"`
	AverageRuntime       int           `json:"averageRuntime"`
	Episodes             []TvdbEpisode `json:"episodes"`
	Overview             string        `json:"overview"`
	Year                 string        `json:"year"`
}

type TvdbMovieData struct {
	ID           int              `json:"id"`
	Image        string           `json:"image"`
	Runtime      int              `json:"runtime"`
	Year         string           `json:"year"`
	Translations TvdbTranslations `json:"translations"`
	Artworks     []TvdbArtwork    `json:"artworks"`
	FirstRelease TvdbMovieRelease `json:"first_release"`
}

type TvdbRemoteIdData = []TvdbRemoteIdResult
type TvdbRemoteIdResult struct {
	Series *any `json:"series"`
	Movie  *struct {
		Id int `json:"id"`
	} `json:"movie"`
	Episode *any `json:"episode"`
	// and more but not needed now
}

type TvdbTranslations struct {
	NameTranslations     []TvdbNameTranslation     `json:"nameTranslations"`
	OverviewTranslations []TvdbOverviewTranslation `json:"overviewTranslations"`
	Aliases              []string                  `json:"aliases"`
}

type TvdbNameTranslation struct {
	Name      string `json:"name"`
	Language  string `json:"language"`
	IsPrimary bool   `json:"isPrimary,omitempty"`
	IsAlias   bool   `json:"isAlias,omitempty"`
}

type TvdbOverviewTranslation struct {
	Overview  string `json:"overview"`
	Language  string `json:"language"`
	IsPrimary bool   `json:"isPrimary,omitempty"`
}

type TvdbAlias struct {
	Language string `json:"language"`
	Name     string `json:"name"`
}

type TvdbStatus struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	RecordType  string `json:"recordType"`
	KeepUpdated bool   `json:"keepUpdated"`
}

type TvdbEpisode struct {
	ID                   int      `json:"id"`
	SeriesID             int      `json:"seriesId"`
	Name                 string   `json:"name"`
	Aired                string   `json:"aired"`
	Runtime              int      `json:"runtime"`
	NameTranslations     []string `json:"nameTranslations"`
	Overview             string   `json:"overview"`
	OverviewTranslations []string `json:"overviewTranslations"`
	Image                string   `json:"image"`
	ImageType            int      `json:"imageType"`
	IsMovie              int      `json:"isMovie"`
	Seasons              any      `json:"seasons"`
	Number               int      `json:"number"`
	AbsoluteNumber       int      `json:"absoluteNumber"`
	SeasonNumber         int      `json:"seasonNumber"`
	LastUpdated          string   `json:"lastUpdated"`
	FinaleType           *string  `json:"finaleType"`
	AirsBeforeSeason     int      `json:"airsBeforeSeason"`
	AirsBeforeEpisode    int      `json:"airsBeforeEpisode"`
	Year                 string   `json:"year"`
}

type TvdbArtwork struct {
	Id           int    `json:"id"`
	Image        string `json:"image"`
	Thumbnail    string `json:"thumbnail"`
	Language     string `json:"language"`
	Type         int    `json:"type"`
	Score        int    `json:"score"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	IncludesText bool   `json:"includesText"`
}

type TvdbMovieRelease struct {
	Country string  `json:"country"`
	Date    string  `json:"date"`
	Detail  *string `json:"detail"`
}

var TvdbArtworkTypes = map[int]string{
	1:  "Banner",
	2:  "Poster",
	3:  "Background",
	5:  "Icon",
	6:  "Banner",
	7:  "Poster",
	8:  "Background",
	10: "Icon",
	11: "16:9 Screencap",
	12: "4:3 Screencap",
	13: "Photo",
	14: "Poster",
	15: "Background",
	16: "Banner",
	18: "Icon",
	19: "Icon",
	20: "Cinemagraph",
	21: "Cinemagraph",
	22: "ClearArt",
	23: "ClearLogo",
	24: "ClearArt",
	25: "ClearLogo",
	26: "Icon",
	27: "Poster",
}
