package announcements

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/client"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var GetAnnouncement = &cobra.Command{
	Use:   "get",
	Short: "Get an announcement",

	Run: func(cmd *cobra.Command, args []string) {
		var uuid string
		uuid, _ = cmd.Flags().GetString("id")
		getAnnouncement(uuid, config.Cfg)
	},
}

func init() {
	GetAnnouncement.Flags().String("id", "", "Get a specific announcement by its id")
}

func getAnnouncement(uuid string, cfg *config.Config) {
	getUrl := config.GetBaseURL(cfg).JoinPath("v1", "announcements", uuid)

	if body, err := client.SendRequestAndReadResponse(getUrl, http.MethodGet, nil); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		utils.PrettyPrintJSON(body)
	}
}
