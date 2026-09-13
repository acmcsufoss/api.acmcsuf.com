package main

import (
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/linter"
)

func main() {
	rules := linter.LinterRules()

	for _, rule := range rules {
		linter.Lint(rule)
	}

	os.Exit(0)
}
