package events

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
	_ "github.com/acmcsufoss/api.acmcsuf.com/utils/dbtypes"
	"github.com/spf13/cobra"
)

var PutEvents = &cobra.Command{
	Use:   "put",
	Short: "Used to update an event",

	Run: func(cmd *cobra.Command, args []string) {

		id, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Println("Event ID is required!")
			return
		}

		payload := Event{}

		// CLI for url
		host, _ := cmd.Flags().GetString("urlhost")
		port, _ := cmd.Flags().GetString("port")

		// CLI for payload
		payload.Uuid, _ = cmd.Flags().GetString("uuid")
		payload.Location, _ = cmd.Flags().GetString("location")
		payload.StartAt, _ = cmd.Flags().GetInt64("startat")
		payload.EndAt, _ = cmd.Flags().GetInt64("endat")
		payload.IsAllDay, _ = cmd.Flags().GetBool("allday")
		payload.Host, _ = cmd.Flags().GetString("host")
		updateEvent(id, host, port, &payload)
	},
}

func init() {
	// URL Flags
	PutEvents.Flags().String("id", "", "Event to update")
	PutEvents.Flags().String("urlhost", "127.0.0.1", "Custom host (ex: 127.0.0.1)")
	PutEvents.Flags().String("port", "8080", "Custom port (ex: 8080)")

	// Payload flags
	PutEvents.Flags().String("uuid", "", "Change uuid of event")
	PutEvents.Flags().String("location", "", "Change location of event")
	PutEvents.Flags().Int("startat", 0, "Change the start time of event")
	PutEvents.Flags().Int("endat", 0, "Change the end time of event")
	PutEvents.Flags().String("host", "", "Change host of event")
	PutEvents.Flags().Bool("allday", false, "Change if event is all day")
}

func updateEvent(id string, host string, port string, payload *Event) {
	if id == "" {
		fmt.Println("Event ID is required!")
		return
	}

	// ----- Constructing Url -----
	host = fmt.Sprint(host, ":", port)

	path := fmt.Sprint("events", "/", id)

	retrievalURL := &url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	// ----- Retrieve payload -----
	getResponse, err := http.Get(retrievalURL.String())
	if err != nil {
		fmt.Printf("Error retrieveing %s: %s", id, err)
		return
	}
	defer getResponse.Body.Close()

	body, err := io.ReadAll(getResponse.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var oldpayload Event
	json.Unmarshal(body, &oldpayload)

	scanner := bufio.NewScanner(os.Stdin)

	// ----- Change the event -----
	// Note: We want to PUT the payload, not old payload
	// payload values are empty if user did not input a value in the command line

	// ----- uuid -----
	if payload.Uuid == "" {
		changeTheEventUuid, err := changePrompt("uuid", oldpayload.Uuid, scanner)
		if err != nil {
			fmt.Println(err) // Custom errors in changePrompt()
			return
		}

		if changeTheEventUuid {
			newUuidBuffer := scanner.Bytes()
			payload.Uuid = string(newUuidBuffer)
		} else {
			payload.Uuid = oldpayload.Uuid
		}

	}

	// ----- Location -----
	if payload.Location == "" {
		changeTheEventLocation, err := changePrompt("location", oldpayload.Location, scanner)
		if err != nil {
			fmt.Println(err)
			return
		}

		if changeTheEventLocation {
			newLocationBuffer := scanner.Bytes()
			payload.Location = string(newLocationBuffer)
		} else {
			payload.Location = oldpayload.Location
		}
	}

	// ----- Start time -----
	if payload.StartAt == 0 {
		changeTheEventStartAt, err := changePrompt("start time", strconv.FormatInt(oldpayload.StartAt, 10), scanner)
		if err != nil {
			fmt.Println(err)
			return
		}

		if changeTheEventStartAt {
			newStartAtBuffer := scanner.Bytes()
			payload.StartAt, err = convert.ByteSlicetoInt64(newStartAtBuffer)
			if err != nil {
				fmt.Println("Error with reading start integer:", err)
				return
			}
		} else {
			payload.StartAt = oldpayload.StartAt
		}
	}

	// ----- End time -----
	if payload.EndAt == 0 {
		changeTheEventEndAt, err := changePrompt("end time", strconv.FormatInt(oldpayload.EndAt, 10), scanner)
		if err != nil {
			fmt.Println(err)
			return
		}

		if changeTheEventEndAt {
			newEndAtBuffer := scanner.Bytes()
			payload.EndAt, err = convert.ByteSlicetoInt64(newEndAtBuffer)
			if err != nil {
				fmt.Println("Error with reading end integer:", err)
				return
			}
		} else {
			payload.EndAt = oldpayload.EndAt
		}
	}

	// ----- All day -----
	// This is kind of awkward but I don't know have a workaround at the moment
	if !payload.IsAllDay {
		changeTheEventAllDay, err := changePrompt("all day status", strconv.FormatBool(oldpayload.IsAllDay), scanner)
		if err != nil {
			fmt.Println(err)
			return
		}

		if changeTheEventAllDay {
			newAllDayBuffer := scanner.Bytes()
			payload.IsAllDay, err = yesOrNo(newAllDayBuffer, scanner)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			payload.IsAllDay = oldpayload.IsAllDay
		}
	}

	// ----- Host -----
	if payload.Host == "" {
		changeTheEventHost, err := changePrompt("host", oldpayload.Host, scanner)
		if err != nil {
			fmt.Println(err)
			return
		}

		if changeTheEventHost {
			newHostBuffer := scanner.Bytes()
			payload.Host = string(newHostBuffer)
		} else {
			payload.Host = oldpayload.Host
		}
	}

	// ----- PUT the payload -----

	/*updatePayload := PutEvent{
		Uuid:     payload.Uuid,
		Location: toNullString(payload.Location),
		StartAt:  toNullInt64(payload.StartAt),
		EndAt:    toNullInt64(payload.EndAt),
		IsAllDay: sql.NullBool{
			Bool:  payload.IsAllDay,
			Valid: true,
		},
		Host: toNullString(payload.Host),
	}*/
	updatePayload := PutEvent{
		Uuid:     &payload.Uuid,
		Location: &payload.Location,
		StartAt:  &payload.StartAt,
		EndAt:    &payload.EndAt,
		IsAllDay: &payload.IsAllDay,
		Host:     &payload.Host,
	}

	newPayload, err := json.Marshal(updatePayload)
	if err != nil {
		fmt.Println("Error marshaling data:", err)
		return
	}

	client := &http.Client{}

	request, err := http.NewRequest(http.MethodPut, retrievalURL.String(), bytes.NewBuffer(newPayload))
	if err != nil {
		fmt.Println("Problem with PUT:", err)
		return
	}

	putResponse, err := client.Do(request)
	if err != nil {
		fmt.Println("Error with response:", err)
		return
	}
	defer putResponse.Body.Close()

	body, err = io.ReadAll(putResponse.Body)
	if err != nil {
		fmt.Println("Error with body:", err)
		return
	}

	fmt.Println(string(body))
}

// ============================================= Helper functions =============================================

// The following function does these things:
// 1. Ask the user if they want to chnage [x] data
// 2. User inputs a response to scanner
// 3. If yesOrNo() is true then:
// 4. Prompt user for new data [x]
// 5. Scanner scans the users response
// 6. Return true. Scanner holds the byte slice which can be used by the function that called this one
// Note: I think this is a little risky, I believe it would be very easy to override the bytes slice by accident.
//
//	If the person reading this thinks there is a better way to manage scanner, then by all means go ahead and try.

// Returns a bool, indicating if the user wants to change the current data field passed into the function
func changePrompt(dataToBeChanged string, currentData string, scanner *bufio.Scanner) (bool, error) {
	fmt.Printf("Would you like to change this event's \x1b[1m%s\x1b[0m?[y/n]\nCurrent event's %s: \x1b[93m%s\x1b[0m\n", dataToBeChanged, dataToBeChanged, currentData)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("error reading input: %s", err)
	}
	userInput := scanner.Bytes()

	changeData, err := yesOrNo(userInput, scanner)
	if err != nil {
		return false, err
	}
	if changeData {
		fmt.Printf("Please enter a new \x1b[1m%s\x1b[0m for the event:\n", dataToBeChanged)
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return false, fmt.Errorf("error reading new %s: %s", dataToBeChanged, err)
		}
		return true, nil
	} else {
		return false, nil
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
