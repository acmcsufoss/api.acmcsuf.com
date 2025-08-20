package events

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var GetEvent = &cobra.Command{
	Use:   "get",
	Short: "Get events",

	Run: func(cmd *cobra.Command, args []string) {

		// If these where global, unexpected behavior would be expected :(
		id, _ := cmd.Flags().GetString("id")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		getEvents(id, port, host)
	},
}

func init() {

	// Url Flags
	GetEvent.Flags().String("id", "", "Get a specific event")
	GetEvent.Flags().String("host", "127.0.0.1", "Custom host")
	GetEvent.Flags().String("port", "8080", "Custom port")

}

func getEvents(id string, port string, host string) {

	// ----- Constructing url -----
	// Combining Host and port
	host = fmt.Sprint(host, ":", port)

	// Constructing Path
	path := "events"
	if id != "" {
		path = fmt.Sprint(path, "/", id)
	}

	getURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// ----- Get -----
	response, err := http.Get(getURL.String())
	if err != nil {
		fmt.Println("Error getting the request:", err)
		return
	}

	if response == nil {
		fmt.Println("no response received")
		return
	}

	defer response.Body.Close()

	// ----- Read Response Information -----
	fmt.Println("Response status:", response.Status)

	if id == "" {
		var getPayload []models.CreateEventParams
		err = json.NewDecoder(response.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body without id:", err)
			return
		}

		for i := range getPayload {
			utils.PrintStruct(getPayload[i])
		}
	} else {
		var getPayload models.CreateEventParams
		err = json.NewDecoder(response.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body with id:", err)
			return
		}

		utils.PrintStruct(getPayload)
	}

}
