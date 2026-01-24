package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/announcements"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/events"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/officers"
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
	Run: func(cmd *cobra.Command, args []string) {
		// do nothing
	},
}

// init() is a special function that always gets run before main
func init() {
	rootCmd.AddCommand(events.CLIEvents)
	rootCmd.AddCommand(announcements.CLIAnnouncements)
	rootCmd.AddCommand(officers.CLIOfficers)
	rootCmd.AddCommand(config.ConfigCmd)

	rootCmd.PersistentFlags().String("host", "", "Override configured/default host")
	rootCmd.PersistentFlags().String("port", "", "Override configured/default port")
	rootCmd.PersistentFlags().String("origin", "", "Override configured/default origin")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		overrides := &config.ConfigOverrides{
			Host: cmd.Flag("host").Value.String(),
			Port: cmd.Flag("port").Value.String(),
		}
		var err error
		config.Cfg, err = config.Load(overrides)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		return nil
	}
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

// Menu function for huh library
func Menu() {
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
					huh.NewOption("Overide Config", "config"),
					huh.NewOption("Version", "version"),
					huh.NewOption("Exit", "exit"),
				).
				Value(&commandState),
		),
	)
	err := commandMenu.Run()
	if err != nil {
		if err == huh.ErrUserAborted {
			fmt.Println("User canceled the form — exiting.")
			return
		}
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
	if commandState == "announcements" {
		announcements.ShowMenu(Menu)
	}
	if commandState == "officers" {
		officers.ShowMenu(Menu)
	}
	if commandState == "events" {
		events.ShowMenu(Menu)
	}
	if commandState == "config" {
		var flagsChosen []string
		err := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					//Ask the user what commands they want to use.
					Title("ACMCSUF-CLI Config Override").
					Description("Choose a command(s). Note: Use spacebar to select and if done click enter.\nTo skip, simply click enter.").
					Options(
						huh.NewOption("Change Host", "host"),
						huh.NewOption("Change Port", "port"),
						huh.NewOption("Change Origin", "origin"),
					).
					Value(&flagsChosen),
			),
		).Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		for index, flag := range flagsChosen {
			switch flag {
			case "host":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Config Override:").
					Description("Please enter the custom host:").
					Prompt("> ").
					Value(&config.CfgOverride.Host).
					Run()
			case "port":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Config Override:").
					Description("Please enter the custom port:").
					Prompt("> ").
					Value(&config.CfgOverride.Port).
					Run()
			case "origin":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Config Override:").
					Description("Please enter the custom origin:").
					Prompt("> ").
					Run()
			}
			if err != nil {
				if err == huh.ErrUserAborted {
					fmt.Println("User canceled the form — exiting.")
				}
				fmt.Println("Uh oh:", err)
				os.Exit(1)
			}
			_ = index
		}
		Menu()

	}
	if commandState == "version" {
		fmt.Println("ACMCSUF-CLI Version:", Version)
		Menu()
	}
}
