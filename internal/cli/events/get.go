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
	"github.com/acmcsufoss/api.acmcsuf.com/utils/requests"
)

var GetEvent = &cobra.Command{
	Use:   "get",
	Short: "Get events",

	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		getEvents(id, config.Cfg)
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

	getURL := baseURL.JoinPath(fmt.Sprint("v1/events/", id))

	// ----- Get -----
	req, err := requests.NewRequestWithAuth(http.MethodGet, getURL.String(), nil)
	if err != nil {
		fmt.Println("Error getting the request:", err)
		return
	}
	requests.AddOrigin(req)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: couldn't make GET request: %v", err)
	}
	fmt.Println(req.Header.Get("Origin"))

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
