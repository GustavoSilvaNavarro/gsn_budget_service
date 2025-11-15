package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/gsn_budget_service/api/handlers"
	"github.com/gsn_budget_service/internal"
)

func HouseholdRoutes(appConns *internal.AppConnections) chi.Router {
	r := chi.NewRouter()

	// Create handler instance with app (which contains queries and other dependencies)
	householdHandler := handlers.NewHouseholdHandler(appConns)

	r.Post("/new", householdHandler.CreateNewHousehold)

	return r
}
