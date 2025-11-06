package routes

import "github.com/go-chi/chi/v5"

func SetupRoutes(router *chi.Mux) {
	router.Get("/healthz", Healthz)
}
