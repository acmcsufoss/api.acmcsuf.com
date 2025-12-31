package officers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var GetOfficers = &cobra.Command{
	Use:   "get [flags]",
	Short: "Get Officers",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		getOfficers(id, config.Cfg)
	},
}

func init() {
	GetOfficers.Flags().String("id", "", "Get a specific officer")
}

func getOfficers(id string, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	// prepare url
	path := fmt.Sprint("v1/board/officers/", id)

	getURL := baseURL.JoinPath(path)

	// getting officer(s)
	response, err := http.Get(getURL.String())
	if err != nil {
		fmt.Println("error getting the request: ", err)
		return
	}
	if response == nil {
		fmt.Println("no response recieved")
		return
	}

	defer response.Body.Close()

	if id == "" {
		var getPayload []models.GetOfficerRow
		err = json.NewDecoder(response.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body without id:", err)
			return
		}

		for i := range getPayload {
			utils.PrintStruct(getPayload[i])
		}
	} else {
		var getPayload models.GetOfficerRow
		err = json.NewDecoder(response.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body with id:", err)
			return
		}

		utils.PrintStruct(getPayload)
	}
}
