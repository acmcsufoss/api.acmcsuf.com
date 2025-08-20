package events

import (
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
