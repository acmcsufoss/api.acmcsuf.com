package events

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/client"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var DeleteEvent = &cobra.Command{
	Use:   "delete",
	Short: "Delete an event with its id",

	Run: func(cmd *cobra.Command, args []string) {
		uuid, _ := cmd.Flags().GetString("id")
		deleteEvent(uuid, config.Cfg)
	},
}

func init() {
	DeleteEvent.Flags().String("id", "", "Delete the identified event")
	DeleteEvent.MarkFlagRequired("id")
}

func deleteEvent(id string, cfg *config.Config) {
	deleteURL := config.GetBaseURL(cfg).JoinPath("v1", "events", id)

	if body, err := client.SendRequestAndReadResponse(deleteURL, true, http.MethodDelete, nil); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		if body != nil {
			utils.PrettyPrintJSONErr(body)
		}
	} else {
		utils.PrettyPrintJSON(body)
	}
}
