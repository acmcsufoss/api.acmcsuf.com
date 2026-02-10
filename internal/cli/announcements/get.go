package announcements

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
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
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	// ----- Constructing the url -----
	path := "v1/announcements"
	if uuid != "" {
		path = fmt.Sprint(path, "/", uuid)
	}

	getUrl := baseURL.JoinPath(path)

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
		fmt.Println("Response status:", res.Status)
		return
	}

	if uuid == "" {
		var getPayload []dbmodels.CreateAnnouncementParams
		err = json.NewDecoder(res.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body without id:", err)
			return
		}

		for i := range getPayload {
			fmt.Println(utils.PrintStruct(getPayload[i]))
		}
	} else {
		var getPayload dbmodels.CreateAnnouncementParams
		err = json.NewDecoder(res.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body with id:", err)
			return
		}

		fmt.Println(utils.PrintStruct(getPayload))
	}

}
