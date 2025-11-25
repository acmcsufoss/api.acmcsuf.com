package tiers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var GetTiers = &cobra.Command{
	Use:   "get [flags]",
	Short: "Get Tiers",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		getTiers(id, port, host)
	},
}

func init() {
	GetTiers.Flags().String("tier", "", "Get a specific tier")
	GetTiers.Flags().String("host", "127.0.0.1", "Custom host")
	GetTiers.Flags().String("port", "8080", "Custom port")
}

func getTiers(id, port, host string) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	// prepare url
	host = fmt.Sprint(host, ":", port)
	path := fmt.Sprint("v1/board/tiers/", id)

	getURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// getting tier(s)
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
		var getPayload []models.CreateTierParams
		err = json.NewDecoder(response.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body without id:", err)
			return
		}

		for i := range getPayload {
			utils.PrintStruct(getPayload[i])
		}
	} else {
		var getPayload models.CreateTierParams
		err = json.NewDecoder(response.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body with id:", err)
			return
		}

		utils.PrintStruct(getPayload)
	}
}
