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
	timeString := string(data)

	startTime, err := time.Parse(time.Layout, timeString)
	if err != nil {
		return -1, err
	}

	return startTime.Unix(), nil
}
