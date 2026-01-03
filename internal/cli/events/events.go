package events

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

type eventFlags struct {
	uuid     bool
	location bool
	startat  bool
	duration bool
	isallday bool
	host     bool
}

var CLIEvents = &cobra.Command{
	Use:   "events HEADER",
	Short: "A command to manage events.",
}

func init() {
	CLIEvents.AddCommand(PostEvent)
	CLIEvents.AddCommand(GetEvent)
	CLIEvents.AddCommand(PutEvents)
	CLIEvents.AddCommand(DeleteEvent)
}

func ShowMenu(backCallback func()) {
	var eventState string
	eventMenu := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("ACMCSUF-CLI Event").
				Description("Choose an option to your heart's content.").
				Options(
					huh.NewOption("Delete", "delete"),
					huh.NewOption("Get", "get"),
					huh.NewOption("Post", "post"),
					huh.NewOption("Put", "put"),
					huh.NewOption("Back", "back"),
				).
				Value(&eventState),
		),
	)
	err := eventMenu.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	if eventState == "delete" {
		DeleteEvent.Run(DeleteEvent, []string{})
		backCallback()
	} else if eventState == "get" {
		GetEvent.Run(GetEvent, []string{})
		backCallback()
	} else if eventState == "post" {
		PostEvent.Run(PostEvent, []string{})
		backCallback()
	} else if eventState == "put" {
		PutEvents.Run(PutEvents, []string{})
		backCallback()
	} else if eventState == "back" {
		backCallback()
	}
}
