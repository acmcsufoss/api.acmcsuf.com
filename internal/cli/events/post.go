package events

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var PostEvent = &cobra.Command{
	Use:   "post",
	Short: "Post a new event.",

	Run: promptEvent,
}

func promptEvent(cmd *cobra.Command, args []string) {
	newEvent := Event{}

	scanner := bufio.NewScanner(os.Stdin)

	// ----- Uuid -----
	fmt.Println("Please enter event's uuid:")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	uuidBuffer := scanner.Bytes()
	newEvent.Uuid = string(uuidBuffer)

	// ----- Location -----
	fmt.Println("please enter the event's location:")
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	locationBuffer := scanner.Bytes()
	newEvent.Location = string(locationBuffer)

	// ----- Start Time -----
	fmt.Println("Please enter the start time of the event in the following format:\n [Month]/[Day] [Hour]:[Minute]:[Second][PM | AM] '[Last 2 digits of year] -0700")
	fmt.Println("For example: \x1b[93m01/02 03:04:05PM '06 -0700\x1b[0m")
	startTime, err := inputTime(scanner)
	if err != nil {
		fmt.Println(err)
		return
	}

	newEvent.StartAt = startTime

	// ----- End Time -----
	fmt.Println("Please enter the end time of the event in the same format as above:")
	endTime, err := inputTime(scanner)
	if err != nil {
		fmt.Println(err)
		return
	}

	newEvent.EndAt = endTime

	// ----- Is all day -----
	fmt.Println("Is the event all day?")
	scanner.Scan()
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	isAllDayBuffer := scanner.Bytes()
	isAllDayString := strings.ToUpper(string(isAllDayBuffer))

	if isAllDayString == "YES" || isAllDayString == "Y" {
		newEvent.IsAllDay = true
	} else if isAllDayString == "NO" || isAllDayString == "N" {
		newEvent.IsAllDay = false
	} else {
		fmt.Println("Invalid input.")
		return
	}

	// ----- Host -----
	fmt.Println("Please enter the event host:")
	scanner.Scan()
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	hostBuffer := scanner.Bytes()
	newEvent.Host = string(hostBuffer)

	// ----- Confirmation -----
	fmt.Println("Is your event data correct? If not, type n or no. [Note that time is displayed in UNIX time.]\n", newEvent)
	scanner.Scan()
	if err = scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	confirmationBuffer := scanner.Bytes()
	confirmationString := strings.ToUpper(string(confirmationBuffer))

	if confirmationString == "NO" || confirmationString == "N" {
		return
	}

	// ----- Convert to Json -----
	jsonEvent, err := json.Marshal(newEvent)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ----- Post -----
	resp, err := http.Post("http://127.0.0.1:8080/events", "application/json", strings.NewReader(string(jsonEvent)))
	if err != nil {
		fmt.Println("Failed to post event:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}

	fmt.Println("Response body:", string(body))

}

func inputTime(scanner *bufio.Scanner) (int64, error) {
	scanner.Scan()
	timeBuffer := scanner.Bytes()
	timeString := string(timeBuffer)

	startTime, err := time.Parse(time.Layout, timeString)
	if err != nil {
		return -1, err
	}

	return startTime.Unix(), nil
}
