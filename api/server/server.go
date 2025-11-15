package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gsn_budget_service/api/routes"
	"github.com/gsn_budget_service/internal"
)

func StartServer(appConnections *internal.AppConnections) *http.Server {
	router := chi.NewRouter()

	// Middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)

	// Routes
	routes.SetupRoutes(router, appConnections)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", appConnections.Config.PORT),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &server
}
