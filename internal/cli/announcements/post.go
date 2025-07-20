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

	"github.com/acmcsufoss/api.acmcsuf.com/utils/convert"
	"github.com/acmcsufoss/api.acmcsuf.com/utils/dbtypes"
	"github.com/spf13/cobra"
)

var PostAnnouncement = &cobra.Command{
	Use:   "post",
	Short: "post a new announcement",

	Run: func(cmd *cobra.Command, args []string) {
		payload := CreateAnnouncement{}

		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		payload.Uuid, _ = cmd.Flags().GetString("uuid")
		payload.Visibility, _ = cmd.Flags().GetString("visibility")
		payload.AnnounceAt, _ = cmd.Flags().GetInt64("announceat")
		channelIdString, _ := cmd.Flags().GetString("channelid")
		messageIdString, _ := cmd.Flags().GetString("messageid")

		payload.DiscordChannelID = dbtypes.StringtoNullString(channelIdString)
		payload.DiscordMessageID = dbtypes.StringtoNullString(messageIdString)

		postAnnouncement(host, port, &payload)
	},
}

func init() {
	// URL flags
	PostAnnouncement.Flags().String("host", "127.0.0.1", "Set a custom host (ex: 127.0.0.1)")
	PostAnnouncement.Flags().String("port", "8080", "Set a custom port (ex: 8080)")
	PostAnnouncement.Flags().String("id", "", "PUT to announcement by it's id")

	// Payload flags
	PostAnnouncement.Flags().StringP("visibility", "v", "", "Set this announcement's visibility")
	PostAnnouncement.Flags().StringP("announceat", "a", "", "Set this annoucement's announce at (Note: in unix time)")
	PostAnnouncement.Flags().StringP("channelid", "c", "", "Set this annoucement's channel id")
	PostAnnouncement.Flags().StringP("messageid", "m", "", "Set this annoucement's message id")
}

func postAnnouncement(host string, port string, payload *CreateAnnouncement) {

	scanner := bufio.NewScanner(os.Stdin)

	// ----- Uuid -----
	if payload.Uuid == "" {
		fmt.Println("Please enter the announcement's uuid:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error with reading uuid:", err)
			return
		}

		uuidBuffer := scanner.Bytes()
		payload.Uuid = string(uuidBuffer)
	}

	// ----- Visibility -----
	if payload.Visibility == "" {
		fmt.Println("Please enter this annoucement's visibility:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error with reading visiblity:", err)
			return
		}

		visibilityBuffer := scanner.Bytes()
		payload.Visibility = string(visibilityBuffer)
	}

	// ----- Announce at -----
	if payload.AnnounceAt == 0 {
		fmt.Println("Please enter the \"announce at\" of the announcement in the following format:\n [Month]/[Day] [Hour]:[Minute]:[Second][PM | AM] '[Last 2 digits of year] -0700")
		fmt.Println("For example: \x1b[93m01/02 03:04:05PM '06 -0700\x1b[0m")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading anounce at:", err)
			return
		}

		announceatBuffer := scanner.Bytes()

		// Sorry
		var err error
		payload.AnnounceAt, err = convert.ByteSlicetoUnix(announceatBuffer)
		if err != nil {
			fmt.Println("error converting byte slice to unix time (of type int64):", err)
			return
		}
	}

	// ----- Discord Channel Id -----
	if !payload.DiscordChannelID.Valid {
		fmt.Println("Please enter this annoucement's discord channel id:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading  discord channel id:", err)
			return
		}

		channelIdBuffer := scanner.Bytes()
		payload.DiscordChannelID = dbtypes.StringtoNullString(string(channelIdBuffer))
	}

	// ----- Discord Message Id -----
	if !payload.DiscordMessageID.Valid {
		fmt.Println("Please enter this annoucement's message id:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading message id:", err)
			return
		}
		messageIdBuffer := scanner.Bytes()
		payload.DiscordMessageID = dbtypes.StringtoNullString(string(messageIdBuffer))
	}

	// ----- Confirmation -----
	fmt.Println("Is your event data correct? If not, type n or no. \n", *payload)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	confirmationBuffer := scanner.Bytes()
	confirmationString := strings.ToUpper(string(confirmationBuffer))

	if confirmationString == "NO" || confirmationString == "N" {
		return
	}

	// ----- Marshalling to Json -----
	jsonPayload, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println("error formating payload to json:", err)
		return
	}

	// ----- Constructing the Url -----
	host = fmt.Sprint(host, ":", port)
	path := "announcements"

	postURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	fmt.Println(postURL.String())
	// ----- Post -----
	response, err := http.Post(postURL.String(), "application/json", strings.NewReader(string(jsonPayload)))
	if err != nil {
		fmt.Println("error with post:", err)
		return
	}

	if response == nil {
		fmt.Println("no response recieved")
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
