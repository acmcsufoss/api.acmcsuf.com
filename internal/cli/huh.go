package main

import (
	"fmt"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/announcements"
	"github.com/charmbracelet/huh"
)

// Main Menu
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

// Commands for CLI option
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

// About CLI
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

// announcement menu
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
		var uuid string
		huh.NewInput().
			Title("ACMCSUF-CLI Announcements Delete:").
			Description("Please enter the announcement's uuid:").
			Prompt("> ").
			Value(&uuid).
			Run()
		announcements.CLIAnnouncements.SetArgs([]string{"delete", "--id", uuid})
		announcements.GetAnnouncement.Execute()
	}
	if announcementState.options == "get" {
		announcements.CLIAnnouncements.SetArgs([]string{"get"})
		announcements.GetAnnouncement.Execute()
	}
	if announcementState.options == "post" {
		var uuid string
		var visibility string
		var announceAt string
		var channelid string
		var messageid string
		huh.NewInput().
			Title("ACMCSUF-CLI Announcements Post:").
			Description("Please enter the announcement's uuid:").
			Prompt("> ").
			Value(&uuid).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Announcements Post:").
			Description("Please enter the announcement's visibility:").
			Prompt("> ").
			Value(&visibility).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Announcements Post:").
			Description("Please enter the \"announce at\" of the announcement in the following format:\n[Month]/[Day]/[Year] [Hour]:[Minutes][PM | AM]").
			Prompt("> ").
			Value(&announceAt).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Announcements Post:").
			Description("Please enter the announcement's discord channel id:").
			Prompt("> ").
			Value(&channelid).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Announcements Post:").
			Description("Please enter the announcement's message id:").
			Prompt("> ").
			Value(&messageid).
			Run()
		announcements.CLIAnnouncements.SetArgs([]string{"post", "--uuid", uuid, "-v", visibility, "-a", announceAt, "--channelid", channelid, "--messageid", messageid})
		announcements.CLIAnnouncements.Execute()

	}
	if announcementState.options == "put" {
		var uuid string
		valid := false
		args:= []string{"put",}
		for !valid {
			huh.NewInput().
				Title("ACMCSUF-CLI Announcements Put:").
				Description("Please enter the announcement's uuid:").
				Prompt("> ").
				Value(&uuid).
				Run()
			if uuid != "" {
				valid = true
				args = append(args, "--id", uuid)
			} else {
			}
		}
		announcements.CLIAnnouncements.SetArgs(args)
		announcements.CLIAnnouncements.Execute()

	}
}
func main() {
	beginning()
}
