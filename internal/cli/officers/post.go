package officers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var PostOfficer = &cobra.Command{
	Use:   "post [flags]",
	Short: "Post a new officer",

	Run: func(cmd *cobra.Command, args []string) {
		var payload models.CreateOfficerParams
		var flagsChosen []string
		err := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					//Ask the user what commands they want to use.
					Title("ACMCSUF-CLI Board Post").
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
		for index, flag := range flagsChosen {
			var hostVal string
			var portVal string
			switch flag {
			case "host":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Board Post:").
					Description("Please enter the custom host:").
					Prompt("> ").
					Value(&hostVal).
					Run()
				cmd.Flags().Set("host", hostVal)
			case "port":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Board Post:").
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
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		payload.Uuid, _ = cmd.Flags().GetString("uuid")
		payload.FullName, _ = cmd.Flags().GetString("name")
		pic, _ := cmd.Flags().GetString("picture")
		payload.Picture = utils.StringtoNullString(pic)
		git, _ := cmd.Flags().GetString("github")
		payload.Github = utils.StringtoNullString(git)
		disc, _ := cmd.Flags().GetString("discord")
		payload.Discord = utils.StringtoNullString(disc)

		changedFlags := officerFlags{
			uuid:     cmd.Flags().Lookup("uuid").Changed,
			fullname: cmd.Flags().Lookup("name").Changed,
			picture:  cmd.Flags().Lookup("picture").Changed,
			github:   cmd.Flags().Lookup("github").Changed,
			discord:  cmd.Flags().Lookup("discord").Changed,
		}

		postOfficer(&payload, &changedFlags, config.Cfg)
	},
}

func init() {
	// Officer flags
	PostOfficer.Flags().StringP("uuid", "u", "", "Set uuid of this officer")
	PostOfficer.Flags().StringP("name", "n", "", "Set the full name of this officer")
	PostOfficer.Flags().StringP("picture", "p", "", "Set the picture of this officer")
	PostOfficer.Flags().StringP("github", "g", "", "Set the github of this officer")
	PostOfficer.Flags().StringP("discord", "d", "", "Set the discord of this officer")
}

func postOfficer(payload *models.CreateOfficerParams, cf *officerFlags, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	// uuid
	for {
		if cf.uuid {
			break
		}

		var uuid string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Board Post:").
			Description("Please enter officer's uuid:").
			Prompt("> ").
			Value(&uuid).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(uuid))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.Uuid = string(scanner.Bytes())
		break
	}

	// full name
	for {
		if cf.fullname {
			break
		}

		var fullName string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Board Post:").
			Description("Please enter officer's full name:").
			Prompt("> ").
			Value(&fullName).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(fullName))
		fmt.Println("Please enter the officer's full name:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.FullName = string(scanner.Bytes())
		break
	}

	// picture
	for {
		if cf.picture {
			break
		}

		var picLink string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Board Post:").
			Description("Please enter the picture link for officer:").
			Prompt("> ").
			Value(&picLink).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(picLink))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.Picture = utils.StringtoNullString(string(scanner.Bytes()))
		break
	}

	// github
	for {
		if cf.github {
			break
		}

		var githubLink string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Board Post:").
			Description("Please enter the github link for officer:").
			Prompt("> ").
			Value(&githubLink).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(githubLink))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.Github = utils.StringtoNullString(string(scanner.Bytes()))
		break
	}

	// discord
	for {
		if cf.discord {
			break
		}

		var discordLink string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Board Post:").
			Description("Please enter the discord link for officer").
			Prompt("> ").
			Value(&discordLink).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(discordLink))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		payload.Discord = utils.StringtoNullString(string(scanner.Bytes()))
		break

	}

	// confirmation
	for {
		var option string
		description := "Is your board data correct?\n" + utils.PrintStruct(payload)
		err := huh.NewSelect[string]().
			Title("ACMCSUF-CLI Board Post:").
			Description(description).
			Options(
				huh.NewOption("Yes", "yes"),
				huh.NewOption("No", "n"),
			).
			Value(&option).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(option))
		utils.PrintStruct(payload)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}

		confirmationBuffer := scanner.Bytes()
		confirmationBool, err := utils.YesOrNo(confirmationBuffer, scanner)
		if err != nil {
			fmt.Println("error with reading confirmation:", err)
		}
		if !confirmationBool {
			// Sorry :(
			return
		} else {
			break
		}
	}

	// marshal to json, and prepare url
	jsonPayload, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println("error formating payload to json: ", err)
		return
	}

	postURL := baseURL.JoinPath("v1/board/officers/")

	// post payload
	response, err := http.Post(postURL.String(), "application/json", strings.NewReader(string(jsonPayload)))
	if err != nil {
		fmt.Println("error with post: ", err)
		return
	}

	if response == nil {
		fmt.Println("error, no response recieved")
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading body: ", err)
		return
	}

	fmt.Println(string(body))
}
