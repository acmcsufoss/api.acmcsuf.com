package officers

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

type officerFlags struct {
	uuid     bool
	fullname bool
	picture  bool
	github   bool
	discord  bool
}

var CLIOfficers = &cobra.Command{
	Use:   "officers HEADER",
	Short: "A command to manage officers.",
}

func init() {
	CLIOfficers.AddCommand(GetOfficers)
	CLIOfficers.AddCommand(DeleteOfficers)
	CLIOfficers.AddCommand(PostOfficer)
	CLIOfficers.AddCommand(PutOfficer)
}

func ShowMenu(backCallback func()) {
	var boardState string
	boardMenu := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("ACMCSUF-CLI Board").
				Description("Choose an option to your heart's content.").
				Options(
					huh.NewOption("Delete", "delete"),
					huh.NewOption("Get", "get"),
					huh.NewOption("Post", "post"),
					huh.NewOption("Put", "put"),
					huh.NewOption("Back", "back"),
				).
				Value(&boardState),
		),
	)
	err := boardMenu.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	if boardState == "delete" {
		DeleteOfficers.Run(DeleteOfficers, []string{})
		backCallback()
	} else if boardState == "get" {
		GetOfficers.Run(GetOfficers, []string{})
		backCallback()
	} else if boardState == "post" {
		PostOfficer.Run(PostOfficer, []string{})
		backCallback()
	} else if boardState == "put" {
		PutOfficer.Run(PutOfficer, []string{})
		backCallback()
	} else if boardState == "back" {
		backCallback()
	}
}
