package events

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

var DeleteEvent = &cobra.Command{
	Use:   "delete",
	Short: "Delete an event with its id",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")

		var overrides config.ConfigOverrides
		overrides.Host, _ = cmd.PersistentFlags().GetString("host")
		overrides.Port, _ = cmd.PersistentFlags().GetString("port")
		cfg, _ := config.Load(&overrides)

		deleteEvent(id, cfg)
	},
}

func init() {
	DeleteEvent.Flags().String("id", "", "Delete the identified event")
	DeleteEvent.MarkFlagRequired("id")
}

func deleteEvent(id string, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println("URL: %s", baseURL.String())
		fmt.Println(err)
		return
	}

	deleteURL := baseURL.JoinPath(fmt.Sprint("/v1/events/", id))

	client := &http.Client{}

	// ----- Delete Request -----
	request, err := requests.NewRequestWithAuth(http.MethodDelete, deleteURL.String(), nil)
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

	// ----- Read Response Info -----
	fmt.Println("Response status:", response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading delete response body:", err)
		return
	}

	fmt.Println(string(body))
}
