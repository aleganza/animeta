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

	if provider == "" {
		api.WriteError(w, http.StatusBadRequest, "missing provider")
		return
	}

	if !provider.IsValid() {
		api.WriteError(
			w,
			http.StatusBadRequest,
			fmt.Sprintf(
				"provider not valid. available providers: %s",
				media.GetProvidersPretty(),
			),
		)
		return
	}

	if id == "" {
		api.WriteError(w, http.StatusBadRequest, "missing id")
		return
	}

	api.WriteSuccess(w, 200, "yep")
}
