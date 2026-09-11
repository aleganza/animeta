package endpoint_meta

import "animeta/lib/data_sources/anime_mappings"

type ProviderAdapter func(string) (anime_mappings.AnimeListFullData, error)
