package announcements

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"fmt"

 	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/requests"
)

var PostAnnouncement = &cobra.Command{
	Use:   "post",
	Short: "post a new announcement",

	Run: func(cmd *cobra.Command, args []string) {
		payload := models.CreateAnnouncementParams{}
		var flagsChosen []string
		err := huh.NewForm(
			huh.NewGroup(
				huh.NewMultiSelect[string]().
					//Ask the user what commands they want to use.
					Title("ACMCSUF-CLI Announcement Post").
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
					Title("ACMCSUF-CLI Announcement Post:").
					Description("Please enter the custom host:").
					Prompt("> ").
					Value(&hostVal).
					Run()
				cmd.Flags().Set("host", hostVal)
			case "port":
				err = huh.NewInput().
					Title("ACMCSUF-CLI Announcement Post:").
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
		payload.Uuid, _ = cmd.Flags().GetString("uuid")
		payload.Visibility, _ = cmd.Flags().GetString("visibility")
		announceString, _ := cmd.Flags().GetString("announceat")

		channelIdString, _ := cmd.Flags().GetString("channelid")
		messageIdString, _ := cmd.Flags().GetString("messageid")

		payload.DiscordChannelID = utils.StringtoNullString(channelIdString)
		payload.DiscordMessageID = utils.StringtoNullString(messageIdString)

		if announceString != "" {
			var err error
			payload.AnnounceAt, err = utils.ByteSlicetoUnix([]byte(announceString))
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		changedFlags := announcementFlags{
			id:         cmd.Flags().Lookup("uuid").Changed,
			visibility: cmd.Flags().Lookup("visibility").Changed,
			announceat: cmd.Flags().Lookup("announceat").Changed,
			channelid:  cmd.Flags().Lookup("channelid").Changed,
			messageid:  cmd.Flags().Lookup("messageid").Changed,
		}

		postAnnouncement(&payload, changedFlags, config.Cfg)
	},
}

func init() {
	// Payload flags
	PostAnnouncement.Flags().String("uuid", "", "Set this announcement's id")
	PostAnnouncement.Flags().StringP("visibility", "v", "", "Set this announcement's visibility")
	PostAnnouncement.Flags().StringP("announceat", "a", "", "Set this announcement's announce at")

	PostAnnouncement.Flags().StringP("channelid", "c", "", "Set this announcement's channel id")
	PostAnnouncement.Flags().StringP("messageid", "m", "", "Set this announcement's message id")
}

func postAnnouncement(payload *models.CreateAnnouncementParams, changedFlags announcementFlags, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	// ----- Uuid -----
	for {
		if changedFlags.id {
			break
		}
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
		if changedFlags.visibility {
			break
		}
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
		if changedFlags.announceat {
			break
		}
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
		if changedFlags.channelid {
			break
		}
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
		if changedFlags.messageid {
			break
		}
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

	// ----- Marshalling to Json -----
	jsonPayload, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println("error formating payload to json:", err)
		return
	}

	// ----- Constructing the Url -----
	postURL := baseURL.JoinPath("v1/announcements")

	fmt.Println(postURL.String())
	// ----- Post -----
	client := http.Client{}
	req, err := requests.NewRequestWithAuth(http.MethodPost, postURL.String(), strings.NewReader(string(jsonPayload)))
	if err != nil {
		fmt.Println("error with post:", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("error with requesting post", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Response status:", res.Status)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading body:", err)
		return
	}

	fmt.Println(string(body))
}
