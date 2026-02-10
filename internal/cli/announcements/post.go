package announcements

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/api/dbmodels"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/oauth"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

var PostAnnouncement = &cobra.Command{
	Use:   "post",
	Short: "post a new announcement",

	Run: func(cmd *cobra.Command, args []string) {
		postAnnouncement(config.Cfg)
	},
}

func postAnnouncement(cfg *config.Config) {
	var payload dbmodels.CreateAnnouncementParams
	// ----- Uuid -----
	for {
		var uuid string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Announcements Post:").
			Description("Please enter the announcement's uuid:").
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
			fmt.Println("error with reading uuid:", err)
			continue
		}

		uuidBuffer := scanner.Bytes()

		payload.Uuid = string(uuidBuffer)
		break
	}

	// ----- Visibility -----
	for {
		var visibility string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Announcements Post:").
			Description("Please enter the announcement's visibility:").
			Prompt("> ").
			Value(&visibility).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(visibility))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error with reading visibility:", err)
			continue
		}

		visibilityBuffer := scanner.Bytes()
		payload.Visibility = string(visibilityBuffer)

		break
	}

	// ----- Announce at -----
	for {
		var announceAt string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Announcements Post:").
			Description("Please enter the \"announce at\" of the announcement in the following format:\n[Month]/[Day]/[Year] [Hour]:[Minutes][PM | AM]\nFor example: \x1b[93m01/02/06 03:04PM\x1b[0m").
			Prompt("> ").
			Value(&announceAt).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(announceAt))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading anounce at:", err)
			continue
		}

		announceatBuffer := scanner.Bytes()

		payload.AnnounceAt, err = utils.ByteSlicetoUnix(announceatBuffer)
		if err != nil {
			fmt.Println("error converting byte slice to unix time (of type int64):", err)
			continue
		}

		break
	}

	// ----- Discord Channel Id -----
	for {
		var discordid string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Announcements Post:").
			Description("Please enter the announcement's discord channel id:").
			Prompt("> ").
			Value(&discordid).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(discordid))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading  discord channel id:", err)
			continue
		}

		channelIdBuffer := scanner.Bytes()
		payload.DiscordChannelID = utils.StringtoNullString(string(channelIdBuffer))

		break
	}

	// ----- Discord Message Id -----
	for {
		var messageid string
		err := huh.NewInput().
			Title("ACMCSUF-CLI Announcements Post:").
			Description("Please enter the announcement's message id:").
			Prompt("> ").
			Value(&messageid).
			Run()
		if err != nil {
			if err == huh.ErrUserAborted {
				fmt.Println("User canceled the form — exiting.")
			}
			fmt.Println("Uh oh:", err)
			os.Exit(1)
		}
		scanner := bufio.NewScanner(strings.NewReader(messageid))
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading message id:", err)
			continue
		}
		messageIdBuffer := scanner.Bytes()
		payload.DiscordMessageID = utils.StringtoNullString(string(messageIdBuffer))

		break
	}

	// ----- Confirmation -----
	for {
		var option string
		description := "Is your announcement data correct?\n" + utils.PrintStruct(payload)
		err := huh.NewSelect[string]().
			Title("ACMCSUF-CLI Announcements Post:").
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

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error: could not marshal JSON:", err)
		return
	}

	postURL := config.GetBaseURL(cfg).JoinPath("v1", "announcements")
	client := http.Client{}
	req, err := oauth.NewRequestWithAuth(http.MethodPost, postURL.String(), strings.NewReader(string(jsonPayload)))
	if err != nil {
		fmt.Println("Error: could not create request:", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: could not send request:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: HTTP", res.Status)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: could not read response body:", err)
		return
	}
	utils.PrettyPrintJSON(body)
}

// TODO: Use DTO models instaad of dbmodels
func form() *dbmodels.CreateAnnouncementParams {
	return nil
}
