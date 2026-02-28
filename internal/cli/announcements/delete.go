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

var DeleteAnnouncements = &cobra.Command{
	Use:   "delete",
	Short: "delete an announcement by its id",

	Run: func(cmd *cobra.Command, args []string) {
		var uuid string
		uuid, _ = cmd.Flags().GetString("id")
		deleteAnnouncement(uuid, config.Cfg)
	},
}

func init() {
	DeleteAnnouncements.Flags().String("id", "", "delete an announcement by its id")
	DeleteAnnouncements.MarkFlagRequired("id")
}

func deleteAnnouncement(id string, cfg *config.Config) {
	deleteUrl := config.GetBaseURL(cfg).JoinPath("v1", "announcements", id)

	if body, err := client.SendRequestAndReadResponse(deleteUrl, http.MethodDelete, nil); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	} else {
		utils.PrettyPrintJSON(body)
	}
}
