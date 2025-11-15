package utils

import (
	"bufio"
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Reoccuring functions for CLI files

// --------------------------- Printing to terminal ---------------------------

// Returns a byte slice, if nil, no changes shall be made. Else, if a byte slice were to return, change the payload value
func ChangePrompt(dataToBeChanged string, currentData string, scanner *bufio.Scanner, entity string) ([]byte, error) {
	fmt.Printf("Would you like to change this %s's \x1b[1m%s\x1b[0m?[y/n]\nCurrent event's %s: \x1b[93m%s\x1b[0m\n", entity, dataToBeChanged, dataToBeChanged, currentData)
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

// For unix times of int64 to readable format of 03:04:05PM 01/02/06
func FormatUnix(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format("01/02/06 03:04PM")
}

// Print the struct passed into it in a nice display in the terminal
func PrintStruct(s any) {
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	// Incase a pointer to a struct was passed
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}

	if val.Kind() != reflect.Struct {
		fmt.Println("error, not a struct")
		return
	}

	fmt.Printf("%s:\n", typ.Name())
	for i := 0; i < val.NumField(); i++ {
		v := val.Field(i)
		t := typ.Field(i)
		display := ""

		// If a value is still in interface, such as CreateEventParams.StartAt (as of writing this) unwrap it.
		if v.Kind() == reflect.Interface && !v.IsNil() {
			v = v.Elem()
		}
		if v.Kind() == reflect.Ptr && !v.IsNil() {
			v = v.Elem()
		}

		switch v.Type() {
		case reflect.TypeOf(sql.NullInt64{}):
			n := v.Interface().(sql.NullInt64)
			if n.Valid {
				display = FormatUnix(n.Int64)
			} else {
				display = "NULL"
			}

		case reflect.TypeOf(sql.NullString{}):
			n := v.Interface().(sql.NullString)
			if n.Valid {
				display = n.String
			} else {
				display = "NULL"
			}

		case reflect.TypeOf(sql.NullBool{}):
			n := v.Interface().(sql.NullBool)
			if n.Valid {
				display = strconv.FormatBool(n.Bool)
			} else {
				display = "NULL"
			}

		default:
			switch v.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				display = FormatUnix(v.Int())

			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				u := v.Uint()
				if u <= math.MaxInt64 {
					display = FormatUnix(int64(u))
				} else {
					display = fmt.Sprintf("%v", u)
				}

			default:
				if v.CanInterface() {
					display = fmt.Sprintf("%v", v.Interface())
				} else {
					display = "<unexported>"
				}
			}
		}

		fmt.Printf("\t%-20s | %s\n", t.Name, display)
	}

	fmt.Println("----------------------------------------------------------------")
}

// --------------------------- Getting values from input ---------------------------

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
	re := regexp.MustCompile(`(\d{2}):(\d{2})`)
	parsedDuration := re.FindStringSubmatch(duration)

	if parsedDuration == nil {
		return -1, fmt.Errorf("error, duration time must be in the format: 03:04")
	}

	durHour := parsedDuration[1]
	durMin := parsedDuration[2]

	// fmt.Println("Parsed times:", durHour, durMin, durSec)
	intDurHour, err := strconv.Atoi(durHour)
	if err != nil {
		return -1, fmt.Errorf("error converting hour to int: %s", err)
	}

	intDurMin, err := strconv.Atoi(durMin)
	if err != nil {
		return -1, fmt.Errorf("error converting minute to int: %s", err)
	}

	totalDuration := startUnix.Add(
		time.Duration(intDurHour)*time.Hour +
			time.Duration(intDurMin)*time.Minute,
	)

	return totalDuration.Unix(), nil
}
