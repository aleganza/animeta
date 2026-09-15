package anidb

type Client struct {
	Name    string
	Version int
}

type AnidbAnime struct {
	ID       int            `xml:"id,attr"`
	Titles   []AnidbTitle   `xml:"titles>title"`
	Episodes []AnidbEpisode `xml:"episodes>episode"`
}

type AnidbTitle struct {
	Type string `xml:"type,attr"`
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
	Name string `xml:",chardata"`
}

type AnidbEpisode struct {
	ID     int            `xml:"id,attr"`
	EpNo   string         `xml:"epno"`
	Length int            `xml:"length"`
	Titles []AnidbEpTitle `xml:"title"`
}

type AnidbEpTitle struct {
	Lang string `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
	Name string `xml:",chardata"`
}
