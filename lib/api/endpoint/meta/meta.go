package endpoint_meta

import (
	"animeta/lib/api"
	"animeta/lib/media"
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	provider := media.Provider(r.PathValue("provider"))
	id := r.PathValue("id")

	// if provider is missing
	if provider == "" {
		api.WriteError(w, http.StatusBadRequest, fmt.Sprintf(
			`missing provider. available providers: %s`,
			media.GetProvidersPretty(),
		))
		return
	}

	// if provider is not valid 
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

	// if id is missing
	if id == "" {
		api.WriteError(w, http.StatusBadRequest, `missing id`)
		return
	}

	data, err := media.Fetch(provider, id)
	if err != nil {
		api.WriteError(w, http.StatusNotFound, err.Error())
		return
	}

	api.WriteSuccess(w, 200, data)
}
