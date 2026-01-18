package db

import (
	"context"
	"fmt"
	"time"

	"github.com/gsn_budget_service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

// Service wraps the database connection pool and generated queries
type Db struct {
	poolConn *pgxpool.Pool
}

func StartDbConnection(ctx context.Context, cfg *config.Config) (*Db, error) {
	config, err := pgxpool.ParseConfig(cfg.DB_URL)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	// Configure connection pool for optimal performance
	config.MaxConns = 25                      // Maximum number of connections
	config.MinConns = 5                       // Minimum number of connections to keep open
	config.MaxConnLifetime = time.Hour        // Max lifetime of a connection
	config.MaxConnIdleTime = 30 * time.Minute // Max time a connection can be idle
	config.HealthCheckPeriod = time.Minute    // How often to check connection health
	config.ConnConfig.ConnectTimeout = 5 * time.Second

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	// Verify connection
	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	log.Info().Msg("📻 PGX - Connection to DB has been established successfully...")
	return &Db{poolConn: pool}, nil
}

// Close closes the database connection pool
func (db *Db) Close() {
	log.Warn().Msg("Closing DB connection...")
	db.poolConn.Close()
}

// Health checks if the database is healthy
func (db *Db) Health(ctx context.Context) error {
	return db.poolConn.Ping(ctx)
}

// GetPool returns the underlying connection pool (useful for advanced usage)
func (db *Db) GetPool() *pgxpool.Pool {
	return db.poolConn
}

// Stats returns pool statistics
func (db *Db) Stats() *pgxpool.Stat {
	return db.poolConn.Stat()
}
