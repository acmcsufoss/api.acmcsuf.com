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

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
	"github.com/acmcsufoss/api.acmcsuf.com/utils"

	"github.com/spf13/cobra"
)

var PostEvent = &cobra.Command{
	Use:   "post",
	Short: "Post a new event.",

	Run: func(cmd *cobra.Command, args []string) {
		payload := models.CreateEventParams{}

		urlhost, _ := cmd.Flags().GetString("urlhost")
		port, _ := cmd.Flags().GetString("port")

		payload.Uuid, _ = cmd.Flags().GetString("uuid")
		payload.Location, _ = cmd.Flags().GetString("location")
		startAtString, _ := cmd.Flags().GetString("startat")
		duration, _ := cmd.Flags().GetString("duration")
		payload.IsAllDay, _ = cmd.Flags().GetBool("isallday")
		payload.Host, _ = cmd.Flags().GetString("host")

		if startAtString != "" {
			var err error
			payload.StartAt, err = utils.ByteSlicetoUnix([]byte(startAtString))
			if err != nil {
				fmt.Println(err)
				return
			}
			if duration != "" {
				var err error
				payload.EndAt, err = utils.TimeAfterDuration(payload.StartAt, duration)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}

		if duration != "" && startAtString == "" {
			fmt.Printf("--startat is required in order to use --duration")
		}

		changedFlags := eventFlags{
			uuid:     cmd.Flags().Lookup("uuid").Changed,
			location: cmd.Flags().Lookup("location").Changed,
			startat:  cmd.Flags().Lookup("startat").Changed,
			duration: cmd.Flags().Lookup("duration").Changed,
			isallday: cmd.Flags().Lookup("isallday").Changed,
			host:     cmd.Flags().Lookup("host").Changed,
		}

		postEvent(urlhost, port, &payload, changedFlags)
	},
}

func init() {
	// URL Flags
	PostEvent.Flags().String("urlhost", "127.0.0.1", "Custom host")
	PostEvent.Flags().String("port", "8080", "Custom port")

	// Payload flags
	PostEvent.Flags().StringP("uuid", "u", "", "Set uuid of new event")
	PostEvent.Flags().StringP("location", "l", "", "Set location of new event")
	PostEvent.Flags().StringP("startat", "s", "", "Set the start time of new event (Format: 03:04:05PM 01/02/06)")
	PostEvent.Flags().StringP("duration", "d", "", "Set the duration of new event (Format: 03:04:05)")

	PostEvent.Flags().StringP("host", "H", "", "Set host of new event")
	PostEvent.Flags().BoolP("isallday", "a", false, "Set if new event is all day")
}

func postEvent(urlhost string, port string, payload *models.CreateEventParams, changedFlag eventFlags) {

	err := utils.CheckConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	// ----- Uuid -----
	for {
		if changedFlag.uuid {
			break
		}

		fmt.Println("Please enter event's uuid:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		uuidBuffer := scanner.Bytes()
		payload.Uuid = string(uuidBuffer)
		break
	}

	// ----- Location -----
	for {
		if changedFlag.location {
			break
		}

		fmt.Println("please enter the event's location:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		locationBuffer := scanner.Bytes()
		payload.Location = string(locationBuffer)
		break
	}

	// ----- Start Time -----
	for {

		if changedFlag.startat {
			break
		}

		fmt.Println("Please enter the start time of the event in the following format:\n [Month]/[Day]/[Year] [Hour]:[Minute][PM | AM]")
		fmt.Println("For example: \x1b[93m01/02/06 03:04PM\x1b[0m")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading start time:", err)
			continue
		}
		startTimeBuffer := scanner.Bytes()
		startTime, err := utils.ByteSlicetoUnix(startTimeBuffer)
		if err != nil {
			fmt.Println(err)
			continue
		}

		payload.StartAt = startTime
		break
	}

	// ----- End Time (Duration) -----
	for {

		if changedFlag.duration {
			break
		}

		fmt.Println("Please enter the duration of the event in the following format:\n [Hour]:[Minute]")
		fmt.Println("For example: \x1b[93m03:04\x1b[0m")

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println("error reading end time:", err)
			continue
		}

		endTimeBuffer := scanner.Bytes()
		endTime, err := utils.TimeAfterDuration(payload.StartAt, string(endTimeBuffer))
		if err != nil {
			fmt.Println(err)
			continue
		}

		payload.EndAt = endTime
		break
	}

	// ----- Is all day -----

	for {
		if changedFlag.isallday {
			break
		}

		fmt.Println("Is the event all day?")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		isAllDayBuffer := scanner.Bytes()

		isAllDay, err := utils.YesOrNo(isAllDayBuffer, scanner)
		if err != nil {
			fmt.Println(err)
		}
		payload.IsAllDay = isAllDay
		break
	}

	// ----- Host -----
	for {
		if changedFlag.host {
			break
		}

		fmt.Println("Please enter the event host:")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Println(err)
			continue
		}

		hostBuffer := scanner.Bytes()
		payload.Host = string(hostBuffer)
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

	// ----- Convert to Json -----
	jsonEvent, err := json.Marshal(*payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ----- Construct Url -----
	urlhost = fmt.Sprint(urlhost, ":", port)
	path := "v1/events"

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
