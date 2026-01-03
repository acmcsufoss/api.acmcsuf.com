package officers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var DeleteOfficers = &cobra.Command{
	Use:   "delete --id <uuid>",
	Short: "Delete an officer with their id",

	Run: func(cmd *cobra.Command, args []string) {
		var flagsChosen []string
		var uuidVal string
		huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					//Ask the user what commands they want to use.
					Title("ACMCSUF-CLI Board Delete").
					Description("Choose a command(s). Note: Use spacebar to select and if done click enter.\nTo skip, simply click enter.").
					Options(
						huh.NewOption("Change Host", "host"),
						huh.NewOption("Change Port", "port"),
					).
					Value(&flagsChosen),
			),
		).Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Board Delete:").
			Description("Please enter the announcement's ID:").
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
					Title("ACMCSUF-CLI Board Delete:").
					Description("Please enter the custom host:").
					Prompt("> ").
					Value(&hostVal).
					Run()
				cmd.Flags().Set("host", hostVal)
			case "port":
				huh.NewInput().
					Title("ACMCSUF-CLI Board Delete:").
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

		deleteOfficer(id, host, port)
	},
}

func init() {
	DeleteOfficers.Flags().String("id", "", "Delete an officer by their id")
	DeleteOfficers.Flags().String("host", "127.0.0.1", "Set a custom host")
	DeleteOfficers.Flags().String("port", "8080", "Set a custom port")

	DeleteOfficers.MarkFlagRequired("id")
}

func deleteOfficer(id, host, port string) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	// req id
	if id == "" {
		fmt.Println("ID is required to delete!")
		return
	}

	// prepare url
	host = fmt.Sprint(host, ":", port)
	path := fmt.Sprint("v1/board/officers/", id)

	deleteURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

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
