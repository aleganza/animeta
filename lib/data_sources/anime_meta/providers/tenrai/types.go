package tenrai

type Client struct {
	ServerKey string
}

type AnimeEpisodesResponse struct {
	Data       []AnimeEpisode `json:"data"`
	Pagination Pagination     `json:"pagination"`
}

type AnimeEpisode struct {
	MalId    int     `json:"mal_id"`
	Title    string  `json:"title"`
	Filler   bool    `json:"filler"`
	Recap    bool    `json:"recap"`
	Aired    *string `json:"aired"`
	Duration int     `json:"duration"`
}

type Pagination struct {
	LastVisiblePage int  `json:"last_visible_page"`
	HasNextPage     bool `json:"has_next_page"`
}
