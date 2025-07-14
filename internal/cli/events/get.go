package events

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

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
	GetEvent.Flags().String("host", "127.0.0.1", "Custom host (ex: 127.0.0.1)")
	GetEvent.Flags().String("port", "8080", "Custom port (ex: 8080)")

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
	defer response.Body.Close()

	// ----- Read Response Information -----
	fmt.Println("Response status:", response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	fmt.Println("Response body:", string(body))

}
