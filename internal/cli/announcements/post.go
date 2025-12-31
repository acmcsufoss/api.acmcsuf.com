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

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/spf13/cobra"
)

var PostAnnouncement = &cobra.Command{
	Use:   "post",
	Short: "post a new announcement",

	Run: func(cmd *cobra.Command, args []string) {
		payload := models.CreateAnnouncementParams{}

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

	scanner := bufio.NewScanner(os.Stdin)

	// ----- Uuid -----
	for {
		if changedFlags.id {
			break
		}

		fmt.Println("Please enter the announcement's uuid:")
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
		fmt.Println("Please enter this announcement's visibility:")
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

		fmt.Println("Please enter the \"announce at\" of the announcement in the following format:\n[Month]/[Day]/[Year] [Hour]:[Minutes][PM | AM]")
		fmt.Println("For example: \x1b[93m01/02/06 03:04PM\x1b[0m")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading anounce at:", err)
			continue
		}

		announceatBuffer := scanner.Bytes()

		var err error
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

		fmt.Println("Please enter this announcement's discord channel id:")
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

		fmt.Println("Please enter this announcement's message id:")
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
		fmt.Println("Is your event data correct? If not, type n or no.")
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
	response, err := http.Post(postURL.String(), "application/json", strings.NewReader(string(jsonPayload)))
	if err != nil {
		fmt.Println("error with post:", err)
		return
	}

	if response == nil {
		fmt.Println("no response received")
		return
	}
	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading body:", err)
		return
	}

	fmt.Println(string(body))
}
