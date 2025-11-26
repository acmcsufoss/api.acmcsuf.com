package events

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/requests"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
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

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	// ----- Constructing url -----
	// Combining Host and port
	host = fmt.Sprint(host, ":", port)

	// Constructing Path
	path := "v1/events"
	if id != "" {
		path = fmt.Sprint(path, "/", id)
	}

	getURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// ----- Get -----
	req, err := requests.NewRequestWithAuth("GET", getURL.String(), nil)
	if err != nil {
		fmt.Println("Error getting the request:", err)
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: couldn't make GET request: %v", err)
	}
	defer resp.Body.Close()

	if req == nil {
		fmt.Println("no response received")
		return
	}

	// ----- Read Response Information -----
	fmt.Println("Response status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: couldn't read response body: %v", err)
	}

	prettyJSON := pretty.Pretty(body)
	fmt.Println(string(prettyJSON))
}
