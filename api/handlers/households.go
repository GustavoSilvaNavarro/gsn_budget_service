package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gsn_budget_service/internal"
	"github.com/gsn_budget_service/internal/db/models"
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
	var payload types.CreateHouseholdRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil {
		log.Error().Err(err).Msg("Failed to decode payload...")
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid payload received, check payload schema.", nil)
		return
	}

	// Validate request using struct tags
	if err := utils.Validate.Struct(&payload); err != nil {
		log.Error().Err(err).Msg("😩 Payload validation failed...")
		errMsg := err.Error()
		utils.SendErrorResponse(w, http.StatusBadRequest, "Error validating payload, check payload schema.", &errMsg)
		return
	}

	// Convert optional *string to pgtype.Text (required for nullable database column)
	address := pgtype.Text{Valid: false} // Default to NULL
	if payload.Address != nil {
		address = pgtype.Text{String: *payload.Address, Valid: true}
	}

	// Create household in database
	newHousehold, err := householdController.appConns.DbQueries.CreateHousehold(r.Context(), models.CreateHouseholdParams{
		Name:    payload.Name,
		Address: address,
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to create household")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create new household in db", nil)
		return
	}

	// Return success response with created household
	utils.SendJsonResponse(w, http.StatusCreated, newHousehold)
}

func (householdController *HouseholdHandler) GetHousehold(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "householdId")

	// 2. Convert string to integer
	householdId, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid household ID", nil)
		return
	}

	id := int32(householdId)
	household, err := householdController.appConns.DbQueries.GetHouseholdByID(r.Context(), id)

	if err != nil {
		log.Error().Err(err).Msgf("Failed to retrieve household with ID: %d", id)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to retrieve household from db", nil)
		return
	}

	// Return success response with created household
	utils.SendJsonResponse(w, http.StatusCreated, household)
}
