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
	"strconv"
	"strings"

	"github.com/acmcsufoss/api.acmcsuf.com/utils/convert"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/dbtypes"
	"github.com/spf13/cobra"
)

var PutAnnouncements = &cobra.Command{
	Use:   "put --id <uuid>",
	Short: "update an existing announcement by it's id",

	Run: func(cmd *cobra.Command, args []string) {
		payload := UpdateAnnouncement{}

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		id, _ := cmd.Flags().GetString("id")

		payload.Uuid, _ = cmd.Flags().GetString("uuid")
		visibilityString, _ := cmd.Flags().GetString("visibility")
		announceAtInt64, _ := cmd.Flags().GetInt64("announceat")
		channelIdString, _ := cmd.Flags().GetString("channelid")
		messageIdString, _ := cmd.Flags().GetString("messageid")

		payload.Visibility = dbtypes.StringtoNullString(visibilityString)
		payload.AnnounceAt = dbtypes.Int64toNullInt64(announceAtInt64)
		payload.DiscordChannelID = dbtypes.StringtoNullString(channelIdString)
		payload.DiscordMessageID = dbtypes.StringtoNullString(messageIdString)

		putAnnouncements(host, port, id, &payload)
	},
}

func init() {

	// Url flags
	PutAnnouncements.Flags().String("host", "127.0.0.1", "Set a custom (Defaults to: 127.0.0.1)")
	PutAnnouncements.Flags().String("port", "8080", "Set a custom (Defaults to: 8080)")
	PutAnnouncements.Flags().String("id", "", "Get an announcement by it's id")

	// Payload flags
	PutAnnouncements.Flags().String("uuid", "", "Change this announcement's uuid")
	PutAnnouncements.Flags().Int64P("announceat", "a", 0, "Change this announcement's announce at")
	PutAnnouncements.Flags().StringP("visibility", "v", "", "Change this announcement's visibility")
	PutAnnouncements.Flags().StringP("channelid", "c", "", "Change this announcement's discord channel id")
	PutAnnouncements.Flags().StringP("messageid", "m", "", "Change this announcement's discord message id")

}

func putAnnouncements(host string, port string, id string, payload *UpdateAnnouncement) {
	// ----- Check if Id was Given -----
	if id == "" {
		fmt.Println("Announcement id required for put! Please use the --id flag")
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
	request, err := http.Get(oldPayloadUrl.String())
	if err != nil {
		fmt.Printf("Error retrieveing %s: %s", payload.Uuid, err)
		return
	}
	defer request.Body.Close()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var oldPayload CreateAnnouncement
	json.Unmarshal(body, &oldPayload)

	// ----- Prompt User for New Values -----
	scanner := bufio.NewScanner(os.Stdin)

	// ----- Uuid -----
	// Known issue: Despite the response going through and output saying uuid has been updated
	// it does not actually update
	if payload.Uuid == "" {
		changeUuid, err := changePrompt("uuid", oldPayload.Uuid, scanner)
		if err != nil {
			fmt.Println("error with changing uuid:", err)
			return
		}
		if changeUuid != nil {
			payload.Uuid = string(changeUuid)
		} else {
			payload.Uuid = oldPayload.Uuid
		}
	}

	// ----- Visibility -----
	if payload.Visibility.String == "" {
		changeVisibility, err := changePrompt("visibility", oldPayload.Visibility, scanner)
		if err != nil {
			fmt.Println("error with changing visibility:", err)
			return
		}
		if changeVisibility != nil {
			payload.Visibility = dbtypes.StringtoNullString(string(changeVisibility))
		} else {
			payload.Visibility = dbtypes.StringtoNullString(oldPayload.Visibility)
		}
	}

	// ----- Announce At ------
	if payload.AnnounceAt.Int64 == 0 {
		oldAnnounceAt := strconv.FormatInt(oldPayload.AnnounceAt, 10)

		changeAnnounceAt, err := changePrompt("announce at", oldAnnounceAt, scanner)
		if err != nil {
			fmt.Println("error with changing announce at:", err)
			return
		}
		if changeAnnounceAt != nil {
			announceAtInt64, err := convert.ByteSlicetoInt64(changeAnnounceAt)
			if err != nil {
				fmt.Println(err)
				return
			}
			payload.AnnounceAt = dbtypes.Int64toNullInt64(announceAtInt64)
		} else {
			payload.AnnounceAt = dbtypes.Int64toNullInt64(oldPayload.AnnounceAt)
		}
	}

	// ----- Discord Channel ID -----
	if payload.DiscordChannelID.String == "" {
		changeDiscordChannelID, err := changePrompt("discord channel id", oldPayload.DiscordChannelID.String, scanner)
		if err != nil {
			fmt.Println("error with changing :", err)
			return
		}
		if changeDiscordChannelID != nil {
			payload.DiscordChannelID = dbtypes.StringtoNullString(string(changeDiscordChannelID))
		} else {
			payload.DiscordChannelID = oldPayload.DiscordChannelID
		}
	}

	// ----- Discord Message ID -----
	if payload.DiscordMessageID.String == "" {
		changeDiscordMessageID, err := changePrompt("discord message id", oldPayload.DiscordMessageID.String, scanner)
		if err != nil {
			fmt.Println("error with changing :", err)
			return
		}
		if changeDiscordMessageID != nil {
			payload.DiscordMessageID = dbtypes.StringtoNullString(string(changeDiscordMessageID))
		} else {
			payload.DiscordMessageID = oldPayload.DiscordMessageID
		}
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

// ============================================= Helper functions =============================================

// Returns a byte slice, if the byte slice is nil, no change will be made. Otherwise, a change will be made
func changePrompt(dataToBeChanged string, currentData string, scanner *bufio.Scanner) ([]byte, error) {
	fmt.Printf("Would you like to change this announcements's \x1b[1m%s\x1b[0m?[y/n]\nCurrent announcement's %s: \x1b[93m%s\x1b[0m\n", dataToBeChanged, dataToBeChanged, currentData)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %s", err)
	}
	userInput := scanner.Bytes()

	changeData, err := yesOrNo(userInput, scanner)
	if err != nil {
		return nil, err
	}
	if changeData {
		fmt.Printf("Please enter a new \x1b[1m%s\x1b[0m for the announcement:\n", dataToBeChanged)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return nil, fmt.Errorf("error reading new %s: %s", dataToBeChanged, err)
		}
		return scanner.Bytes(), nil
	} else {
		return nil, nil
	}
}

func yesOrNo(userInput []byte, scanner *bufio.Scanner) (bool, error) {
	userInputString := strings.ToUpper(string(userInput))

	if userInputString == "YES" || userInputString == "Y" || userInputString == "TRUE" {
		return true, nil
	} else if userInputString == "NO" || userInputString == "N" || userInputString == "FALSE" {
		return false, nil
	} else {
		fmt.Println("Invalid input, please try again.")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return false, fmt.Errorf("error scanning new input: %s", err)
		}
		return yesOrNo(scanner.Bytes(), scanner)
	}
}
