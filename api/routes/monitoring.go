package routes

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte("OK")); err != nil {
		log.Error().Err(err).Msg("GET /healthz failed")
	}
}
