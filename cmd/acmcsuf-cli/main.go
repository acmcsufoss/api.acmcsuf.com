package main

// Initalizing cobra here
// To use the CLI, first, cd into this directory (/api.acmcsuf.com/cmd/csuf)
// Next, run: go install .
// Now, if you have not already, export the go bin path: export PATH="$HOME/go/bin:$PATH"
// Now type csuf in your command line and see what happens!

import (
	"log"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/announcements"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/events"

	"github.com/spf13/cobra"
)

var Version = "dev"

var rootCmd = &cobra.Command{
	Use:     "csuf",
	Short:   "A CLI tool to help manage the API of the CSUF ACM website",
	Version: Version,
}

func Execute() {
	// Logging the error, prefix is date, time, and what file the log is from
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err := rootCmd.Execute(); err != nil {
		log.Println("Error with CLI:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(events.CLIEvents)
	rootCmd.AddCommand(announcements.CLIAnnouncements)
}

func main() {
	Execute()
}
