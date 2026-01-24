package main

import (
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/cli"
)

func main() {
	cli.Menu()
	exitCode := cli.Execute()
	os.Exit(int(exitCode))
}
