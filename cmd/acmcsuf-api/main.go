package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/routes"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/services"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

var Version = "dev"

func main() {

	var showVersion = flag.Bool("version", false, "Show version")
	flag.Parse()

	if *showVersion {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		log.Println("Shutting down the server...")
		cancel()
	}()

	db, closer, err := db.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	// Now we init services & gin router, and then start the server
	queries := models.New(db)
	eventsService := services.NewEventsService(queries)
	announcementService := services.NewAnnouncementService(queries)
	router := gin.Default()

	router.SetTrustedProxies([]string{
		"127.0.0.1/32",
	})
	routes.SetupRoutes(router, eventsService, announcementService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	go func() {
		serverAddr := fmt.Sprintf("localhost:%s", port)
		log.Printf("\033[32m Server started on http://%s \033[0m\n", serverAddr)

		if err := router.Run(serverAddr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Server shut down.")
}
