package officers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var GetOfficers = &cobra.Command{
	Use:   "get [flags]",
	Short: "Get Officers",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		getOfficers(id, port, host)
	},
}

func init() {
	GetOfficers.Flags().String("id", "", "Get a specific officer")
	GetOfficers.Flags().String("host", "127.0.0.1", "Custom host")
	GetOfficers.Flags().String("port", "8080", "Custom port")
}

func getOfficers(id, port, host string) {
	// prepare url
	host = fmt.Sprint(host, ":", port)
	path := fmt.Sprint("v1/officers/", id)

	getURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// getting officer(s)
	response, err := http.Get(getURL.String())
	if err != nil {
		fmt.Println("error getting the request: ", err)
	}
	if response == nil {
		fmt.Println("no response recieved")
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
