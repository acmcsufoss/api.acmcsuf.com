// This file (server.go) contains server initialization logic that's called by main.go

package api

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/logging"
	mw "github.com/acmcsufoss/api.acmcsuf.com/internal/api/middleware"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/routes"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
)

// Run initializes the database, services, and router, then starts the server.
// It waits for the context to be canceled to initiate a graceful shutdown.
func Run(ctx context.Context, logger *slog.Logger) {
	cfg := config.Load()

	db, closer, err := NewDB(ctx, cfg.DatabaseURL)
	if err != nil {
		logging.Fatal(logger, "could not open database", "error", err)
	}
	defer closer()

	// Apply db migrations
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		logging.Fatal(logger, "could not create sqlite3 driver", "error", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://sql/migrations",
		"sqlite3",
		driver,
	)
	if err != nil {
		logging.Fatal(logger, "could not create migration instance", "error", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logging.Fatal(logger, "could not run db migrations", "error", err)
	}

	// Now we init services & gin router, and then start the server
	queries := dbmodels.New(db)
	eventsService := services.NewEventsService(queries)
	announcementService := services.NewAnnouncementService(queries)
	boardService := services.NewBoardService(queries, db)
	router := gin.New()
	router.Use(logging.RequestLogger(logger), gin.Recovery(), mw.Cors(), mw.Ratelimiter())

	router.SetTrustedProxies(cfg.TrustedProxies)
	routes.SetupRoot(router)
	routes.SetupV1(router, eventsService, announcementService, boardService)

	go func() {
		serverAddr := ":" + cfg.Port
		if cfg.Env == "development" {
			// this binds the server to the loopback interface in dev mode for security reasons
			serverAddr = "localhost:" + cfg.Port
		}

		if err := router.Run(serverAddr); err != nil {
			logging.Fatal(logger, "failed to start server", "error", err)
		}
	}()

	// This is a blocking call that prevents the function from finishing until the signal
	// is received.
	<-ctx.Done()
	logger.Info("server shut down")
}

func NewDB(ctx context.Context, url string) (*sql.DB, func(), error) {
	db, err := sql.Open("sqlite3", url)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening SQLite database: %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return db, func() {
		db.Close()
	}, nil
}
