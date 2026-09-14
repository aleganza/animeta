package endpoint_meta

import (
	"animeta/lib/data_sources/anime_mappings"
	"animeta/lib/media"
)

type ProviderAdapter func(string) (anime_mappings.AnimeListFullData, error)

type MetaResponse struct {
	media.Media
	Mappings anime_mappings.AnimeListFullData `json:"mappings"`
}
