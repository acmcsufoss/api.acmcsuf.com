package events

import (
	"fmt"
	"net/http"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/client"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var GetEvent = &cobra.Command{
	Use:   "get",
	Short: "Get one or all events",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		getEvents(id, config.Cfg)
	},
}

func init() {
	GetEvent.Flags().String("id", "", "Get a specific event")
}

func getEvents(id string, cfg *config.Config) {
	getUrl := config.GetBaseURL(cfg).JoinPath("v1", "events", id)

	if body, err := client.SendRequestAndReadResponse(getUrl, false, http.MethodGet, nil); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		if body != nil {
			utils.PrettyPrintJSONErr(body)
		}
	} else {
		utils.PrettyPrintJSON(body)
	}
}
