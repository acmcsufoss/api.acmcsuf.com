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
	// =================== Command Line Arg Parsing ===================
	var showVersion = flag.Bool("version", false, "Show version")
	flag.Parse()

	if *showVersion {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}
	// ================================================================


	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		log.Println("Shutting down the server...")
		// when cancel is called, it sends a "done" signal to ctx
		cancel()
	}()

	// =================== Inialize database ===================
	db, closer, err := db.New(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer closer()

	// =================== Inialize all services ===================
	queries := models.New(db)
	eventsService := services.NewEventsService(queries)
	announcementService := services.NewAnnouncementService(queries)

	// =================== Server configuration ===================
	router := gin.Default()
	router.SetTrustedProxies([]string{
		"127.0.0.1/32",
	})

	// This hooks the initialized services up to their handlers and plugs them into the router
	routes.SetupRoutes(router, eventsService, announcementService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// =================== Start server in goroutine ===================
	go func() {
		serverAddr := fmt.Sprintf("localhost:%s", port)
		log.Printf("\033[32m Server started on http://%s \033[0m\n", serverAddr)

		if err := router.Run(serverAddr); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// This essentially pauses the main function until the "done" signal is received
	<-ctx.Done()
}
