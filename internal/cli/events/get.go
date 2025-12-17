package events

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var GetEvent = &cobra.Command{
	Use:   "get",
	Short: "Get events",

	Run: func(cmd *cobra.Command, args []string) {

		// If these where global, unexpected behavior would be expected :(
		id, _ := cmd.Flags().GetString("id")
		var overrides config.ConfigOverrides
		overrides.Host, _ = cmd.PersistentFlags().GetString("host")
		overrides.Port, _ = cmd.PersistentFlags().GetString("port")
		cfg, _ := config.Load(&overrides)

		getEvents(id, cfg)
	},
}

func init() {
	GetEvent.Flags().String("id", "", "Get a specific event")
}

func getEvents(id string, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	getURL := baseURL.JoinPath(fmt.Sprint("v1/events/"))

	// ----- Get -----
	req, err := http.NewRequest("GET", getURL.String(), nil)
	if err != nil {
		fmt.Println("Error getting the request:", err)
		return
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: couldn't make GET request: %v", err)
	}

	// ----- Read Response Information -----
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Response status:", resp.Status)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: couldn't read response body: %v", err)
	}
	utils.PrettyPrintJSON(body)
}
