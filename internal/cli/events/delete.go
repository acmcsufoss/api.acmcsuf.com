package events

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var DeleteEvent = &cobra.Command{
	Use:   "delete",
	Short: "Delete an event with its id",

	Run: func(cmd *cobra.Command, args []string) {
		var flagsChosen []string
		var uuidVal string
		huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					//Ask the user what commands they want to use.
					Title("ACMCSUF-CLI Event Delete").
					Description("Choose a command(s). Note: Use spacebar to select and if done click enter.\nTo skip, simply click enter.").
					Options(
						huh.NewOption("Change Host", "host"),
						huh.NewOption("Change Port", "port"),
					).
					Value(&flagsChosen),
			),
		).Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Event Delete:").
			Description("Please enter the event's uuid:").
			Prompt("> ").
			Value(&uuidVal).
			Run()
		cmd.Flags().Set("id", uuidVal)
		for index, flag := range flagsChosen {
			var hostVal string
			var portVal string
			switch flag {
			case "host":
				huh.NewInput().
					Title("ACMCSUF-CLI Event Delete:").
					Description("Please enter the custom host:").
					Prompt("> ").
					Value(&hostVal).
					Run()
				cmd.Flags().Set("host", hostVal)
			case "port":
				huh.NewInput().
					Title("ACMCSUF-CLI Event Delete:").
					Description("Please enter the custom port:").
					Prompt("> ").
					Value(&portVal).
					Run()
				cmd.Flags().Set("port", portVal)
			}
			_ = index
		}
		id, _ := cmd.Flags().GetString("id")
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		deleteEvent(id, host, port)
	},
}

func init() {

	// Url flags
	DeleteEvent.Flags().String("id", "", "Delete the identified event")
	DeleteEvent.Flags().String("host", "127.0.0.1", "Set a custom host")
	DeleteEvent.Flags().String("port", "8080", "Set a custom port")

	DeleteEvent.MarkFlagRequired("id")

}

func deleteEvent(id string, host string, port string) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	// ----- Check if Event Id was Given -----
	if id == "" {
		fmt.Println("Event ID is required to delete!")
		return
	}

	host = fmt.Sprint(host, ":", port)
	path := fmt.Sprint("v1/events/", id)

	deleteURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	client := &http.Client{}

	// ----- Delete Request -----
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

	// ----- Read Response Info -----
	fmt.Println("Response status:", response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading delete response body:", err)
		return
	}

	fmt.Println(string(body))
}
