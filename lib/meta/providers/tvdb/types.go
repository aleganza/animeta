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

type Alias struct {
	Language string
	Name     string
}

type MediaStatus struct {
	Id          uint
	Name        string
	RecordType  string
	KeepUpdated bool
}

type Translation struct {
	Language  string
	Name      string
	Overview  string
	IsPrimary bool
	IsAlias   bool
	Aliases   []string
}

type Translations struct {
	NameTranslations     []Translation
	OverviewTranslations []Translation
}

type Artwork struct {
	Id           uint
	Image        string
	Thumbnail    string
	Language     string
	Type         uint
	Score        float64
	Width        uint
	Height       uint
	IncludesText bool
}

type Episode struct {
	Id             uint
	SeriesId       uint
	Name           string
	Aired          string
	Runtime        *uint
	Overview       string
	Image          string
	Number         uint
	SeasonNumber   uint
	AbsoluteNumber uint
	Year           string
}

type TvdbSeriesExtendedResponse struct {
	Status string
	Data   struct {
		Id           uint
		Name         string
		Slug         string
		Image        string
		Translations Translations
		Aliases      []Alias

		FirstAired string
		LastAired  string
		NextAired  string

		Score uint

		Status MediaStatus

		OriginalCountry  string
		OriginalLanguage string

		AverageRuntime uint

		Artworks []Artwork
		Episodes []Episode

		Overview string
		Year     string
	}
}

type TvdbSeriesEpisodesResponse struct {
	Status string
	Data   struct {
		Series struct {
			Id   uint
			Name string
		}
		Episodes []Episode
	}
}

type TvdbMovieResponse struct {
	Status string
	Data   struct {
		Id           uint
		Name         string
		Slug         string
		Image        string
		Translations Translations
		Aliases      []Alias

		Score   uint
		Runtime uint

		Status MediaStatus

		LastUpdated string
		Year        string

		Trailers []struct {
			Id       uint
			Name     string
			URL      string
			Language string
			Runtime  uint
		}

		Genres []struct {
			Id   uint
			Name string
			Slug string
		}

		Releases []struct {
			Country string
			Date    string
			Detail  *string
		}

		OriginalCountry  string
		OriginalLanguage string

		Studios []struct {
			Id   uint
			Name string
		}

		Artworks []Artwork

		Budget    string
		BoxOffice string

		Overview string
	}
}