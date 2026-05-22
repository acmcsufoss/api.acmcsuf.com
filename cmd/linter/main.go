package main

import (
	"errors"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
)

var filesToSkip []string = []string{".", "swagger.go"}

func main() {

	// Go to api/handlers
	// This should be run in the project root
	handlersPath := "internal/api/handlers/"
	os.Chdir(handlersPath)

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if slices.Contains(filesToSkip, path) {
			return nil
		}

		// Creating an AST tree by parsing
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		// Now we can search for specific things like functions and imports
		ast.Inspect(file, func(n ast.Node) bool {
			// looking for an import, we shouldn't have dbmodels in handler
			switch x := n.(type) {
			case *ast.ImportSpec:

				importedModule, _ := strconv.Unquote(x.Path.Value)

				// yah messy, but for the sake of proposal
				check := (importedModule == "github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels")
				if check {
					// This error is going to make me throw up lol
					log.Println(errors.New("Bad import found: " + importedModule + " in " + handlersPath + path))
					os.Exit(1)
				}
			}
			return true
		})
		return nil
	})
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
