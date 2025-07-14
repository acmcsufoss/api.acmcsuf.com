package convert

import (
	"fmt"
	"strconv"
	"time"
)

func ByteSlicetoInt64(data []byte) (int64, error) {
	// Padding, if number is too short we run into the eof and get an error (see binary.Read)
	/*if len(data) < 8 {
		padded := make([]byte, 8-len(data), 8)
		data = append(padded, data...)
	}
	*/
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
