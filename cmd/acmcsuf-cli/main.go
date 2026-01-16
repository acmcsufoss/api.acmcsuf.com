package main

import (
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli"
)

func main() {
	exitCode := cli.Execute()
	cli.Menu()
	os.Exit(int(exitCode))
}
