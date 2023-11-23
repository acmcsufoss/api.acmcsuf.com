package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/cli/interactors_generator"
)

// configFile is the path to the config file.
var configFile = flag.String("config", "config.json", "the path to the config file")

// outputFile is the path to the output file.
var outputFile = flag.String("output", "output.go", "the path to the output file")

func main() {
	// Parse the CLI flags.
	flag.Parse()

	// Read the config file.
	config := &interactors_generator.Config{}
	err := interactors_generator.ReadConfigFile(*configFile, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate the code.
	generatedCode := ""
	err = config.Render(&generatedCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write the code to the output file.
	err = os.WriteFile(*outputFile, []byte(generatedCode), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
