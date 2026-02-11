package announcements

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/forms"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/oauth"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var DeleteAnnouncements = &cobra.Command{
	Use:   "delete",
	Short: "delete an announcement by its id",

	Run: func(cmd *cobra.Command, args []string) {
		var uuid string

		if cmd.Flags().Changed("id") {
			uuid, _ = cmd.Flags().GetString("id")
		} else {
			uuid, _ = forms.GetIdInteractive()
		}
		if uuid == "" {
			fmt.Println("Error: no ID specified")
			return
		}
		deleteAnnouncement(uuid, config.Cfg)
	},
}

func init() {
	DeleteAnnouncements.Flags().String("id", "", "delete an announcement by its id")
	DeleteAnnouncements.MarkFlagRequired("id")
}

func deleteAnnouncement(id string, cfg *config.Config) {
	deleteUrl := config.GetBaseURL(cfg).JoinPath("v1", "announcements", id)

	// ----- Delete -----
	request, err := oauth.NewRequestWithAuth(http.MethodDelete, deleteUrl.String(), nil)
	if err != nil {
		fmt.Println("Error: failed to construct delete request:", err)
		return
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error: failed to send delete request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP", response.Status)
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: failed to read response body:", err)
		return
	}
	utils.PrettyPrintJSON(body)
}
