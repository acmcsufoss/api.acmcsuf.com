package events

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/acmcsufoss/api.acmcsuf.com/utils/convert"
	"github.com/spf13/cobra"
)

var PostEvent = &cobra.Command{
	Use:   "post",
	Short: "Post a new event.",

	Run: func(cmd *cobra.Command, args []string) {
		payload := CreateEvent{}

		urlhost, _ := cmd.Flags().GetString("urlhost")
		port, _ := cmd.Flags().GetString("port")

		payload.Uuid, _ = cmd.Flags().GetString("uuid")
		payload.Location, _ = cmd.Flags().GetString("location")
		payload.StartAt, _ = cmd.Flags().GetInt64("startat")
		payload.EndAt, _ = cmd.Flags().GetInt64("endat")
		payload.IsAllDay, _ = cmd.Flags().GetBool("isallday")
		payload.Host, _ = cmd.Flags().GetString("host")

		postEvent(urlhost, port, &payload)
	},
}

func init() {

	// URL Flags
	PostEvent.Flags().String("urlhost", "127.0.0.1", "Custom host (ex: 127.0.0.1)")
	PostEvent.Flags().String("port", "8080", "Custom port (ex: 8080)")

	// Payload flags
	PostEvent.Flags().StringP("uuid", "u", "", "Set uuid of new event")
	PostEvent.Flags().StringP("location", "l", "", "Set location of new event")
	PostEvent.Flags().Int64P("startat", "s", 0, "Set the start time of new event (Note: flag takes Unix time)")
	PostEvent.Flags().Int64P("endat", "e", 0, "Set the end time of new event (Note: flag takes unix time)")
	PostEvent.Flags().StringP("host", "H", "", "Set host of new event")
	PostEvent.Flags().BoolP("allday", "a", false, "Set if new event is all day")

}
func postEvent(urlhost string, port string, payload *CreateEvent) {

	scanner := bufio.NewScanner(os.Stdin)

	// ----- Uuid -----
	if payload.Uuid == "" {
		fmt.Println("Please enter event's uuid:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}

		uuidBuffer := scanner.Bytes()
		payload.Uuid = string(uuidBuffer)
	}

	// ----- Location -----
	if payload.Location == "" {
		fmt.Println("please enter the event's location:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}

		locationBuffer := scanner.Bytes()
		payload.Location = string(locationBuffer)
	}

	// ----- Start Time -----
	if payload.StartAt == 0 {
		fmt.Println("Please enter the start time of the event in the following format:\n [Month]/[Day] [Hour]:[Minute]:[Second][PM | AM] '[Last 2 digits of year] -0700")
		fmt.Println("For example: \x1b[93m01/02 03:04:05PM '06 -0700\x1b[0m")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading start time:", err)
			return
		}
		startTimeBuffer := scanner.Bytes()
		startTime, err := convert.ByteSlicetoUnix(startTimeBuffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		payload.StartAt = startTime
	}

	// ----- End Time -----
	if payload.EndAt == 0 {
		fmt.Println("Please enter the end time of the event in the following format:\n [Month]/[Day] [Hour]:[Minute]:[Second][PM | AM] '[Last 2 digits of year] -0700")
		fmt.Println("For example: \x1b[93m01/02 03:04:05PM '06 -0700\x1b[0m")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading end time:", err)
			return
		}
		endTimeBuffer := scanner.Bytes()
		endTime, err := convert.ByteSlicetoUnix(endTimeBuffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		payload.EndAt = endTime
	}

	// ----- Is all day -----

	// This is kind of awkward
	if !payload.IsAllDay {
		fmt.Println("Is the event all day?")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}

		isAllDayBuffer := scanner.Bytes()
		isAllDayString := strings.ToUpper(string(isAllDayBuffer))

		switch isAllDayString {
		case "YES", "Y":
			payload.IsAllDay = true
		case "NO", "N":
			payload.IsAllDay = false
		default:
			fmt.Println("Invalid input.")
			return
		}
	}

	// ----- Host -----
	if payload.Host == "" {
		fmt.Println("Please enter the event host:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return
		}

		hostBuffer := scanner.Bytes()
		payload.Host = string(hostBuffer)
	}

	// ----- Confirmation -----
	fmt.Println("Is your event data correct? If not, type n or no. [Note that time is displayed in UNIX time.]\n", payload)
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

	// ----- Convert to Json -----
	jsonEvent, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ----- Construct Url -----
	urlhost = fmt.Sprint(urlhost, ":", port)
	path := "events"

	postUrl := &url.URL{
		Scheme: "http",
		Host:   urlhost,
		Path:   path,
	}

	// ----- Post -----
	response, err := http.Post(postUrl.String(), "application/json", strings.NewReader(string(jsonEvent)))
	if err != nil {
		fmt.Println("Failed to post event:", err)
		return
	}

	if response == nil {
		fmt.Println("no response received")
		return
	}

	defer response.Body.Close()

	// ----- Read Response Info -----
	fmt.Println("Response Status:", response.Status)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	fmt.Println("Response body:", string(body))

}
