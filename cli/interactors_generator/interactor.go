package main

import (
	"encoding/json"
	"os"
	"strings"
)

const (
	swaggest_package_import = "github.com/swaggest/usecase"
)

// TODO: Copy CRUDL source code from ./crudl.go to generated file.

type interactorConfig struct {
	//  pascalName is the name of the interactor in PascalCase.
	pascalName string `json:"pascalName"`

	// camelName is the name of the interactor in camelCase.
	camelName string `json:"camelName"`

	// kabobName is the name of the interactor in kabob-case.
	kabobName string `json:"kabobName"`

	// methodNamesMap is a map of method names to their corresponding store method names.
	methodNamesMap map[string]string `json:"methodNamesMap"`
}

func (c *interactorConfig) render() string {
	var b strings.Builder
	b.WriteString("func use")
	b.WriteString(c.pascalName)
	b.WriteString("(service *web.Service, store api.Store) {\n")
	b.WriteString("\tuseCRUDL(\n")
	b.WriteString("\t\tservice,\n")
	b.WriteString("\t\twithPrefix(\"/")
	b.WriteString(c.kabobName)
	b.WriteString("\"),\n")

	for methodName, storeMethodName := range c.methodNamesMap {
		if methodName == "" {
			continue
		}

		b.WriteString("\t\twith")
		b.WriteString(methodName)
		b.WriteString("(usecase.NewInteractor(func(ctx context.Context, input api.")
		b.WriteString(methodName)
		b.WriteString(c.pascalName)
		b.WriteString("Request, output api.")
		b.WriteString(methodName)
		b.WriteString(c.pascalName)
		b.WriteString("Response) error {\n")
		b.WriteString("\t\t\t_, err := store.")
		b.WriteString(storeMethodName)
		b.WriteString(c.pascalName)
		b.WriteString("(input)\n")
		b.WriteString("\t\t\tif err != nil {\n")
		b.WriteString("\t\t\t\treturn err\n")
		b.WriteString("\t\t\t}\n")
		b.WriteString("\n")
		b.WriteString("\t\t\treturn nil\n")
		b.WriteString("\t\t})),\n")
	}

	b.WriteString("\t)\n")
	b.WriteString("}\n")

	return b.String()
}

type interactorsConfig struct {
	// packageName is the package name of the generated file.
	packageName string `json:"packageName"`

	// storePackageImport is the import path for the store package.
	storePackageImport string `json:"storePackageImport"`

	// defaultMethodNamesMap is a map of method names to their corresponding store method names.
	defaultMethodNamesMap map[string]string `json:"defaultMethodNamesMap"`

	//  interactors is a list of interactors to generate.
	interactors []interactorConfig `json:"interactors"`
}

func (c *interactorsConfig) render() string {
	var b strings.Builder
	b.WriteString("// Code is generated. DO NOT EDIT.\n\n")
	b.WriteString("package ")
	b.WriteString(c.packageName)
	b.WriteString("\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"")
	b.WriteString(c.storePackageImport)
	b.WriteString("\"\n")
	b.WriteString("\t\"")
	b.WriteString(swaggest_package_import)
	b.WriteString("\"\n")
	b.WriteString(")\n\n")

	for _, interactor := range c.interactors {
		if interactor.methodNamesMap == nil {
			interactor.methodNamesMap = c.defaultMethodNamesMap
		}

		b.WriteString(interactor.render())
		b.WriteString("\n")
	}

	return b.String()
}

func unmarshalInteractors(data []byte, c *interactorsConfig) error {
	return json.Unmarshal(data, c)
}

func readInteractorsFromFile(path string, c *interactorsConfig) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return unmarshalInteractors(data, c)
}
