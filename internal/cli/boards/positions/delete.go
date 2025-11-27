package positions

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var DeletePosition = &cobra.Command{
	Use:   "delete --oid <oid>",
	Short: "Delete a position",

	Run: func(cmd *cobra.Command, args []string) {
		oid, _ := cmd.Flags().GetString("oid")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		deletePosition(oid, host, port)
	},
}

func init() {
	DeletePosition.Flags().String("oid", "", "Delete a position")
	DeletePosition.Flags().String("host", "127.0.0.1", "Set a custom host")
	DeletePosition.Flags().String("port", "8080", "Set a custom port")

	DeletePosition.MarkFlagRequired("oid")
}

func deletePosition(id, host, port string) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	// req id
	if id == "" {
		fmt.Println("oid is required to delete!")
		return
	}

	// prepare url
	host = fmt.Sprint(host, ":", port)
	path := fmt.Sprint("v1/board/positions/", id)

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
