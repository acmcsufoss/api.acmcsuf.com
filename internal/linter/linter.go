package linter 

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

func Lint(rule lintRule) {
	
	os.Chdir(rule.path)

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if slices.Contains(rule.skipFiles, path) {
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
			switch x := n.(type) {

			// Import check
			case *ast.ImportSpec:

				importedModule, _ := strconv.Unquote(x.Path.Value)

				if  slices.Contains(rule.badImports, importedModule) {
					log.Println(errors.New("Bad import found: " + importedModule + " in " + rule.path + path))
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
