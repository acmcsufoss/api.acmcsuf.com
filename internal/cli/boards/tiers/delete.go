package tiers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var DeleteTier = &cobra.Command{
	Use:   "delete --tier <uuid>",
	Short: "Delete a tier",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("tier")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		deleteTier(id, host, port)
	},
}

func init() {
	DeleteTier.Flags().String("tier", "", "Delete a tier")
	DeleteTier.Flags().String("host", "127.0.0.1", "Set a custom host")
	DeleteTier.Flags().String("port", "8080", "Set a custom port")

	DeleteTier.MarkFlagRequired("tier")
}

func deleteTier(id, host, port string) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	// req id
	if id == "" {
		fmt.Println("ID is required to delete!")
		return
	}

	// prepare url
	host = fmt.Sprint(host, ":", port)
	path := fmt.Sprint("v1/board/officers/", id)

	deleteURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// send delete request
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
