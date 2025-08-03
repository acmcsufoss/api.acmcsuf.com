package announcements

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/utils/cli"
	"github.com/spf13/cobra"
)

var GetAnnouncement = &cobra.Command{
	Use:   "get",
	Short: "Get an announcement",

	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		uuid, _ := cmd.Flags().GetString("id")

		getAnnouncement(host, port, uuid)
	},
}

func init() {

	// Url flags
	GetAnnouncement.Flags().String("host", "127.0.0.1", "Set a custom host")
	GetAnnouncement.Flags().String("port", "8080", "Set a custom port")
	GetAnnouncement.Flags().String("id", "", "Get a specific announcement by its id")

}

func getAnnouncement(host string, port string, uuid string) {
	// ----- Constructing the url -----
	host = fmt.Sprint(host, ":", port)
	path := "announcements"
	if uuid != "" {
		path = fmt.Sprint(path, "/", uuid)
	}

	getUrl := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// ----- Requesting Get -----
	response, err := http.Get(getUrl.String())
	if err != nil {
		fmt.Println("error with request:", err)
		return
	}

	if response == nil {
		fmt.Println("no response received")
		return
	}

	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)

	if id == "" {
		var getPayload []CreateAnnouncement
		err = json.NewDecoder(response.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body without id:", err)
			return
		}

		for i := range getPayload {
			cli.PrintStruct(getPayload[i])
		}
	} else {
		var getPayload CreateAnnouncement
		err = json.NewDecoder(response.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body with id:", err)
			return
		}

		cli.PrintStruct(getPayload)
	}
}
