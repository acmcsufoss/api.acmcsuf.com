package main

import (
	"flag"
	"fmt"

	"github.com/acmcsufoss/api.acmcsuf.com/cli/interactors_generator"
)

// configFile is the path to the config file.
var configFile = flag.String("config", "config.json", "the path to the config file")

func main() {
	// Parse the CLI flags.
	flag.Parse()

	// Read the config file.
	var config *interactors_generator.Config
	err := interactors_generator.ReadConfigFile(*configFile, config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate the code.
	var generatedCode *string
	err = config.Render(generatedCode)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Write the generated code to stdout.
	fmt.Println(generatedCode)
}
