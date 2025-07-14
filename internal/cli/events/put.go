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
	"github.com/acmcsufoss/api.acmcsuf.com/utils/dbtypes"
	_ "github.com/acmcsufoss/api.acmcsuf.com/utils/dbtypes"
	"github.com/spf13/cobra"
)

var PutEvents = &cobra.Command{
	Use:   "put",
	Short: "Used to update an event",

	Run: func(cmd *cobra.Command, args []string) {

		payload := CreateEvent{}

		// CLI for url
		host, _ := cmd.Flags().GetString("urlhost")
		port, _ := cmd.Flags().GetString("port")
		id, _ := cmd.Flags().GetString("id")

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
	PutEvents.Flags().StringP("uuid", "u", "", "Set uuid of new event")
	PutEvents.Flags().StringP("location", "l", "", "Set location of new event")
	PutEvents.Flags().Int64P("startat", "s", 0, "Set the start time of new event (Note: flag takes Unix time)")
	PutEvents.Flags().Int64P("endat", "e", 0, "Set the end time of new event (Note: flag takes unix time)")
	PutEvents.Flags().StringP("host", "H", "", "Set host of new event")
	PutEvents.Flags().BoolP("allday", "a", false, "Set if new event is all day")

}

func updateEvent(id string, host string, port string, payload *CreateEvent) {
	// ----- Check for Event Id -----
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

	var oldpayload CreateEvent
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

		if changeTheEventUuid != nil {
			payload.Uuid = string(changeTheEventUuid)
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

		if changeTheEventLocation != nil {
			payload.Location = string(changeTheEventLocation)
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

		if changeTheEventStartAt != nil {
			payload.StartAt, err = convert.ByteSlicetoInt64(changeTheEventStartAt)
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

		if changeTheEventEndAt != nil {
			payload.EndAt, err = convert.ByteSlicetoInt64(changeTheEventEndAt)
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

		if changeTheEventAllDay != nil {
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

		if changeTheEventHost != nil {
			payload.Host = string(changeTheEventHost)
		} else {
			payload.Host = oldpayload.Host
		}
	}

	// ----- PUT the payload -----

	updatePayload := UpdateEvent{
		Uuid:     payload.Uuid,
		Location: dbtypes.StringtoNullString(payload.Location),
		StartAt:  dbtypes.Int64toNullInt64(payload.StartAt),
		EndAt:    dbtypes.Int64toNullInt64(payload.EndAt),
		IsAllDay: dbtypes.BooltoNullBool(payload.IsAllDay),
		Host:     dbtypes.StringtoNullString(payload.Host),
	}

	// ----- Put the Payload -----
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

	// ----- Read Response Info -----
	fmt.Println("Response status:", putResponse.Status)

	body, err = io.ReadAll(putResponse.Body)
	if err != nil {
		fmt.Println("Error with body:", err)
		return
	}

	fmt.Println(string(body))
}

// ============================================= Helper functions =============================================

// Returns a byte slice, if nil, no changes shall be made. Else, if a byte slice were to return, change the payload value
func changePrompt(dataToBeChanged string, currentData string, scanner *bufio.Scanner) ([]byte, error) {
	fmt.Printf("Would you like to change this event's \x1b[1m%s\x1b[0m?[y/n]\nCurrent event's %s: \x1b[93m%s\x1b[0m\n", dataToBeChanged, dataToBeChanged, currentData)
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
		fmt.Printf("Please enter a new \x1b[1m%s\x1b[0m for the event:\n", dataToBeChanged)
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
