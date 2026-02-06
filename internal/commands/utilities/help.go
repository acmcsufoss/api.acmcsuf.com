package utilities

import (
	"bufio"
	"fmt"

	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

// https://go.dev/wiki/Iota
const (
	PromptDateTime = iota
	PromptTime
	PromptChange
)

type InputContext struct {
	Prompt string
	Extra  map[string]interface{}
}

// Reusable prompt function
// Alan Turing was similing from above when this function was written I think.
func InputPrompt(scan *bufio.Scanner, ctx InputContext, specialPrompt ...int) (interface{}, error) {
	fmt.Println(ctx.Prompt)
	scan.Scan()
	if err := scan.Err(); err != nil {
		return "", err
	}
	input := scan.Bytes()

	if len(specialPrompt) != 0 {
		for _, elm := range specialPrompt {
			switch elm {
			case PromptDateTime:
				// handle date time
			case PromptTime:
				// handle time
			case PromptChange:
				for {
					// for loop shouldnt be starting with yes or no prompt
					confirmation, err := utils.YesOrNo(input, scan)
					if err != nil {
						fmt.Println("Invalid input, please type yes or no")
						continue
					}

					if confirmation {
						fmt.Printf("Please enter a new \x1b[1m%s\x1b[0m for the %s:\n", ctx.Extra["dataToBeChanged"], ctx.Extra["entity"])
						scan.Scan()
						if err := scan.Err(); err != nil {
							fmt.Println("Error reading input:", err)
							continue
						}

						return scan.Bytes(), nil
					} else {
						fmt.Println("No channges made.")
					}
				}
			}
		}
	}

	return string(input), nil
}
