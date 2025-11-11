package routes

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/gsn_budget_service/internal/config"
)

func SetupRoutes(router *chi.Mux) {
	router.Get("/healthz", Healthz)

	router.Route(fmt.Sprintf("/%s", config.Cfg.URL_PREFIX), func(r chi.Router) {
		r.Mount("/households", HouseholdRoutes())
	})
}
