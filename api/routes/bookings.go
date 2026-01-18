package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/gsn_budget_service/api/handlers"
	"github.com/gsn_budget_service/internal"
)

func BookingRoutes(appConns *internal.AppConnections) chi.Router {
	r := chi.NewRouter()

	// Create handler instance with app (which contains queries and other dependencies)
	bookingController := handlers.NewBookingController(appConns)

	r.Post("/new-booking", bookingController.CreateBooking)
	r.Get("/book/{householdId}", bookingController.GetListOfBookingsBasedOnHouseHoldId)

	return r
}
