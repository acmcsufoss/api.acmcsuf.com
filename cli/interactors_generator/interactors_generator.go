package interactors_generator

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	_ "embed"
)

//go:embed crudl.go.tmpl
var crudlTemplateString string

const (
	swaggestPackageImport = "github.com/swaggest/usecase"
)

// CRUDLMethod is a CRUDL method name.
type CRUDLMethod string

const (
	// CRUDLMethodCreate is the create method name.
	CRUDLMethodCreate CRUDLMethod = "Create"

	// CRUDLMethodRead is the read method name.
	CRUDLMethodRead CRUDLMethod = "Read"

	// CRUDLMethodUpdate is the update method name.
	CRUDLMethodUpdate CRUDLMethod = "Update"

	// CRUDLMethodDelete is the delete method name.
	CRUDLMethodDelete CRUDLMethod = "Delete"

	// CRUDLMethodList is the list method name.
	CRUDLMethodList CRUDLMethod = "List"
)

func (m CRUDLMethod) String() string {
	return string(m)
}

// CRUDLMethodsMap is a map of CRUDL method names to their corresponding store method names.
type CRUDLMethodsMap map[CRUDLMethod]string

// InteractorConfig is the configuration for a single interactor.
type InteractorConfig struct {
	//  PascalName is the name of the interactor in PascalCase.
	PascalName string `json:"pascalName"`

	// CamelName is the name of the interactor in camelCase.
	CamelName string `json:"camelName"`

	// KabobName is the name of the interactor in kabob-case.
	KabobName string `json:"kabobName"`

	// MethodsMap is a map of method names to their corresponding store method names.
	MethodsMap CRUDLMethodsMap `json:"methodsMap"`
}

// Render generates the source code for the interactor.
func (c *InteractorConfig) Render(result *string, storePackage string) error {
	var b strings.Builder
	b.WriteString("func use")
	b.WriteString(c.PascalName)
	b.WriteString("(service *web.Service, store ")
	b.WriteString(storePackage)
	b.WriteString(".Store")
	b.WriteString(") {\n")
	b.WriteString("\tuseCRUDL(\n")
	b.WriteString("\t\tservice,\n")
	b.WriteString("\t\twithPrefix(\"/")
	b.WriteString(c.KabobName)
	b.WriteString("\"),\n")

	for methodName, storeMethodName := range c.MethodsMap {
		b.WriteString("\t\twith")
		b.WriteString(methodName.String())
		b.WriteString("(usecase.NewInteractor(func(ctx context.Context, input ")
		b.WriteString(storePackage)
		b.WriteString(".")
		b.WriteString(storeMethodName)
		b.WriteString("Input) (output ")
		b.WriteString(storePackage)
		b.WriteString(".")
		b.WriteString(storeMethodName)
		b.WriteString("Output, err error) {\n")
		b.WriteString("\t\t\treturn store.")
		b.WriteString(storeMethodName)
		b.WriteString("(ctx, input)\n")
		b.WriteString("\t\t})),\n")
	}

	b.WriteString("\t)\n")
	b.WriteString("}\n")

	*result = b.String()
	return nil
}

// Config is the configuration for the interactors generator.
type Config struct {
	// Package is the package name of the generated file.
	Package string `json:"package"`

	// StorePackageImport is the import path string for the store package e.g. "github.com/example/api.example.com/store".
	StorePackageImport string `json:"storePackageImport"`

	// StorePackage is the package name of the store package e.g. "store".
	StorePackage string `json:"storePackage"`

	// DefaultMethodsMap is a map of method names to their corresponding store method names.
	DefaultMethodsMap CRUDLMethodsMap `json:"defaultMethodsMap"`

	// Interactors is a list of interactors to generate.
	Interactors []InteractorConfig `json:"interactors"`
}

// Render generates the source code for the configured interactors.
func (c *Config) Render(result *string) error {
	var b strings.Builder
	b.WriteString("// Code is generated. DO NOT EDIT.\n\n")
	b.WriteString("package ")
	b.WriteString(c.Package)
	b.WriteString("\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"context\"\n\n")
	b.WriteString("\t\"net/http\"\n\n")
	b.WriteString("\t\"" + swaggestPackageImport + "/nethttp\"\n")
	b.WriteString("\t\"" + swaggestPackageImport + "/rest/web\"\n")
	b.WriteString("\t\"" + swaggestPackageImport + "/usecase\"\n\n")
	b.WriteString("\t\"")
	b.WriteString(c.StorePackageImport)
	b.WriteString("\"\n")
	b.WriteString(")\n\n")

	var s string
	for _, interactorConfig := range c.Interactors {
		if interactorConfig.PascalName == "" {
			return fmt.Errorf("PascalName (example: 'PascalName') is empty")
		}

		if interactorConfig.CamelName == "" {
			return fmt.Errorf("CamelName (example: 'camelName') is empty")
		}

		if interactorConfig.KabobName == "" {
			return fmt.Errorf("KabobName (example: 'kabob-name') is empty")
		}

		if interactorConfig.MethodsMap == nil || len(interactorConfig.MethodsMap) == 0 {
			interactorConfig.MethodsMap = c.DefaultMethodsMap
		}

		err := interactorConfig.Render(&s, c.StorePackage)
		if err != nil {
			return err
		}

		b.WriteString(s)
		b.WriteString("\n")
	}

	// Append crudl.go.tmpl to the string builder.
	b.WriteString("\n")
	b.WriteString(crudlTemplateString)

	// Write the generated code to the result.
	*result = b.String()
	return nil
}

// UnmarshalConfig unmarshals the given JSON data into the given Config.
func UnmarshalConfig(data []byte, c *Config) error {
	return json.Unmarshal(data, c)
}

// ReadConfigFile reads the given file and unmarshals it into the given Config.
func ReadConfigFile(path string, c *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return UnmarshalConfig(data, c)
}
