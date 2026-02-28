package officers

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var GetOfficers = &cobra.Command{
	Use:   "get [flags]",
	Short: "Get Officers",

	Run: func(cmd *cobra.Command, args []string) {
		blankUUID := ""
		cmd.Flags().Set("id", blankUUID)
		var flagsChosen []string
		err := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					//Ask the user what commands they want to use.
					Title("ACMCSUF-CLI Officers Get").
					Description("Choose a command(s). Note: Use spacebar to select and if done click enter.\nTo get all officers, simply click enter.").
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
					Title("ACMCSUF-CLI Officers Get:").
					Description("Please enter the officer's ID:").
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

		id, _ := cmd.Flags().GetString("id")
		getOfficers(id, config.Cfg)
	},
}

func init() {
	GetOfficers.Flags().String("id", "", "Get a specific officer")
}

func getOfficers(uuid string, cfg *config.Config) {
	getUrl := config.GetBaseURL(cfg).JoinPath("v1", "announcements", uuid)

	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, getUrl.String(), nil)
	if err != nil {
		fmt.Println("Error: failed to construct request:", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: failed to send request:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP", res.Status)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: failed to read response body:", err)
		return
	}
	utils.PrettyPrintJSON(body)
}
