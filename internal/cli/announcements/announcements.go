package announcements

import (
	"database/sql"

	"github.com/spf13/cobra"
)

var CLIAnnouncements = &cobra.Command{
	Use:   "announcements",
	Short: "Manage ACM CSUF's Announcements",
}

type CreateAnnouncement struct {
	Uuid             string         `json:"uuid"`
	Visibility       string         `json:"visibility"`
	AnnounceAt       int64          `json:"announce_at"`
	DiscordChannelID sql.NullString `json:"discord_channel_id"`
	DiscordMessageID sql.NullString `json:"discord_message_id"`
}

type UpdateAnnouncement struct {
	Visibility       sql.NullString `json:"visibility"`
	AnnounceAt       sql.NullInt64  `json:"announce_at"`
	DiscordChannelID sql.NullString `json:"discord_channel_id"`
	DiscordMessageID sql.NullString `json:"discord_message_id"`
	Uuid             string         `json:"uuid"`
}

func init() {
	CLIAnnouncements.AddCommand(GetAnnouncement)
	//CLIAnnouncements.AddCommand(PostAnnouncement)
	//CLIAnnouncements.AddCommand(DeleteAnnouncements)
	//CLIAnnouncements.AddCommand(PutAnnouncements)
}
