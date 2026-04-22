// This file (server.go) contains server initialization logic that's called by main.go

package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
	mw "github.com/acmcsufoss/api.acmcsuf.com/internal/api/middleware"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/routes"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"
)

// Run initializes the database, services, and router, then starts the server.
// It waits for the context to be canceled to initiate a graceful shutdown.
func Run(ctx context.Context) {
	cfg := config.Load()

	db, closer, err := NewDB(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	// Apply db migrations
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("could not create sqlite3 driver: %v\n", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://sql/migrations",
		"sqlite3",
		driver,
	)
	if err != nil {
		log.Fatalf("could not create migration instance: %v\n", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("could not run db migrations: %v\n", err)
	}

	// Now we init services & gin router, and then start the server
	queries := dbmodels.New(db)
	eventsService := services.NewEventsService(queries)
	announcementService := services.NewAnnouncementService(queries)
	boardService := services.NewBoardService(queries, db)
	router := gin.Default()
	router.Use(mw.Cors(), mw.Ratelimiter())

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
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// This is a blocking call that prevents the function from finishing until the signal
	// is received.
	<-ctx.Done()
	log.Println("\x1b[32mServer shut down.\x1b[0m")
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
