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

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var PostOfficer = &cobra.Command{
	Use:   "post [flags]",
	Short: "Post a new officer",

	Run: func(cmd *cobra.Command, args []string) {
		var payload models.CreateOfficerParams

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

		postOfficer(&payload, &changedFlags, host, port)
	},
}

func init() {
	// Url flags
	PostOfficer.Flags().String("host", "127.0.0.1", "Set a custom host")
	PostOfficer.Flags().String("port", "8080", "Set a custom port")

	// Officer flags
	PostOfficer.Flags().StringP("uuid", "u", "", "Set uuid of this officer")
	PostOfficer.Flags().StringP("name", "n", "", "Set the full name of this officer")
	PostOfficer.Flags().StringP("picture", "p", "", "Set the picture of this officer")
	PostOfficer.Flags().StringP("github", "g", "", "Set the github of this officer")
	PostOfficer.Flags().StringP("discord", "d", "", "Set the discord of this officer")
}

func postOfficer(payload *models.CreateOfficerParams, cf *officerFlags, host, port string) {
	scanner := bufio.NewScanner(os.Stdin)

	// uuid
	for {
		if cf.uuid {
			break
		}

		fmt.Println("Please enter officer's uuid:")
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

		fmt.Println("Please enter the picture link for officer:")
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

		fmt.Println("Please enter the github link for officer:")
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

		fmt.Println("Please enter the discord link for officer:")
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
		fmt.Println("Is your officer data correct? If not, type n or no.")
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

	host = fmt.Sprint(host, ":", port)
	path := "v1/officers"

	postURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

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
