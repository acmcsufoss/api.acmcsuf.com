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


	"github.com/acmcsufoss/api.acmcsuf.com/utils/cli"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/convert"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/dbtypes"

	"github.com/spf13/cobra"
)

var PutAnnouncements = &cobra.Command{
	Use:   "put --id <uuid>",
	Short: "update an existing announcement by its id",

	Run: func(cmd *cobra.Command, args []string) {
		payload := UpdateAnnouncement{}

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		id, _ := cmd.Flags().GetString("id")

		payload.Uuid, _ = cmd.Flags().GetString("uuid")
		visibilityString, _ := cmd.Flags().GetString("visibility")
		announceAtString, _ := cmd.Flags().GetString("announceat")

		channelIdString, _ := cmd.Flags().GetString("channelid")
		messageIdString, _ := cmd.Flags().GetString("messageid")

		payload.Visibility = dbtypes.StringtoNullString(visibilityString)
		payload.DiscordChannelID = dbtypes.StringtoNullString(channelIdString)
		payload.DiscordMessageID = dbtypes.StringtoNullString(messageIdString)

		if announceAtString != "" {
			announceAtUnix, err := convert.ByteSlicetoUnix([]byte(announceAtString))
			if err != nil {
				fmt.Println(err)
				return
			}
			payload.AnnounceAt = dbtypes.Int64toNullInt64(announceAtUnix)
		}
    
		putAnnouncements(host, port, id, &payload)
	},
}

func init() {

	// Url flags
	PutAnnouncements.Flags().String("host", "127.0.0.1", "Set a custom host")
	PutAnnouncements.Flags().String("port", "8080", "Set a custom port")

	PutAnnouncements.Flags().String("id", "", "Get an announcement by its id")

	// Payload flags
	PutAnnouncements.Flags().String("uuid", "", "Change this announcement's uuid")
	PutAnnouncements.Flags().StringP("announceat", "a", "", "Change this announcement's announce at")

	PutAnnouncements.Flags().StringP("visibility", "v", "", "Change this announcement's visibility")
	PutAnnouncements.Flags().StringP("channelid", "c", "", "Change this announcement's discord channel id")
	PutAnnouncements.Flags().StringP("messageid", "m", "", "Change this announcement's discord message id")

	PutAnnouncements.MarkFlagRequired("id")

}

func putAnnouncements(host string, port string, id string, payload *UpdateAnnouncement) {
	// ----- Check if Id was Given -----
	if id == "" {
		fmt.Println("Announcement id required for put! Please use the --id flag")
		return
	}

	// ----- Retrieving old Announcement -----
	// ----- Constructing Url -----
	host = fmt.Sprint(host, ":", port)
	path := fmt.Sprint("announcements/", id)

	oldPayloadUrl := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// ----- Get the Announcement We Want to Update -----
	response, err := http.Get(oldPayloadUrl.String())
	if err != nil {
		fmt.Printf("Error retrieveing %s: %s", payload.Uuid, err)
		return
	}

	if response == nil {
		fmt.Println("no response received")
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var oldPayload CreateAnnouncement
	err = json.Unmarshal(body, &oldPayload)
	if err != nil {
		fmt.Println("error unmarshaling previous announcement data:", err)
		return
	}


	// ----- Prompt User for New Values -----
	scanner := bufio.NewScanner(os.Stdin)

	// ----- Uuid -----
	// Known issue: Despite the response going through and output saying uuid has been updated
	// it does not actually update
	for {
		if payload.Uuid == "" {
			changeUuid, err := cli.ChangePrompt("uuid", oldPayload.Uuid, scanner)
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
		if payload.Visibility.String == "" {
			changeVisibility, err := cli.ChangePrompt("visibility", oldPayload.Visibility, scanner)
			if err != nil {
				fmt.Println("error with changing visibility:", err)
				continue
			}
			if changeVisibility != nil {
				payload.Visibility = dbtypes.StringtoNullString(string(changeVisibility))
			} else {
				payload.Visibility = dbtypes.StringtoNullString(oldPayload.Visibility)
			}
		}

		break
	}

	// ----- Announce At ------
	for {
		if payload.AnnounceAt.Int64 == 0 {
			oldAnnounceAt := cli.FormatUnix(oldPayload.AnnounceAt)

			// Yah this might be a little sloppy in the terminal. forgive me.
			changeAnnounceAt, err := cli.ChangePrompt("announce at (Note: format for new announcment is \"03:04:05PM 01/02/06\")", oldAnnounceAt, scanner)
			if err != nil {
				fmt.Println("error with changing announce at:", err)
				continue
			}
			if changeAnnounceAt != nil {
				announceAtInt64, err := convert.ByteSlicetoUnix(changeAnnounceAt)
				if err != nil {
					fmt.Println(err)
					continue
				}
				payload.AnnounceAt = dbtypes.Int64toNullInt64(announceAtInt64)
			} else {
				payload.AnnounceAt = dbtypes.Int64toNullInt64(oldPayload.AnnounceAt)
			}
		}

		break
	}

	// ----- Discord Channel ID -----
	for {
		if payload.DiscordChannelID.String == "" {
			changeDiscordChannelID, err := cli.ChangePrompt("discord channel id", oldPayload.DiscordChannelID.String, scanner)
			if err != nil {
				fmt.Println("error with changing :", err)
				continue
			}
			if changeDiscordChannelID != nil {
				payload.DiscordChannelID = dbtypes.StringtoNullString(string(changeDiscordChannelID))
			} else {
				payload.DiscordChannelID = oldPayload.DiscordChannelID
			}
		}

		break
	}

	// ----- Discord Message ID -----
	for {
		if payload.DiscordMessageID.String == "" {
			changeDiscordMessageID, err := cli.ChangePrompt("discord message id", oldPayload.DiscordMessageID.String, scanner)
			if err != nil {
				fmt.Println("error with changing :", err)
				continue
			}
			if changeDiscordMessageID != nil {
				payload.DiscordMessageID = dbtypes.StringtoNullString(string(changeDiscordMessageID))
			} else {
				payload.DiscordMessageID = oldPayload.DiscordMessageID
			}
		}

		break
	}
	// ----- Confirmation -----
	for {
		fmt.Println("Is your event data correct? If not, type n or no.")
		cli.PrintStruct(payload)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}

		confirmationBuffer := scanner.Bytes()
		confirmationBool, err := cli.YesOrNo(confirmationBuffer, scanner)
		if err != nil {
			fmt.Println("error with reading confirmation:", err)
		}
		if !confirmationBool {
			// Sorry :(
			return
		}
		break

	}

	// ----- Marshal Payload to Json -----
	jsonPayload, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	// ----- Put Payload -----
	client := &http.Client{}

	putRequest, err := http.NewRequest(http.MethodPut, oldPayloadUrl.String(), bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Problem with PUT:", err)
		return
	}

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
