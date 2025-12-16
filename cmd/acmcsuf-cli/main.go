package main

import (
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli"
)

func main() {
	exitCode := cli.Execute()
	os.Exit(int(exitCode))
}
