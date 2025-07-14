package cmd

import (
	"log"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/events"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/announcements"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/events"
)

var rootCmd = &cobra.Command{
	Use:   "csuf",
	Short: "A CLI tool to help manage the API of the CSUF ACM website",
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
