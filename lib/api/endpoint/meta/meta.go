package endpoint_meta

import (
	"animeta/lib/api"
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

	// find tvdb identifiers
	adapter, ok := providerAdapters[provider]

	if !ok {
		api.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf(
				`provider not supported. available providers: %s`,
				media.GetProvidersPretty(),
			),
		)
		return
	}

	mapping, err := adapter(id)

	if err != nil {
		api.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	tvdbId := mapping.TVDBID
	
	// standalone movies have no series/season mapping, so Season is nil.
	tvdbSeasonNumber := 0
	if mapping.Season != nil {
			tvdbSeasonNumber = mapping.Season.TVDB
	}

	data, err := GetTvdbData(tvdbId, tvdbSeasonNumber, mapping.IMDbID)
	if err != nil {
		api.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	api.WriteSuccess(w, 200, data)
}
