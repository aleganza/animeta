package endpoint_meta

import (
	"animeta/lib/api"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	provider := r.PathValue("provider")
	id := r.PathValue("id")

	if provider == "" {
		api.WriteError(w, http.StatusBadRequest, "Missing provider")
		return
	}

	if id == "" {
		api.WriteError(w, http.StatusBadRequest, "Missing id")
		return
	}

	
}
