package announcements

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/requests"
	"github.com/spf13/cobra"
)

var GetAnnouncement = &cobra.Command{
	Use:   "get",
	Short: "Get an announcement",

	Run: func(cmd *cobra.Command, args []string) {
		uuid, _ := cmd.Flags().GetString("id")
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
	req, err := requests.NewRequestWithAuth(http.MethodGet, getUrl.String(), nil)
	if err != nil {
		fmt.Println("error with request:", err)
		return
	}
	requests.AddOrigin(req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error getting announcements:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)

	if uuid == "" {
		var getPayload []models.CreateAnnouncementParams
		err = json.NewDecoder(res.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body without id:", err)
			return
		}

		for i := range getPayload {
			utils.PrintStruct(getPayload[i])
		}
	} else {
		var getPayload models.CreateAnnouncementParams
		err = json.NewDecoder(res.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body with id:", err)
			return
		}

		utils.PrintStruct(getPayload)
	}

}
