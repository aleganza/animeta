package endpoint_meta

import (
	"animeta/lib/api"
	"animeta/lib/data_sources/anime_mappings"
	"animeta/lib/media"
	"fmt"
	"net/http"
)

// TODO: missing movies integration, only works with series

func Handler(w http.ResponseWriter, r *http.Request) {
	provider := media.Provider(r.PathValue("provider"))
	id := r.PathValue("id")

	if provider == "" {
		api.WriteError(w, http.StatusBadRequest, fmt.Sprintf(
			`missing provider. available providers: %s`,
			media.GetProvidersPretty(),
		))
		return
	}

	if !provider.IsValid() {
		api.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf(
				`provider not valid. available providers: %s`,
				media.GetProvidersPretty(),
			),
		)
		return
	}

	if id == "" {
		api.WriteError(w, http.StatusBadRequest, `missing id`)
		return
	}

	// find tvdb id

	var tvdbId int
	var tvdbSeasonNumber int

	if provider == media.ProviderAniList {
		anilist_mappings, err := anime_mappings.GetMappingsFromAniListId(id)

		if err != nil {
			api.WriteError(w, http.StatusNotFound, err.Error())
			return
		}

		tvdbId = anilist_mappings.TVDBID
		tvdbSeasonNumber = anilist_mappings.Season.TVDB
	} else if provider == media.ProviderMAL {
		mal_mappings, err := anime_mappings.GetMappingsFromMALId(id)

		if err != nil {
			api.WriteError(w, http.StatusNotFound, err.Error())
			return
		}

		tvdbId = mal_mappings.TVDBID
		tvdbSeasonNumber = mal_mappings.Season.TVDB
	}

	// fetch series data from tvdb

	data, err := media.FetchSeries(tvdbId, tvdbSeasonNumber)
	if err != nil {
		api.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	api.WriteSuccess(w, 200, data)
}
