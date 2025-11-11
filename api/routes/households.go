package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/gsn_budget_service/api/handlers"
)

func HouseholdRoutes() chi.Router {
	r := chi.NewRouter()

	r.Post("/new", handlers.CreateNewHousehold)

	return r
}
