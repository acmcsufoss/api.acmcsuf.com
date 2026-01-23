package announcements

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

 	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/requests"
)

var DeleteAnnouncements = &cobra.Command{
	Use:   "delete",
	Short: "delete an announcement by its id",

	Run: func(cmd *cobra.Command, args []string) {
		var flagsChosen []string
		var uuidVal string
		err := huh.NewForm(
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
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		err = huh.NewInput().
			Title("ACMCSUF-CLI Announcement Delete:").
			Description("Please enter the announcement's ID:").
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
		for index, flag := range flagsChosen {
			var hostVal string
			var portVal string
			switch flag {
			case "host":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Announcement Delete:").
					Description("Please enter the custom host:").
					Prompt("> ").
					Value(&hostVal).
					Run()
				cmd.Flags().Set("host", hostVal)
			case "port":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Announcement Delete:").
					Description("Please enter the custom port:").
					Prompt("> ").
					Value(&portVal).
					Run()
				cmd.Flags().Set("port", portVal)
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
		deleteAnnouncement(uuid, config.Cfg)
	},
}

func init() {
	DeleteAnnouncements.Flags().String("id", "", "delete an announcement by its id")
	DeleteAnnouncements.MarkFlagRequired("id")
}

func deleteAnnouncement(id string, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	if id == "" {
		fmt.Println("ID is required to delete an announcement! Please use the --id flag")
		return
	}

	// ----- Constructing Url -----
	deleteUrl := baseURL.JoinPath("v1/announcements/", id)

	// ----- Delete -----
	request, err := requests.NewRequestWithAuth(http.MethodDelete, deleteUrl.String(), nil)
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
	if response.StatusCode != http.StatusOK {
		fmt.Println("response status:", response.Status)
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
