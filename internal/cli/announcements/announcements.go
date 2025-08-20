package announcements

import (
	"github.com/spf13/cobra"
)

var CLIAnnouncements = &cobra.Command{
	Use:   "announcements",
	Short: "Manage ACM CSUF's Announcements",
}

type announcementFlags struct {
	id         bool
	visibility bool
	announceat bool
	channelid  bool
	messageid  bool
}

func init() {
	CLIAnnouncements.AddCommand(GetAnnouncement)
	CLIAnnouncements.AddCommand(PostAnnouncement)
	CLIAnnouncements.AddCommand(DeleteAnnouncements)
	CLIAnnouncements.AddCommand(PutAnnouncements)
}
