package events

import (
	"github.com/spf13/cobra"
)

var CLIEvents = &cobra.Command{
	Use:   "events",
	Short: "Manage ACM CSUF's events",
}

func init() {
	CLIEvents.AddCommand(PostEvent)
	CLIEvents.AddCommand(GetEvent)
	CLIEvents.AddCommand(PutEvents)
	CLIEvents.AddCommand(DeleteEvent)
}
