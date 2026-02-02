package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gsn_budget_service/internal"
	"github.com/gsn_budget_service/internal/db/models"
	"github.com/gsn_budget_service/pkg/types"
	"github.com/gsn_budget_service/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type UseHandlers struct {
	AppConns *internal.AppConnections
}

func NewUserControllers(conns *internal.AppConnections) *UseHandlers {
	return &UseHandlers{
		AppConns: conns,
	}
}

func (userController *UseHandlers) NewUser(writer http.ResponseWriter, req *http.Request) {
	var payload types.NewUser

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil {
		log.Error().Err(err).Msg("Failed to parser payload")
		errMsg := err.Error()
		utils.SendErrorResponse(writer, http.StatusBadRequest, "Error validating payload, please check payload schema.", &errMsg)
		return
	}

	if err := utils.Validate.Struct(&payload); err != nil {
		log.Error().Err(err).Msg("New user payload is invalid, check schema")
		errMsg := err.Error()
		utils.SendErrorResponse(writer, http.StatusBadRequest, "Validation failed", &errMsg)
		return
	}

	if payload.Role == nil {
		defaultRole := "user"
		payload.Role = &defaultRole
	}

	_, err := userController.AppConns.DbQueries.GetHouseholdByID(req.Context(), payload.HouseholdId)
	if err != nil {
		log.Error().Err(err).Msgf("Household does not exist, for the following ID => %d", payload.HouseholdId)
		errMsg := err.Error()
		utils.SendErrorResponse(writer, http.StatusBadRequest, "Household does not exist", &errMsg)
		return
	}

	newUser, err := userController.AppConns.DbQueries.CreateNewUser(req.Context(), models.CreateNewUserParams{
		Email:       payload.Email,
		Username:    payload.Username,
		Lastname:    payload.Lastname,
		Gender:      strings.ToUpper(payload.Gender),
		Role:        strings.ToUpper(*payload.Role),
		HouseholdID: pgtype.Int4{Int32: payload.HouseholdId, Valid: true},
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to create new user")
		utils.SendErrorResponse(writer, http.StatusInternalServerError, "Failed to create new user", nil)
		return
	}

	utils.SendJsonResponse(writer, http.StatusCreated, newUser)
}
