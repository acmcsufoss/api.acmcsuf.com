package main

import (
	"fmt"
	"log"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/announcements"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/boards/officers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/events"
	"github.com/charmbracelet/huh"

	"github.com/spf13/cobra"
)

var Version = "dev"

var rootCmd = &cobra.Command{
	Use:     os.Args[0],
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
	rootCmd.AddCommand(officers.CLIOfficers)
}

func menu() {
	var commandState string
	commandMenu := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				//Ask the user what commands they want to use.
				Title("ACMCSUF-CLI Available Commands").
				Description("A CLI tool to help manage the API of the CSUF ACM website.").
				Options(
					huh.NewOption("Announcements", "announcements"),
					huh.NewOption("Officers", "officers"),
					huh.NewOption("Events", "events"),
					huh.NewOption("Version", "version"),
					huh.NewOption("Exit", "exit"),
				).
				Value(&commandState),
		),
	)
	err := commandMenu.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
	if commandState == "back" {

	}
	if commandState == "announcements" {
		announcements.ShowMenu(menu)
	}
	if commandState == "officers" {

	}
	if commandState == "events" {

	}
	if commandState == "completion" {

	}
}
func main() {
	exitCode := cli.Execute()
	menu()
	os.Exit(int(exitCode))
}
