package main

import (
	"fmt"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/announcements"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/boards/officers"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/events"
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
					huh.NewOption("Exit", "exit"),
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
	}
	if state.menuOptions == "flags" {
		about()
	}
	if state.menuOptions == "exit" {}
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
					huh.NewOption("Officers", "officers"),
					huh.NewOption("Events", "events"),
					huh.NewOption("Help", "help"),
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
	if commandState.options == "officers" {
		boardMenu()
	}
	if commandState.options == "events" {
		eventsMenu()
	}
	if commandState.options == "completion" {

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
		announcements.CLIAnnouncements.Execute()
	}
	if announcementState.options == "get" {
		announcements.CLIAnnouncements.SetArgs([]string{"get"})
		announcements.CLIAnnouncements.Execute()
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
		args := []string{"put"}
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
	beginning()
}

type boardOptions struct {
	options string
}

func boardMenu() {
	var boardState boardOptions
	boardMenu := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				//Ask the user what commands they want to use.
				Title("ACMCSUF-CLI Boards").
				Description("Choose an option to your heart's content.").
				Options(
					huh.NewOption("Delete", "delete"),
					huh.NewOption("Get", "get"),
					huh.NewOption("Post", "post"),
					huh.NewOption("Put", "put"),
					huh.NewOption("Back", "back"),
				).
				Value(&boardState.options),
		),
	)
	err := boardMenu.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
	if boardState.options == "delete"{
		var uuid string
		huh.NewInput().
			Title("ACMCSUF-CLI Board Delete:").
			Description("Please enter the officer's uuid:").
			Prompt("> ").
			Value(&uuid).
			Run()
		officers.CLIOfficers.SetArgs([]string{"delete", "--id", uuid})
		officers.CLIOfficers.Execute()
	}
	if boardState.options == "get" {
		officers.CLIOfficers.SetArgs([]string{"get"})
		officers.CLIOfficers.Execute()
	}
	if boardState.options == "post" {
		var uuid string
		var fullName string
		var pictureLink string
		var githubLink string
		var discordLink string
		huh.NewInput().
			Title("ACMCSUF-CLI Board Post:").
			Description("Please enter officer's uuid:").
			Prompt("> ").
			Value(&uuid).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Board Post:").
			Description("Please enter the officer's full name:").
			Prompt("> ").
			Value(&fullName).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Board Post:").
			Description("Please enter the picture link for officer:").
			Prompt("> ").
			Value(&pictureLink).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Board Post:").
			Description("Please enter the github link for officer:").
			Prompt("> ").
			Value(&githubLink).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Board Post:").
			Description("Please enter the discord link for officer").
			Prompt("> ").
			Value(&discordLink).
			Run()
		officers.CLIOfficers.SetArgs([]string{"post", "--uuid", uuid, "-n", fullName, "-p", pictureLink, "-g", githubLink, "-d", discordLink})
		officers.CLIOfficers.Execute()
	}
	if boardState.options == "put" {
		var uuid string
		valid := false
		args := []string{"put"}
		for !valid {
			huh.NewInput().
				Title("ACMCSUF-CLI Board Put:").
				Description("Please enter the officer's uuid:").
				Prompt("> ").
				Value(&uuid).
				Run()
			if uuid != "" {
				valid = true
				args = append(args, "--id", uuid)
			} else {
			}
		}
		officers.CLIOfficers.SetArgs(args)
		officers.CLIOfficers.Execute()

	}
	beginning()
}
type eventOptions struct {
	options string
}
func eventsMenu() {
	var eventState eventOptions
	eventMenu := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				//Ask the user what commands they want to use.
				Title("ACMCSUF-CLI Menu").
				Description("Choose an option to your heart's content.").
				Options(
					huh.NewOption("Delete", "delete"),
					huh.NewOption("Get", "get"),
					huh.NewOption("Post", "post"),
					huh.NewOption("Put", "put"),
					huh.NewOption("Back", "back"),
				).
				Value(&eventState.options),
		),
	)
	err := eventMenu.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
	if eventState.options == "delete"{
		var uuid string
		huh.NewInput().
			Title("ACMCSUF-CLI Event Delete:").
			Description("Please enter the events's uuid:").
			Prompt("> ").
			Value(&uuid).
			Run()
		events.CLIEvents.SetArgs([]string{"delete", "--id", uuid})
		events.CLIEvents.Execute()
	}
	if eventState.options == "get" {
		events.CLIEvents.SetArgs([]string{"get"})
		events.CLIEvents.Execute()
	}
	if eventState.options == "post" {
		var uuid string
		var location string
		var timeStart string
		var duration string
		var allDayYes string
		var host string
		huh.NewInput().
			Title("ACMCSUF-CLI Event Post:").
			Description("Please enter event's uuid:").
			Prompt("> ").
			Value(&uuid).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Event Post:").
			Description("Please enter the event's location:").
			Prompt("> ").
			Value(&location).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Event Post:").
			Description("Please enter the start time of the event in the following format:\n [Month]/[Day]/[Year] [Hour]:[Minute][PM | AM]\nFor example: \x1b[93m01/02/06 03:04PM\x1b[0m").
			Prompt("> ").
			Value(&timeStart).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Event Post:").
			Description("Please enter the duration of the event in the following format:\n [Hour]:[Minute]\nFor example: \x1b[93m03:04\x1b[0m").
			Prompt("> ").
			Value(&duration).
			Run()
		huh.NewSelect[string]().
			Title("ACMCSUF-CLI Event Post:").
			Description("Is your event all day?").
			Options(
				huh.NewOption("Yes", "yes"),
				huh.NewOption("No", "n"),
			).
			Value(&allDayYes).
			Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Event Post:").
			Description("Please enter the event host").
			Prompt("> ").
			Value(&host).
			Run()
		events.CLIEvents.SetArgs([]string{"post", "--uuid", uuid, "-l", location, "-s", timeStart, "-d", duration, "-a", allDayYes, "-H", host})
		events.CLIEvents.Execute()
	}
	if eventState.options == "put" {
		var uuid string
		valid := false
		args := []string{"put"}
		for !valid {
			huh.NewInput().
				Title("ACMCSUF-CLI Event Put:").
				Description("Please enter the events's uuid:").
				Prompt("> ").
				Value(&uuid).
				Run()
			if uuid != "" {
				valid = true
				args = append(args, "--id", uuid)
			} else {
			}
		}
		events.CLIEvents.SetArgs(args)
		events.CLIEvents.Execute()
	}
	beginning()
}


func main() {
	beginning()
}
