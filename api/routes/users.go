package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/gsn_budget_service/api/handlers"
	"github.com/gsn_budget_service/internal"
)

func UserRoutes(appConns *internal.AppConnections) chi.Router {
	r := chi.NewRouter()
	userHandler := handlers.NewUserControllers(appConns)

	r.Post("/new", userHandler.NewUser)

	return r
}
