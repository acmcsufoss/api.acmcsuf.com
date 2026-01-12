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
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var GetEvent = &cobra.Command{
	Use:   "get",
	Short: "Get events",

	Run: func(cmd *cobra.Command, args []string) {
		var flagsChosen []string
		err := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					//Ask the user what commands they want to use.
					Title("ACMCSUF-CLI Event Get").
					Description("Choose a command(s). Note: Use spacebar to select and if done click enter.\nTo get all events, simply click enter.").
					Options(
						huh.NewOption("Change Host", "host"),
						huh.NewOption("Change Port", "port"),
						huh.NewOption("Get Specific ID", "id"),
					).
					Value(&flagsChosen),
			),
		).Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		for index, flag := range flagsChosen {
			var hostVal string
			var portVal string
			var uuidVal string
			switch flag {
			case "host":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Event Get:").
					Description("Please enter the custom host:").
					Prompt("> ").
					Value(&hostVal).
					Run()
				cmd.Flags().Set("host", hostVal)
			case "port":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Event Get:").
					Description("Please enter the custom port:").
					Prompt("> ").
					Value(&portVal).
					Run()
				cmd.Flags().Set("port", portVal)
			case "id":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Event Get:").
					Description("Please enter the event's ID:").
					Prompt("> ").
					Value(&uuidVal).
					Run()
				cmd.Flags().Set("id", uuidVal)
			}
			if err != nil {
				if err == huh.ErrUserAborted {
					fmt.Println("User canceled the form — exiting.")
				}
				fmt.Println("Uh oh:", err)
				os.Exit(1)
			}
			_ = index
		}
		// If these where global, unexpected behavior would be expected :(
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
	req, err := http.NewRequest("GET", getURL.String(), nil)
	if err != nil {
		fmt.Println("Error getting the request:", err)
		return
	}

	client := &http.Client{}
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
