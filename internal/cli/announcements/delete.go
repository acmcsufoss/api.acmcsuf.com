package announcements

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

var DeleteAnnouncements = &cobra.Command{
	Use:   "delete",
	Short: "delete an announcement by its id",

	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		uuid, _ := cmd.Flags().GetString("id")

		deleteAnnouncement(host, port, uuid)
	},
}

func init() {

	// Url flags
	DeleteAnnouncements.Flags().String("host", "127.0.0.1", "set a custom host")
	DeleteAnnouncements.Flags().String("port", "8080", "set a custom port")
	DeleteAnnouncements.Flags().String("id", "", "delete an announcement by its id")

	DeleteAnnouncements.MarkFlagRequired("id")
}

func deleteAnnouncement(host string, port string, id string) {
	if id == "" {
		fmt.Println("ID is required to delete an announcement! Please use the --id flag")
		return
	}

	// ----- Constructing Url -----
	host = fmt.Sprint(host, ":", port)
	path := fmt.Sprint("announcements/", id)

	deleteUrl := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// ----- Delete -----
	request, err := http.NewRequest(http.MethodDelete, deleteUrl.String(), nil)
	if err != nil {
		fmt.Println("error with delete request:", err)
		return
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error with delete response:", err)
		return
	}

	if response == nil {
		fmt.Println("no response received")
		return
	}

	defer response.Body.Close()

	// ----- Read Response Information -----
	fmt.Println("Request status:", response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error with reading body:", err)
		return
	}

	fmt.Println(string(body))
}
