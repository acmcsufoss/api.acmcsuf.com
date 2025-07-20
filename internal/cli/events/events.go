package events

import (
	"database/sql"

	"github.com/spf13/cobra"
)

var CLIEvents = &cobra.Command{
	Use:   "events HEADER",
	Short: "A command to manage events.",
}

type CreateEvent struct {
	Uuid     string `json:"uuid"`
	Location string `json:"location"`
	StartAt  int64  `json:"start_at"`
	EndAt    int64  `json:"end_at"`
	IsAllDay bool   `json:"is_all_day"`
	Host     string `json:"host"`
}

type UpdateEvent struct {
	Uuid     string         `json:"uuid"`
	Location sql.NullString `json:"location"`
	StartAt  sql.NullInt64  `json:"start_at"`
	EndAt    sql.NullInt64  `json:"end_at"`
	IsAllDay sql.NullBool   `json:"is_all_day"`
	Host     sql.NullString `json:"host"`
}

func init() {
	CLIEvents.AddCommand(PostEvent)
	CLIEvents.AddCommand(GetEvent)
	CLIEvents.AddCommand(PutEvents)
	CLIEvents.AddCommand(DeleteEvent)
}
