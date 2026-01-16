package announcements

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

var GetAnnouncement = &cobra.Command{
	Use:   "get",
	Short: "Get an announcement",

	Run: func(cmd *cobra.Command, args []string) {
		var flagsChosen []string
		err := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					//Ask the user what commands they want to use.
					Title("ACMCSUF-CLI Announcement Get").
					Description("Choose a command(s). Note: Use spacebar to select and if done click enter.\nTo get all announcements, simply click enter.").
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
					Title("ACMCSUF-CLI Announcement Get:").
					Description("Please enter the custom host:").
					Prompt("> ").
					Value(&hostVal).
					Run()
				cmd.Flags().Set("host", hostVal)
			case "port":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Announcement Get:").
					Description("Please enter the custom port:").
					Prompt("> ").
					Value(&portVal).
					Run()
				cmd.Flags().Set("port", portVal)
			case "id":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Announcement Get:").
					Description("Please enter the announcement's ID:").
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

		uuid, _ := cmd.Flags().GetString("id")
		getAnnouncement(uuid, config.Cfg)
	},
}

func init() {
	GetAnnouncement.Flags().String("id", "", "Get a specific announcement by its id")
}

func getAnnouncement(uuid string, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	// ----- Constructing the url -----
	path := "v1/announcements"
	if uuid != "" {
		path = fmt.Sprint(path, "/", uuid)
	}

	getUrl := baseURL.JoinPath(path)

	// ----- Requesting Get -----
	response, err := http.Get(getUrl.String())
	if err != nil {
		fmt.Println("error with request:", err)
		return
	}

	if response == nil {
		fmt.Println("no response received")
		return
	}

	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)

	if uuid == "" {
		var getPayload []models.CreateAnnouncementParams
		err = json.NewDecoder(response.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body without id:", err)
			return
		}

		for i := range getPayload {
			utils.PrintStruct(getPayload[i])
		}
	} else {
		var getPayload models.CreateAnnouncementParams
		err = json.NewDecoder(response.Body).Decode(&getPayload)
		if err != nil {
			fmt.Println("Failed to read response body with id:", err)
			return
		}

		utils.PrintStruct(getPayload)
	}

}
