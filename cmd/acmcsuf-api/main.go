package main

//go:generate swag init -d .,../../internal/api/handlers,../../internal/db/models -o ../../docs --parseDependency

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api"
	_ "modernc.org/sqlite"
)

var Version = "dev"

func main() {
	// =================== Command line arg parsing ===================
	var showVersion = flag.Bool("version", false, "Show version")
	flag.Parse()

	if *showVersion {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}

	// =================== Goroutine management ===================
	log.SetPrefix("[SERVER] ")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		log.Println("\x1b[32mShutting down the server...\x1b[0m")
		// when cancel is called, it sends a "done" signal to ctx
		cancel()
	}()

	// =================== Start the server ===================
	api.Run(ctx)
}
