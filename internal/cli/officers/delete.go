package officers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/requests"
)

var DeleteOfficers = &cobra.Command{
	Use:   "delete --id <uuid>",
	Short: "Delete an officer with their id",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		deleteOfficer(id, config.Cfg)
	},
}

func init() {
	DeleteOfficers.Flags().String("id", "", "Delete an officer by their id")
	DeleteOfficers.MarkFlagRequired("id")
}

func deleteOfficer(id string, cfg *config.Config) {
	// prepare url
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	deleteURL := baseURL.JoinPath(fmt.Sprint("v1/board/officers/", id))

	request, err := requests.NewRequestWithAuth(http.MethodDelete, deleteURL.String(), nil)
	if err != nil {
		fmt.Println("Error making delete request:", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error with delete response:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Response status:", response.Status)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading delete response body:", err)
		return
	}
	utils.PrettyPrintJSON(body)
}
