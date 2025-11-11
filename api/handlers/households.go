package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"
)

func CreateNewHousehold(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// One-Liner using a map
	// Note: We use "msg" here to match the desired JSON key
	response := map[string]string{"msg": "success"}

	// Encode and write the map directly
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Error().Msg("Error during calling of the endpoint")
	}
}
