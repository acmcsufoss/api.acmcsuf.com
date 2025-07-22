package cli

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Reoccuring functions for CLI files

// Returns a byte slice, if nil, no changes shall be made. Else, if a byte slice were to return, change the payload value
func ChangePrompt(dataToBeChanged string, currentData string, scanner *bufio.Scanner) ([]byte, error) {
	fmt.Printf("Would you like to change this event's \x1b[1m%s\x1b[0m?[y/n]\nCurrent event's %s: \x1b[93m%s\x1b[0m\n", dataToBeChanged, dataToBeChanged, currentData)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %s", err)
	}
	userInput := scanner.Bytes()

	changeData, err := YesOrNo(userInput, scanner)
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

// YesOrNo checks the user input for a yes or no response.
func YesOrNo(userInput []byte, scanner *bufio.Scanner) (bool, error) {
	userInputString := strings.ToUpper(string(userInput))

	switch userInputString {
	case "YES", "Y", "TRUE":
		return true, nil
	case "NO", "N", "FALSE":
		return false, nil
	default:
		fmt.Println("Invalid input, please try again.")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return false, fmt.Errorf("error scanning new input: %s", err)
		}
		return YesOrNo(scanner.Bytes(), scanner)
	}
}

func TimeAfterDuration(startTime int64, duration string) (int64, error) {
	startUnix := time.Unix(startTime, 0)

	// Since I cant go from 03:04:05 -> time.Time directly, I am left doing some... parsing
	re := regexp.MustCompile(`(\d{2}):(\d{2}):(\d{2})`)
	parsedDuration := re.FindStringSubmatch(duration)

	if parsedDuration == nil {
		return -1, fmt.Errorf("error, duration time must be in the format: 03:04:05")
	}

	durHour := parsedDuration[1]
	durMin := parsedDuration[2]
	durSec := parsedDuration[3]

	//fmt.Println("Parsed times:", durHour, durMin, durSec)
	intDurHour, err := strconv.Atoi(durHour)
	if err != nil {
		return -1, fmt.Errorf("error converting hour to int: %s", err)
	}

	intDurMin, err := strconv.Atoi(durMin)
	if err != nil {
		return -1, fmt.Errorf("error converting minute to int: %s", err)
	}

	intDurSec, err := strconv.Atoi(durSec)
	if err != nil {
		return -1, fmt.Errorf("error converting second to int: %s", err)
	}

	totalDuration := startUnix.Add(
		time.Duration(intDurHour)*time.Hour +
			time.Duration(intDurMin)*time.Minute +
			time.Duration(intDurSec)*time.Second,
	)

	return totalDuration.Unix(), nil
}

// For unix times of int64 to readable format of 03:04:05PM 01/02/06
func FormatUnix(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format("03:04:05PM 01/02/06")
}
