package routes

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/gsn_budget_service/internal"
)

func SetupRoutes(router *chi.Mux, appConns *internal.AppConnections) {
	router.Get("/healthz", Healthz)

	router.Route(fmt.Sprintf("/%s", appConns.Config.URL_PREFIX), func(r chi.Router) {
		r.Mount("/households", HouseholdRoutes(appConns))
		r.Mount("/users", UserRoutes(appConns))
		r.Mount("/bookings", BookingRoutes(appConns))
	})
}
