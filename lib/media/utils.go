package media

import (
	"animeta/lib/data_sources/anime_meta/providers/tvdb"
)

func resolveTvdbArtworks(tvdbArtworks []tvdb.TvdbArtwork) Artworks {
	var artworks Artworks

	for _, artwork := range tvdbArtworks {
		if artwork.Language != "eng" && artwork.Language != "" {
			continue
		}

		artworkType := tvdb.TvdbArtworkTypes[artwork.Type]

		switch artworkType {
		case "Banner":
			artworks = addOnlyFirstArtwork(artworks, "banner", artwork.Image)
		case "ClearLogo":
			artworks = addOnlyFirstArtwork(artworks, "logo", artwork.Image)
		case "Poster":
			if artwork.IncludesText {
				artworks = addOnlyFirstArtwork(artworks, "poster", artwork.Image)
			} else {
				artworks = addOnlyFirstArtwork(artworks, "poster_blank", artwork.Image)
			}
		}
	}

	return artworks
}

func addArtwork(artworks Artworks, artworkType string, url string) Artworks {
	return append(artworks, Artwork{
		Type: artworkType,
		URL:  url,
	})
}

func addOnlyFirstArtwork(artworks Artworks, artworkType string, url string) Artworks {
	for _, a := range artworks {
		if a.Type == artworkType {
			return artworks
		}
	}
	return addArtwork(artworks, artworkType, url)
}
