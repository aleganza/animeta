package endpoint_root

import (
	"net/http"
)

const banner = `animeta`

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(banner))
}
