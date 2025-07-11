package events

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

var DeleteEvent = &cobra.Command{
	Use:   "delete",
	Short: "Delete an event with its id",

	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Println("Event ID is required to delete!")
			return
		}
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		deleteEvent(id, host, port)
	},
}

func init() {
	DeleteEvent.Flags().String("id", "", "Delete the identified event")
	DeleteEvent.Flags().String("host", "127.0.0.1", "Set a custom host (ex: 127.0.0.1)")
	DeleteEvent.Flags().String("port", "8080", "Set a custom port (ex: 8080)")
}

func deleteEvent(id string, host string, port string) {
	if id == "" {
		fmt.Println("Event ID is required to delete!")
		return
	}

	host = fmt.Sprint(host, ":", port)
	path := fmt.Sprint("events/", id)

	deleteURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	//fmt.Println(deleteURL.String())

	client := &http.Client{}

	request, err := http.NewRequest(http.MethodDelete, deleteURL.String(), nil)
	if err != nil {
		fmt.Println("Error making delete request:", err)
		return
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error with delete response:", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading delete response body:,", err)
		return
	}

	fmt.Println(string(body))
}
