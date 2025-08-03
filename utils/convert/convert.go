package convert

import (
	"fmt"
	"strconv"
	"time"
)

// Common Byte Slice Conversions

func ByteSlicetoInt64(data []byte) (int64, error) {
	fmt.Printf("%d, %v", data, data)
	number, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, fmt.Errorf("error converting bytes to int: %s", err)
	}

	return int64(number), nil
}

func ByteSlicetoUnix(data []byte) (int64, error) {
	// 12 hour format. Note: AM AND PM HAVE TO BE CAPITALIZED
	layout := "03:04:05PM 01/02/06"
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return -1, fmt.Errorf("error in getting location for time: %s", err)
	}

	timeString := string(data)

	// Parse time in relation to los angeles time, or more familiarly, PST time
	startTime, err := time.ParseInLocation(layout, timeString, loc)
	if err != nil {
		return -1, fmt.Errorf("error parsing time format: %s", err)

	}

	return startTime.Unix(), nil
}
