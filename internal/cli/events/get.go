package events

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var GetEvent = &cobra.Command{
	Use:   "get",
	Short: "Get events",

	Run: func(cmd *cobra.Command, args []string) {
		blankUUID := ""
		cmd.Flags().Set("id", blankUUID)
		var flagsChosen []string
		err := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					//Ask the user what commands they want to use.
					Title("ACMCSUF-CLI Event Get").
					Description("Choose a command(s). Note: Use spacebar to select and if done click enter.\nTo get all events, simply click enter.").
					Options(
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
		for _, flag := range flagsChosen {
			var uuidVal string
			switch flag {
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
	req, err := http.NewRequest(http.MethodGet, getURL.String(), nil)
	if err != nil {
		fmt.Println("Error getting the request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: couldn't make GET request: %v", err)
	}
	defer resp.Body.Close()

	// ----- Read Response Information -----
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Response status:", resp.Status)
		return
	}

	if id == "" {
		var getPayload []models.CreateEventParams
		err = json.NewDecoder(resp.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body without id:", err)
			return
		}

		for i := range getPayload {
			fmt.Println(utils.PrintStruct(getPayload[i]))
		}
	} else {
		var getPayload models.CreateEventParams
		err = json.NewDecoder(resp.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body with id:", err)
			return
		}

		fmt.Println(utils.PrintStruct(getPayload))
	}
}
