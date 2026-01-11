package announcements

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/config"
	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/requests"
	"github.com/spf13/cobra"
)

var PutAnnouncements = &cobra.Command{
	Use:   "put --id <uuid>",
	Short: "update an existing announcement by its id",

	Run: func(cmd *cobra.Command, args []string) {
		payload := models.UpdateAnnouncementParams{}

		id, _ := cmd.Flags().GetString("id")

		payload.Uuid, _ = cmd.Flags().GetString("uuid")
		visibilityString, _ := cmd.Flags().GetString("visibility")
		announceAtString, _ := cmd.Flags().GetString("announceat")

		channelIdString, _ := cmd.Flags().GetString("channelid")
		messageIdString, _ := cmd.Flags().GetString("messageid")

		payload.Visibility = utils.StringtoNullString(visibilityString)
		payload.DiscordChannelID = utils.StringtoNullString(channelIdString)
		payload.DiscordMessageID = utils.StringtoNullString(messageIdString)

		if announceAtString != "" {
			announceAtUnix, err := utils.ByteSlicetoUnix([]byte(announceAtString))
			if err != nil {
				fmt.Println(err)
				return
			}
			payload.AnnounceAt = utils.Int64toNullInt64(announceAtUnix)
		}

		changedFlags := announcementFlags{
			id:         cmd.Flags().Lookup("uuid").Changed,
			visibility: cmd.Flags().Lookup("visibility").Changed,
			announceat: cmd.Flags().Lookup("announceat").Changed,
			channelid:  cmd.Flags().Lookup("channelid").Changed,
			messageid:  cmd.Flags().Lookup("messageid").Changed,
		}

		putAnnouncements(id, &payload, changedFlags, config.Cfg)
	},
}

func init() {
	PutAnnouncements.Flags().String("id", "", "Get an announcement by its id")

	// Payload flags
	PutAnnouncements.Flags().String("uuid", "", "Change this announcement's uuid")
	PutAnnouncements.Flags().StringP("announceat", "a", "", "Change this announcement's announce at")

	PutAnnouncements.Flags().StringP("visibility", "v", "", "Change this announcement's visibility")
	PutAnnouncements.Flags().StringP("channelid", "c", "", "Change this announcement's discord channel id")
	PutAnnouncements.Flags().StringP("messageid", "m", "", "Change this announcement's discord message id")

	PutAnnouncements.MarkFlagRequired("id")
}

func putAnnouncements(id string, payload *models.UpdateAnnouncementParams, changedFlags announcementFlags, cfg *config.Config) {
	baseURL := &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	if err := utils.CheckConnection(baseURL.JoinPath("/health").String()); err != nil {
		fmt.Println(err)
		return
	}

	// ----- Check if Id was Given -----
	if id == "" {
		fmt.Println("Announcement id required for put! Please use the --id flag")
		return
	}

	// ----- Retrieving old Announcement -----
	// ----- Constructing Url -----
	oldPayloadUrl := baseURL.JoinPath("v1/announcements/", id)

	// ----- Get the Announcement We Want to Update -----
	client := http.Client{}

	getReq, err := requests.NewRequestWithAuth(http.MethodGet, oldPayloadUrl.String(), nil)
	if err != nil {
		fmt.Printf("Error retrieveing %s: %s", payload.Uuid, err)
		return
	}
	requests.AddOrigin(getReq)

	getRes, err := client.Do(getReq)
	if err != nil {
		fmt.Println("error getting old payload", err)
		return
	}
	defer getRes.Body.Close()

	if getRes == nil {
		fmt.Println("no response received")
		return
	}
	defer getRes.Body.Close()

	body, err := io.ReadAll(getRes.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var oldPayload models.CreateAnnouncementParams
	err = json.Unmarshal(body, &oldPayload)
	if err != nil {
		fmt.Println("error unmarshaling previous announcement data:", err)
		return
	}

	// ----- Prompt User for New Values -----
	scanner := bufio.NewScanner(os.Stdin)

	// ----- Uuid -----
	// Will probably remove uuid from all put cli soon
	for {
		if payload.Uuid == "" {
			changeUuid, err := utils.ChangePrompt("uuid", oldPayload.Uuid, scanner, "announcment")
			if err != nil {
				fmt.Println("error with changing uuid:", err)
				continue
			}
			if changeUuid != nil {
				payload.Uuid = string(changeUuid)
			} else {
				payload.Uuid = oldPayload.Uuid
			}
		}

		break
	}

	// ----- Visibility -----
	for {
		if changedFlags.visibility {
			break
		}
		changeVisibility, err := utils.ChangePrompt("visibility", oldPayload.Visibility, scanner, "announcment")
		if err != nil {
			fmt.Println("error with changing visibility:", err)
			continue
		}
		if changeVisibility != nil {
			payload.Visibility = utils.StringtoNullString(string(changeVisibility))
		} else {
			payload.Visibility = utils.StringtoNullString(oldPayload.Visibility)
		}

		break
	}

	// ----- Announce At ------
	for {
		if changedFlags.announceat {
			break
		}
		oldAnnounceAt := utils.FormatUnix(oldPayload.AnnounceAt)

		// Yah this might be a little sloppy in the terminal. forgive me.
		changeAnnounceAt, err := utils.ChangePrompt("announce at (Note: format for new announcment is \"01/02/06 03:04PM\")", oldAnnounceAt, scanner, "announcment")
		if err != nil {
			fmt.Println("error with changing announce at:", err)
			continue
		}
		if changeAnnounceAt != nil {
			announceAtInt64, err := utils.ByteSlicetoUnix(changeAnnounceAt)
			if err != nil {
				fmt.Println(err)
				continue
			}
			payload.AnnounceAt = utils.Int64toNullInt64(announceAtInt64)
		} else {
			payload.AnnounceAt = utils.Int64toNullInt64(oldPayload.AnnounceAt)
		}

		break
	}

	// ----- Discord Channel ID -----
	for {
		if changedFlags.channelid {
			break
		}

		changeDiscordChannelID, err := utils.ChangePrompt("discord channel id", oldPayload.DiscordChannelID.String, scanner, "announcment")
		if err != nil {
			fmt.Println("error with changing :", err)
			continue
		}
		if changeDiscordChannelID != nil {
			payload.DiscordChannelID = utils.StringtoNullString(string(changeDiscordChannelID))
		} else {
			payload.DiscordChannelID = oldPayload.DiscordChannelID
		}

		break
	}

	// ----- Discord Message ID -----
	for {
		if changedFlags.messageid {
			break
		}
		changeDiscordMessageID, err := utils.ChangePrompt("discord message id", oldPayload.DiscordMessageID.String, scanner, "announcment")
		if err != nil {
			fmt.Println("error with changing :", err)
			continue
		}
		if changeDiscordMessageID != nil {
			payload.DiscordMessageID = utils.StringtoNullString(string(changeDiscordMessageID))
		} else {
			payload.DiscordMessageID = oldPayload.DiscordMessageID
		}

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

	// ----- Marshal Payload to Json -----
	jsonPayload, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	// ----- Put Payload -----

	putRequest, err := requests.NewRequestWithAuth(http.MethodPut, oldPayloadUrl.String(), bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Problem with PUT:", err)
		return
	}
	requests.AddOrigin(putRequest)

	putResponse, err := client.Do(putRequest)
	if err != nil {
		fmt.Println("Error with response:", err)
		return
	}

	if putResponse == nil {
		fmt.Println("no response received")
		return
	}

	defer putResponse.Body.Close()

	// ----- Reading Response Status -----
	fmt.Println("PUT status:", putResponse.Status)

	body, err = io.ReadAll(putResponse.Body)
	if err != nil {
		fmt.Println("Error with body:", err)
		return
	}

	fmt.Println(string(body))
}
