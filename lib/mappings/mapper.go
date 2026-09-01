package mappings

import (
	"animeta/lib/mappings/providers"
	"fmt"
)

// TODO: should take care of those provider who dont only use ids but seasons too
func FetchMappingsFromProviderId(provider mapping_providers.MappingProvider, id string) (AnimeListFullData, error) {
	if !provider.IsValid() {
		return AnimeListFullData{}, fmt.Errorf(`"provider" argument should be of type MappingProvider`)
	}

	indices, err := FetchIndicesByProviderName(provider)
	if err != nil {
		return AnimeListFullData{}, err
	}

	entry, ok := indices[id]
	if !ok {
		return AnimeListFullData{}, fmt.Errorf(
			"mapping not found for provider %q and id %q",
			provider,
			id,
		)
	}

	if len(entry.AnimeList) == 0 {
		return AnimeListFullData{}, fmt.Errorf(
			"no anime-list mapping found for provider %q and id %q",
			provider,
			id,
		)
	}

	index := entry.AnimeList[0]

	full, err := FetchAnimeListFull()
	if err != nil {
		return AnimeListFullData{}, err
	}

	return full[index], nil
}
