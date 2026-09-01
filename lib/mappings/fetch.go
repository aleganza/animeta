package mappings

import (
	"animeta/lib/core/fetch"
	mapping_providers "animeta/lib/mappings/providers"
	"fmt"
)

// not convinced about filename

func fetchWrapper[T any](path string) (T, error) {
	resp, err := fetch.ExecuteRequest("GET", BaseUrl+path, nil)
	if err != nil {
		var zero T
		return zero, err
	}

	var data T

	if err := fetch.ExtractResponseJsonBody(resp, &data); err != nil {
		var zero T
		return zero, err
	}

	return data, nil
}

func FetchAnimeListFull() ([]AnimeListFullData, error) {
	return fetchWrapper[[]AnimeListFullData]("anime-list-full.json")
}

func FetchIndicesByProviderName(provider mapping_providers.MappingProvider) (AnimeListIndex, error) {
	if !provider.IsValid() {
		return AnimeListIndex{}, fmt.Errorf(`"provider" argument should be of type MappingProvider`)
	}

	return fetchWrapper[AnimeListIndex]("indices/" + string(provider) + "_index.json")
}
