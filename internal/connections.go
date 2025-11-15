package internal

import (
	"context"

	"github.com/gsn_budget_service/internal/config"
	"github.com/gsn_budget_service/internal/db"
)

// App holds all application dependencies and services
// This can be shared between API handlers, background tasks, cronjobs, etc.
type AppConnections struct {
	Config              *config.Config
	DB                  *db.Db
	Queries             *db.Queries
	ConnectionsShutdown func() error
}

// New creates a new App instance with all dependencies initialized
func New(ctx context.Context, cfg *config.Config) (*AppConnections, error) {
	// Initialize database connection
	dbConnection, err := db.StartDbConnection(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// Create queries from database pool
	queries := db.New(dbConnection.GetPool())

	app := &AppConnections{
		Config:  cfg,
		DB:      dbConnection,
		Queries: queries,
		ConnectionsShutdown: func() error {
			dbConnection.Close()
			return nil
		},
	}

	return app, nil
}

// Close gracefully closes all application resources
func (conns *AppConnections) CloseAppConnections() error {
	if conns.ConnectionsShutdown != nil {
		return conns.ConnectionsShutdown()
	}
	return nil
}
