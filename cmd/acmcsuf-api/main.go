/*
Package main is the entry point for the acmcsuf-api server.

It is responsible for:
- Parsing command-line arguments like --version.
- Setting up a context that listens for OS interrupt signals (Ctrl+C)
  to enable graceful shutdown of the server.
- Calling the Run function in the api package to start the server.
*/
package main

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
		log.Println("[32mShutting down the server...[0m")
		// when cancel is called, it sends a "done" signal to ctx
		cancel()
	}()

	// =================== Start the server ===================
	api.Run(ctx)
}
