package officers

import (
	"bufio"
	"bytes"
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

var PutOfficer = &cobra.Command{
	Use:   "put --id <uuid> [flags]",
	Short: "update an existing officer by id",

	Run: func(cmd *cobra.Command, args []string) {
		payload := models.UpdateOfficerParams{}
		var flagsChosen []string
		var uuidVal string
		huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					//Ask the user what commands they want to use.
					Title("ACMCSUF-CLI Board Put").
					Description("Choose a command(s). Note: Use spacebar to select and if done click enter.\nTo skip, simply click enter.").
					Options(
						huh.NewOption("Change Host", "host"),
						huh.NewOption("Change Port", "port"),
					).
					Value(&flagsChosen),
			),
		).Run()
		huh.NewInput().
			Title("ACMCSUF-CLI Board Put:").
			Description("Please enter the officer's ID:").
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
					Title("ACMCSUF-CLI Board Put:").
					Description("Please enter the custom host:").
					Prompt("> ").
					Value(&hostVal).
					Run()
				cmd.Flags().Set("host", hostVal)
			case "port":
				huh.NewInput().
					Title("ACMCSUF-CLI Board Put:").
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
		id, _ := cmd.Flags().GetString("id")

		fullname, _ := cmd.Flags().GetString("fullname")
		picture, _ := cmd.Flags().GetString("picture")
		github, _ := cmd.Flags().GetString("github")
		discord, _ := cmd.Flags().GetString("discord")
		uuid, _ := cmd.Flags().GetString("uuid")

		payload.FullName = fullname
		payload.Picture = utils.StringtoNullString(picture)
		payload.Github = utils.StringtoNullString(github)
		payload.Discord = utils.StringtoNullString(discord)
		payload.Uuid = uuid

		flags := officerFlags{
			fullname: cmd.Flags().Lookup("fullname").Changed,
			picture:  cmd.Flags().Lookup("picture").Changed,
			github:   cmd.Flags().Lookup("github").Changed,
			discord:  cmd.Flags().Lookup("discord").Changed,
			uuid:     cmd.Flags().Lookup("uuid").Changed,
		}

		putOfficer(id, &payload, flags, config.Cfg)
	},
}

func init() {
	PutOfficer.Flags().String("id", "", "Officer ID to update")

	PutOfficer.Flags().String("fullname", "", "Change full name")
	PutOfficer.Flags().String("picture", "", "Change picture URL")
	PutOfficer.Flags().String("github", "", "Change GitHub username")
	PutOfficer.Flags().String("discord", "", "Change Discord tag")
	PutOfficer.Flags().String("uuid", "", "Change uuid")

	PutOfficer.MarkFlagRequired("id")
}

func putOfficer(id string, payload *models.UpdateOfficerParams, flags officerFlags, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	if id == "" {
		fmt.Println("Officer id required for put! Use --id")
		return
	}

	// construct url
	u := baseURL.JoinPath("v1/board/officers/", id)

	// getting old officer
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("error retrieving %s: %s\n", id, err)
		return
	}
	if resp == nil {
		fmt.Println("no response received")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		return
	}

	var old models.CreateOfficerParams
	if err := json.Unmarshal(body, &old); err != nil {
		fmt.Println("error unmarshaling previous officer data:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	// full name
	for {
		if flags.fullname {
			break
		}

		change, err := utils.ChangePrompt("full name", old.FullName, scanner, "officer")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if change != nil {
			payload.FullName = string(change)
		} else {
			payload.FullName = old.FullName
		}
		break
	}

	// picture
	for {
		if flags.picture {
			break
		}

		change, err := utils.ChangePrompt("picture", old.Picture.String, scanner, "officer")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if change != nil {
			payload.Picture = utils.StringtoNullString(string(change))
		} else {
			payload.Picture = old.Picture
		}
		break
	}

	// github
	for {
		if flags.github {
			break
		}

		change, err := utils.ChangePrompt("github", old.Github.String, scanner, "officer")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if change != nil {
			payload.Github = utils.StringtoNullString(string(change))
		} else {
			payload.Github = old.Github
		}
		break
	}

	// discord
	for {
		if flags.discord {
			break
		}

		change, err := utils.ChangePrompt("discord", old.Discord.String, scanner, "officer")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if change != nil {
			payload.Discord = utils.StringtoNullString(string(change))
		} else {
			payload.Discord = old.Discord
		}
		break
	}

	// uuid
	for {
		if flags.uuid {
			break
		}

		change, err := utils.ChangePrompt("uuid", old.Uuid, scanner, "officer")
		if err != nil {
			fmt.Println(err)
			continue
		}
		if change != nil {
			payload.Uuid = string(change)
		} else {
			payload.Uuid = old.Uuid
		}
		break
	}

	// Confirm
	for {
		var option string
		huh.NewSelect[string]().
			Title("ACMCSUF-CLI Board Put:").
			Description("Is your event data correct? If not, type n or no.").
			Options(
				huh.NewOption("Yes", "yes"),
				huh.NewOption("No", "n"),
			).
			Value(&option).
			Run()
		scanner := bufio.NewScanner(strings.NewReader(option))
		utils.PrintStruct(payload)
		scanner.Scan()
		confirmation := scanner.Bytes()

		ok, err := utils.YesOrNo(confirmation, scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if !ok {
			return
		}
		break
	}

	// marshal payload
	jsonPayload, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	// PUT
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, u.String(), bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Problem with PUT:", err)
		return
	}

	putResp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error with response:", err)
		return
	}
	if putResp == nil {
		fmt.Println("no response received")
		return
	}
	defer putResp.Body.Close()

	fmt.Println("PUT status:", putResp.Status)

	body, err = io.ReadAll(putResp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	fmt.Println(string(body))
}
