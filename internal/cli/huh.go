package main

import (
	"fmt"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/announcements"
	"github.com/charmbracelet/huh"
)

//Main Menu
type mainMenuState struct {
	menuOptions string
}

func beginning() {
	var state mainMenuState
	mainMenu := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				// Ask the user if they want to see API commands or learn more about the API.
				Title("ACMCSUF-CLI").
				Description("A CLI tool to help manage the API of the CSUF ACM website").
				Options(
					huh.NewOption("View Commands", "commands"),
					huh.NewOption("Learn More about the API", "flags"),
				).
				Value(&state.menuOptions),
		),
	)

	err := mainMenu.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
	if state.menuOptions == "commands" {
		commands()
	} else {
		about()
	}
}

//Commands for CLI option
type commandCli struct {
	options string
}

func commands() {
	var commandState commandCli
	commandMenu := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				//Ask the user what commands they want to use.
				Title("ACMCSUF-CLI Available Commands").
				Description("Choose a command to your heart's content.").
				Options(
					huh.NewOption("Announcements", "announcements"),
					huh.NewOption("Completion", "completion"),
					huh.NewOption("Events", "events"),
					huh.NewOption("Help", "help"),
					huh.NewOption("Officers", "officers"),
					huh.NewOption("Back", "back"),
				).
				Value(&commandState.options),
		),
	)
	err := commandMenu.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
	if commandState.options == "back" {
		beginning()
	}
	if commandState.options == "announcements" {
		announcementsMenu()
	}
}

//About CLI
type flags struct {
	options string
}

func about() {
	var flagState flags
	flagMenu := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				//Ask the user what commands they want to use.
				Title("About ACMCSUF-CLI").
				Description("Choose an option to your heart's content.").
				Options(
					huh.NewOption("Help with CLI", "help"),
					huh.NewOption("Version", "version"),
					huh.NewOption("Back", "back"),
				).
				Value(&flagState.options),
		),
	)
	err := flagMenu.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
	if flagState.options == "back" {
		beginning()
	}
	if flagState.options == "version" {
		
	}
}

//announcement menu
type announcementOptions struct {
	options string
}

func announcementsMenu() {
	var announcementState announcementOptions
	announcementMenu := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				//Ask the user what commands they want to use.
				Title("ACMCSUF-CLI Announcements").
				Description("Choose an option to your heart's content.").
				Options(
					huh.NewOption("Delete", "delete"),
					huh.NewOption("Get", "get"),
					huh.NewOption("Post", "post"),
					huh.NewOption("Put", "put"),
					huh.NewOption("Back", "back"),
				).
				Value(&announcementState.options),
		),
	)
	err := announcementMenu.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
	if announcementState.options == "delete" {

	}
	if announcementState.options == "get" {
		announcements.CLIAnnouncements.SetArgs([]string{"get"})
		announcements.GetAnnouncement.Execute()
	}
	if announcementState.options == "post" {
		announcements.CLIAnnouncements.SetArgs([]string{"post"})
		announcements.GetAnnouncement.Execute()
	}
}
func main() {
	beginning()
}
