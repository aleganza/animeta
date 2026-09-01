package tvdb

// === AUTH ===

type Client struct {
	Token string
}

type LoginResponse struct {
	Status string
	Data   struct {
		Token string
	}
}

// === MEDIA ===

type TvdbResponse[T any] struct {
	Status string
	Data   T
}

type TvdbSeriesTranslationsResponse = TvdbResponse[TvdbSeriesTranslationsData]
type TvdbSeriesEpisodesResponse = TvdbResponse[TvdbSeriesEpisodesData]
type TvdbMovieResponse = TvdbResponse[TvdbMovieData]

type TvdbSeriesTranslationsData struct {
	
}

type TvdbSeriesEpisodesData struct {
	ID                   int
	Name                 string
	Slug                 string
	Image                string
	NameTranslations     []string
	OverviewTranslations []string
	Aliases              []TvdbAlias
	FirstAired           string
	LastAired            string
	NextAired            string
	Score                int
	Status               TvdbStatus
	OriginalCountry      string
	OriginalLanguage     string
	DefaultSeasonType    int
	IsOrderRandomized    bool
	LastUpdated          string
	AverageRuntime       int
	Episodes             []TvdbEpisode
	Overview             string
	Year                 string
}

type TvdbAlias struct {
	Language string
	Name     string
}

type TvdbStatus struct {
	ID          int
	Name        string
	RecordType  string
	KeepUpdated bool
}

type TvdbEpisode struct {
	ID                   int
	SeriesID             int
	Name                 string
	Aired                string
	Runtime              int
	NameTranslations     []string
	Overview             string
	OverviewTranslations []string
	Image                string
	ImageType            int
	IsMovie              int
	Seasons              any
	Number               int
	AbsoluteNumber       int
	SeasonNumber         int
	LastUpdated          string
	FinaleType           *string
	AirsBeforeSeason     int
	AirsBeforeEpisode    int
	Year                 string
}

type TvdbMovieData struct {
	
}
