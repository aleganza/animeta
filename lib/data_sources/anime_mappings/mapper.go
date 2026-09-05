package anime_mappings

import (
	"fmt"
)

func GetMappingsFromProviderId(provider MappingProvider, id string) (AnimeListFullData, error) {
	if !provider.IsValid() {
		return AnimeListFullData{}, fmt.Errorf(`"%s" provider should be of type MappingProvider`, provider)
	}

	if !provider.isMappingsRetrievalHandled() {
		return AnimeListFullData{}, fmt.Errorf(`"%s" provider not handled`, provider)
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
