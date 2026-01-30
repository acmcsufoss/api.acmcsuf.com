package announcements

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
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

func ShowMenu(backCallback func()) {
	var announcementState string
	announcementMenu := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("ACMCSUF-CLI Announcements").
				Description("Choose an option to your heart's content.").
				Options(
					huh.NewOption("Delete", "delete"),
					huh.NewOption("Get", "get"),
					huh.NewOption("Post", "post"),
					huh.NewOption("Put", "put"),
					huh.NewOption("Back", "back"),
				).
				Value(&announcementState),
		),
	)
	err := announcementMenu.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	if announcementState == "delete" {
		DeleteAnnouncements.Run(DeleteAnnouncements, []string{})
		backCallback()
	} else if announcementState == "get" {
		GetAnnouncement.Run(GetAnnouncement, []string{})
		backCallback()
	} else if announcementState == "post" {
		PostAnnouncement.Run(PostAnnouncement, []string{})
		backCallback()
	} else if announcementState == "put" {
		PutAnnouncements.Run(PutAnnouncements, []string{})
		backCallback()
	} else if announcementState == "back" {
		backCallback()
	}
}
