package announcements

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var DeleteAnnouncements = &cobra.Command{
	Use:   "delete",
	Short: "delete an announcement by its id",

	Run: func(cmd *cobra.Command, args []string) {
		var flagsChosen []string
		var uuidVal string
		huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					//Ask the user what commands they want to use.
					Title("ACMCSUF-CLI Announcement Delete").
					Description("Choose a command(s). Note: Use spacebar to select and if done click enter.\nTo skip, simply click enter.").
					Options(
						huh.NewOption("Change Host", "host"),
						huh.NewOption("Change Port", "port"),
					).
					Value(&flagsChosen),
			),
		).Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Announcement Delete:").
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
					Title("ACMCSUF-CLI Announcement Delete:").
					Description("Please enter the custom host:").
					Prompt("> ").
					Value(&hostVal).
					Run()
				cmd.Flags().Set("host", hostVal)
			case "port":
				huh.NewInput().
					Title("ACMCSUF-CLI Announcement Delete:").
					Description("Please enter the custom port:").
					Prompt("> ").
					Value(&portVal).
					Run()
				cmd.Flags().Set("port", portVal)
			}
			_ = index
		}
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		uuid, _ := cmd.Flags().GetString("id")

		deleteAnnouncement(host, port, uuid)
	},
}

func init() {

	// Url flags
	DeleteAnnouncements.Flags().String("host", "127.0.0.1", "set a custom host")
	DeleteAnnouncements.Flags().String("port", "8080", "set a custom port")
	DeleteAnnouncements.Flags().String("id", "", "delete an announcement by its id")

}

func deleteAnnouncement(host string, port string, id string) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	if id == "" {
		fmt.Println("ID is required to delete an announcement! Please use the --id flag")
		return
	}

	// ----- Constructing Url -----
	host = fmt.Sprint(host, ":", port)
	path := fmt.Sprint("v1/announcements/", id)

	deleteUrl := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// ----- Delete -----
	request, err := http.NewRequest(http.MethodDelete, deleteUrl.String(), nil)
	if err != nil {
		fmt.Println("error with delete request:", err)
		return
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("error with delete response:", err)
		return
	}

	if response == nil {
		fmt.Println("no response received")
		return
	}

	defer response.Body.Close()

	// ----- Read Response Information -----
	fmt.Println("Request status:", response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error with reading body:", err)
		return
	}

	fmt.Println(string(body))
}
