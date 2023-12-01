package tests

import (
	"encoding/json"
	"fmt"
	"os"
)

type TestData map[string]interface{}

func GetTestData(filename string) TestData {
	fileBytes, _ := os.ReadFile(filename)
	var data TestData

	err := json.Unmarshal(fileBytes, &data)

	if err != nil {
		return nil
	}

	return data
}

func TestResources() {
	data := GetTestData("test_data/data.json")
	fmt.Println(data)
}
