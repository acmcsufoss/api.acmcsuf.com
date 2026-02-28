package officers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/oauth"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var DeleteOfficers = &cobra.Command{
	Use:   "delete --id <uuid>",
	Short: "Delete an officer with their id",

	Run: func(cmd *cobra.Command, args []string) {
		var uuid string
		uuid, _ = cmd.Flags().GetString("id")
		deleteOfficer(uuid, config.Cfg)
	},
}

func init() {
	DeleteOfficers.Flags().String("id", "", "Delete an officer by their id")
	DeleteOfficers.MarkFlagRequired("id")
}

func deleteOfficer(id string, cfg *config.Config) {
	deleteUrl := config.GetBaseURL(cfg).JoinPath("v1", "board", "officers", id)

	request, err := oauth.NewRequestWithAuth(http.MethodDelete, deleteUrl.String(), nil)
	if err != nil {
		fmt.Println("Error: failed to construct delete request:", err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error: failed to send delete request:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP", response.Status)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: failed to read response body:", err)
		return
	}
	utils.PrettyPrintJSON(body)
}
