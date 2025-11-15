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

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var PutOfficer = &cobra.Command{
	Use:   "put --id <uuid> [flags]",
	Short: "update an existing officer by id",

	Run: func(cmd *cobra.Command, args []string) {
		payload := models.UpdateOfficerParams{}

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

		putOfficer(host, port, id, &payload, flags)
	},
}

func init() {
	PutOfficer.Flags().String("host", "127.0.0.1", "Set a custom host")
	PutOfficer.Flags().String("port", "8080", "Set a custom port")

	PutOfficer.Flags().String("id", "", "Officer ID to update")

	PutOfficer.Flags().String("fullname", "", "Change full name")
	PutOfficer.Flags().String("picture", "", "Change picture URL")
	PutOfficer.Flags().String("github", "", "Change GitHub username")
	PutOfficer.Flags().String("discord", "", "Change Discord tag")
	PutOfficer.Flags().String("uuid", "", "Change uuid")

	PutOfficer.MarkFlagRequired("id")
}

func putOfficer(host, port, id string, payload *models.UpdateOfficerParams, flags officerFlags) {
	if id == "" {
		fmt.Println("Officer id required for put! Use --id")
		return
	}

	// construct url
	hostPort := fmt.Sprint(host, ":", port)
	path := "v1/officers/" + id

	u := &url.URL{
		Scheme: "http",
		Host:   hostPort,
		Path:   path,
	}

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
		fmt.Println("Is the officer data correct? (y/n)")
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
