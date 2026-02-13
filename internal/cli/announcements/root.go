package announcements

import (
	"github.com/spf13/cobra"
)

var CLIAnnouncements = &cobra.Command{
	Use:   "announcements",
	Short: "Manage ACM CSUF's Announcements",
}

func init() {
	CLIAnnouncements.AddCommand(GetAnnouncement)
	CLIAnnouncements.AddCommand(PostAnnouncement)
	CLIAnnouncements.AddCommand(DeleteAnnouncements)
	CLIAnnouncements.AddCommand(PutAnnouncements)
}
