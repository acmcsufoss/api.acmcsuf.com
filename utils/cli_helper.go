package utils

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// For unix times of int64 to readable format of 03:04:05PM 01/02/06
func FormatUnix(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format("01/02/06 03:04PM")
}

func CheckConnection(url string) error {
	_, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("\x1b[1;37;41mUNABLE TO CONNECT\x1b[0m | %s\n\t↳ %v",
			"Did you forget to start the server?",
			err)
	}
	return nil
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
