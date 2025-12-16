package cli

import (
	"log"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/announcements"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/officers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/events"

	"github.com/spf13/cobra"
)

type exitCode int

// These are stolen from github.com/cli/cli, not all have to be utilized
const (
	exitOK      exitCode = 0
	exitError   exitCode = 1
	exitCancel  exitCode = 2
	exitAuth    exitCode = 4
	exitPending exitCode = 8
)

var Version = "dev"

var rootCmd = &cobra.Command{
	Use:     os.Args[0],
	Short:   "A CLI tool to help manage the API of the CSUF ACM website",
	Version: Version,
}

// init() is a special function that always gets run before main
func init() {
	rootCmd.AddCommand(events.CLIEvents)
	rootCmd.AddCommand(announcements.CLIAnnouncements)
	rootCmd.AddCommand(officers.CLIOfficers)
}

// Function that gets called by main
func Execute() exitCode {
	// Logging the error, prefix is date, time, and what file the log is from
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if err := rootCmd.Execute(); err != nil {
		log.Println("Error:", err)
		return exitError
	}

	return exitOK
}
