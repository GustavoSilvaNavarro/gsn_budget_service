package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gsn_budget_service/internal"
	"github.com/gsn_budget_service/internal/db"
	"github.com/gsn_budget_service/pkg/types"
	"github.com/gsn_budget_service/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

// holds dependencies for household operations
type HouseholdHandler struct {
	appConns *internal.AppConnections
}

// appConns contains all dependencies (queries, db, config, etc.)
func NewHouseholdHandler(conns *internal.AppConnections) *HouseholdHandler {
	return &HouseholdHandler{
		appConns: conns,
	}
}

func (householdController *HouseholdHandler) CreateNewHousehold(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var payload types.CreateHouseholdRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Error().Err(err).Msg("Failed to decode payload...")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Validate request using struct tags
	if err := utils.Validate.Struct(&payload); err != nil {
		log.Error().Err(err).Msg("Validation failed")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Prepare address (convert *string to pgtype.Text)
	var address pgtype.Text
	if payload.Address != nil {
		address = pgtype.Text{String: *payload.Address, Valid: true}
	}

	// Create household in database
	household, err := householdController.appConns.Queries.CreateHousehold(r.Context(), db.CreateHouseholdParams{
		Name:    payload.Name,
		Address: address,
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to create household")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to create household"})
		return
	}

	// Return success response with created household
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(household); err != nil {
		log.Error().Err(err).Msg("Failed to encode response")
	}
}
