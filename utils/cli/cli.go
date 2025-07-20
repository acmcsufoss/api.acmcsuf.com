package cli

import (
	"bufio"
	"fmt"
	"strings"
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
