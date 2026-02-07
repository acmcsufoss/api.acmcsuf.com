// This file (server.go) contains server initialization logic that's called by main.go

package api

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	mw "github.com/acmcsufoss/api.acmcsuf.com/internal/api/middleware"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/routes"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
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
		serverAddr := fmt.Sprintf("localhost:%s", cfg.Port)
		log.Printf("\x1b[32mServer started on http://%s\x1b[0m", serverAddr)

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
	db, err := sql.Open("sqlite", url)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening SQLite database: %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, nil, fmt.Errorf("error connecting to database: %v", err)
	}

	// schemaBytes, err := os.ReadFile("internal/db/sql/schemas/schema.sql")
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("error reading schema file: %v", err)
	// }
	//
	// if _, err := db.ExecContext(ctx, string(schemaBytes)); err != nil {
	// 	return nil, nil, fmt.Errorf("error initializing db schema: %v", err)
	//
	// }

	return db, func() {
		db.Close()
	}, nil

}
