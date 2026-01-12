package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gsn_budget_service/internal"
	"github.com/gsn_budget_service/internal/db/models"
	"github.com/gsn_budget_service/pkg/types"
	"github.com/gsn_budget_service/pkg/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
)

type BookingController struct {
	AppConns *internal.AppConnections
}

func NewBookingController(conns *internal.AppConnections) *BookingController {
	return &BookingController{
		AppConns: conns,
	}
}

func (bookingHandler *BookingController) CreateBooking(w http.ResponseWriter, req *http.Request) {
	var payload types.NewBooking

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&payload); err != nil {
		log.Error().Err(err).Msg("Failed to parser booking payload")
		errMsg := err.Error()
		utils.SendErrorResponse(w, http.StatusBadRequest, "Error validating payload, please check payload schema.", &errMsg)
		return
	}

	var amountNumeric pgtype.Numeric
	if err := amountNumeric.Scan(fmt.Sprintf("%f", payload.Amount)); err != nil {
		log.Error().Err(err).Msg("Fail to parse amount")
		errMsg := err.Error()
		utils.SendErrorResponse(w, http.StatusBadRequest, "Error parsing booking amount", &errMsg)
		return
	}

	newBooking, err := bookingHandler.AppConns.DbQueries.CreateNewBooking(req.Context(), models.CreateNewBookingParams{
		Amount:           amountNumeric,
		UserID:           payload.UserID,
		BookingPlatform:  payload.BookingPlatform,
		FreeCancelBefore: pgtype.Timestamp{Time: payload.FreeCancelBefore, Valid: true},
		BookingStart:     pgtype.Timestamp{Time: payload.BookingStart, Valid: true},
		BookingEnd:       pgtype.Timestamp{Time: payload.BookingEnd, Valid: true},
		Description:      payload.Description,
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to create new booking")
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create new booking", nil)
		return
	}

	utils.SendJsonResponse(w, http.StatusCreated, newBooking)
}
