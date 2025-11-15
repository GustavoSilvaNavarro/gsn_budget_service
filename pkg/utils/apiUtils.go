package utils

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func SendJsonResponse(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Error().Err(err).Msg("☠️ Failed to reply back, check connectivity...")
	}
}

func SendErrorResponse(w http.ResponseWriter, status int, msg string, details *string) {
	SendJsonResponse(w, status, map[string]any{"message": msg, "details": details})
}
