package officers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var DeleteOfficers = &cobra.Command{
	Use:   "delete --id <uuid>",
	Short: "Delete an officer with their id",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")

		var overrides config.Config
		overrides.Host, _ = cmd.PersistentFlags().GetString("host")
		overrides.Port, _ = cmd.PersistentFlags().GetString("port")
		cfg, _ := config.Load(&overrides)

		deleteOfficer(id, cfg)
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
		fmt.Println("URL: %s", baseURL.String())
		fmt.Println(err)
		return
	}

	if id == "" {
		fmt.Println("ID is required to delete!")
		return
	}
	deleteURL := baseURL.JoinPath(fmt.Sprint("v1/board/officers/", id))

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
