package officers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var DeleteOfficers = &cobra.Command{
	Use:   "delete --id <uuid>",
	Short: "Delete an officer with their id",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		deleteOfficer(id, host, port)
	},
}

func init() {
	// ----- URL Flags -----
	DeleteOfficers.Flags().String("id", "", "Delete an officer by their id")
	DeleteOfficers.Flags().String("host", "127.0.0.1", "Set a custom host")
	DeleteOfficers.Flags().String("port", "8080", "Set a custom port")

	DeleteOfficers.MarkFlagRequired("id")
}

func deleteOfficer(id, host, port string) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	// ----- Req id -----
	if id == "" {
		fmt.Println("ID is required to delete!")
		return
	}

	// ----- Prepare url -----
	host = fmt.Sprint(host, ":", port)
	path := fmt.Sprint("v1/board/officers/", id)

	deleteURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// ----- Send delete request -----
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

	if response == nil {
		fmt.Println("no response received")
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading delete response body:", err)
		return
	}

	fmt.Println(string(body))
}
