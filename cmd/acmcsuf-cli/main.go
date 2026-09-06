package main

import (
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli/command"
)

func main() {
	exitCode := command.Execute()
	os.Exit(int(exitCode))
}
