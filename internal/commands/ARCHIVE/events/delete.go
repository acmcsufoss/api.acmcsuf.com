package events

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/requests"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var DeleteEvent = &cobra.Command{
	Use:   "delete",
	Short: "Delete an event with its id",

	Run: func(cmd *cobra.Command, args []string) {
		var uuidVal string
		cmd.Flags().Set("id", uuidVal)
		err := huh.NewForm().Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		err = huh.NewInput().
			Title("ACMCSUF-CLI Event Delete:").
			Description("Please enter the event's uuid:").
			Prompt("> ").
			Value(&uuidVal).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		cmd.Flags().Set("id", uuidVal)
		id, _ := cmd.Flags().GetString("id")
		deleteEvent(id, config.Cfg)
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
		fmt.Println(err)
		return
	}

	deleteURL := baseURL.JoinPath(fmt.Sprint("/v1/events/", id))

	// ----- Delete Request -----
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

	// ----- Read Response Info -----
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
