package announcements

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/forms"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var GetAnnouncement = &cobra.Command{
	Use:   "get",
	Short: "Get an announcement",

	Run: func(cmd *cobra.Command, args []string) {
		var uuid string
		if cmd.Flags().Changed("id") {
			uuid, _ = cmd.Flags().GetString("id")
		} else {
			uuid, _ = forms.GetIdInteractive()
		}
		getAnnouncement(uuid, config.Cfg)
	},
}

func init() {
	GetAnnouncement.Flags().String("id", "", "Get a specific announcement by its id")
}

func getAnnouncement(uuid string, cfg *config.Config) {
	getUrl := config.GetBaseURL(cfg).JoinPath("v1", "announcements", uuid)

	// ----- Requesting Get -----
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, getUrl.String(), nil)
	if err != nil {
		fmt.Println("error with request:", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error getting announcements:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP", res.Status)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: failed to read response body:", err)
		return
	}
	utils.PrettyPrintJSON(body)
}
