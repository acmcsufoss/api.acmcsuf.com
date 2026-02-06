package utilities

import (
	"bufio"
	"fmt"
)

// https://go.dev/wiki/Iota
const (
	PromptDateTime = iota
	PromptTime
	PromptChange
)

// Constants for change prompt
const (
	Entity       = "Entity"
	DataToChange = "DataToChange"
	CurrentData  = "CurrentData"
)

type InputContext struct {
	Prompt string
	Extra  map[string]interface{}
}

// Reusable prompt function
// Alan Turing was similing from above when this function was written I think.
func InputPrompt(scan *bufio.Scanner, ctx InputContext, specialPrompt ...int) (interface{}, error) {

	if len(specialPrompt) != 0 {
		for _, elm := range specialPrompt {
			switch elm {
			case PromptDateTime:
				// TODO: handle date time
			case PromptTime:
				// TODO: handle time
			case PromptChange:
				for {
					fmt.Printf(
						"Would you like to change this %s's \x1b[1m%s\x1b[0m?\nCurrent %s's %s: \x1b[93m%s\x1b[0m\n",
						ctx.Extra[Entity],
						ctx.Extra[DataToChange],
						ctx.Extra[Entity],
						ctx.Extra[DataToChange],
						ctx.Extra[CurrentData],
					)

					confirmation, err := YesOrNo(scan)
					if err != nil {
						fmt.Println("Invalid input, please type yes or no")
						continue
					}

					if confirmation {
						fmt.Printf(
							"Please enter a new \x1b[1m%s\x1b[0m for the %s:\n",
							ctx.Extra[DataToChange],
							ctx.Extra[Entity],
						)
						scan.Scan()
						if err := scan.Err(); err != nil {
							fmt.Println("Error reading input:", err)
							continue
						}

						// TODO: adjust for if user wants to change time
						newData := getIn(scan)

						return newData, nil
					} else {
						fmt.Println("No channges made.")
					}
				}
			}
		}

		// In our default case we will just ask the prompt and take a byte slice of inputs
	} else {
		fmt.Println(ctx.Prompt)
		scan.Scan()
		if err := scan.Err(); err != nil {
			return "", err
		}

		input := getIn(scan)

		return input, nil
	}

}

// scanner is a pointer to a buffer, therefore to avoid possible
// unexpected values we can copy its current pointer into a byte slice
func getIn(scan *bufio.Scanner) []byte {
	raw := scan.Bytes()
	input := make([]byte, len(raw))
	copy(input, raw)

	return input
}
