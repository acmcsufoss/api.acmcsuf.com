package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const timeFormat = "01/02/06 03:04PM"

// For unix times of int64 to readable format of 03:04:05PM 01/02/06
func FormatUnix(unixTime int64) string {
	t := time.Unix(unixTime, 0)
	return t.Format(timeFormat)
}

// Parses readable time format to Unix time integer
func ParseTime(timeStr string) (int64, error) {
	t, err := time.Parse(timeFormat, timeStr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse time: %v", err)
	}
	return t.Unix(), nil
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

func UnixToTime(v int64) time.Time {
	return time.Unix(v, 0)
}

func UnixToTimePtr(v *int64) *time.Time {
	if v == nil {
		return nil
	}
	t := time.Unix(*v, 0)
	return &t
}
