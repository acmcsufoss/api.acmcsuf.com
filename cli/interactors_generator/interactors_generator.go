package interactors_generator

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	_ "embed"
)

//go:embed crudl.go.tmpl
var crudlTemplateString string

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

// CRUDLMethodsMap is a map of CRUDL methods to store methods.
type CRUDLMethodsMap map[CRUDLMethod]string

// InteractorConfig is the configuration for an interactor.
type InteractorConfig struct {
	// PascalName is the PascalCase name of the interactor.
	PascalName string `json:"pascalName"`

	// PascalPluralName is the PascalCase plural name of the interactor.
	// PascalPluralName string `json:"pascalPluralName"`

	// CamelName is the camelCase name of the interactor.
	CamelName string `json:"camelName"`

	// CamelPluralName is the camelCase plural name of the interactor.
	// CamelPluralName string `json:"camelPluralName"`

	// KabobName is the kabob-case name of the interactor.
	// KabobName string `json:"kabobName"`

	// KabobPluralName is the kabob-case plural name of the interactor.
	KabobPluralName string `json:"kabobPluralName"`

	// MethodsMap is the optional map of CRUDL methods to store methods.
	MethodsMap CRUDLMethodsMap `json:"methodsMap"`
}

// Render renders the interactor configuration to code.
func (c *InteractorConfig) Render(result *string, storePackage string) error {
	var b strings.Builder
	b.WriteString("// Use")
	b.WriteString(c.PascalName)
	b.WriteString(" uses a generated ")
	b.WriteString(c.PascalName)
	b.WriteString(" interactor.\n")
	b.WriteString(fmt.Sprintf("func Use%s(service *web.Service, store %s.Store) {\n", c.PascalName, storePackage))
	b.WriteString("\tuseCRUDL(\n")
	b.WriteString("\t\tservice,\n")
	b.WriteString("\t\twithPrefix(\"/")
	b.WriteString(c.KabobPluralName)
	b.WriteString("\"),\n")

	methodNames := []CRUDLMethod{CRUDLMethodCreate, CRUDLMethodRead, CRUDLMethodUpdate, CRUDLMethodDelete, CRUDLMethodList}
	for _, methodName := range methodNames {
		storeMethodName, ok := c.MethodsMap[methodName]
		if !ok {
			continue
		}

		b.WriteString(fmt.Sprintf("\t\twith%s(usecase.NewInteractor(func(ctx context.Context, request *%s.%s%sRequest, response *%s.%s%sResponse) (err error) {\n", methodName.String(), storePackage, storeMethodName, c.PascalName, storePackage, storeMethodName, c.PascalName))
		b.WriteString("\t\t\t*response, err = store.")
		b.WriteString(storeMethodName)
		b.WriteString(c.PascalName)
		b.WriteString("(request)\n")
		b.WriteString("\t\t\treturn err\n")
		b.WriteString("\t\t})),\n")
	}

	b.WriteString("\t)\n")
	b.WriteString("}\n")

	*result = b.String()
	return nil
}

// Config is the configuration for the interactors generator.
type Config struct {
	// Package is the package name.
	Package string `json:"package"`

	// StorePackageImport is the import path for the store package.
	StorePackageImport string `json:"storePackageImport"`

	// StorePackage is the package name for the store package.
	StorePackage string `json:"storePackage"`

	// DefaultMethodsMap is the default map of CRUDL methods to store methods.
	DefaultMethodsMap CRUDLMethodsMap `json:"defaultMethodsMap"`

	// Interactors is the list of interactors.
	Interactors []InteractorConfig `json:"interactors"`
}

// Render renders the configuration to code.
func (c *Config) Render(result *string) error {
	var b strings.Builder
	b.WriteString("// Code is generated. DO NOT EDIT.\n")
	b.WriteString("package ")
	b.WriteString(c.Package)
	b.WriteString("\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"context\"\n")
	b.WriteString("\t\"net/http\"\n\n")
	b.WriteString(fmt.Sprintf("\t\"%s\"\n", "github.com/swaggest/rest/nethttp"))
	b.WriteString(fmt.Sprintf("\t\"%s\"\n", "github.com/swaggest/rest/web"))
	b.WriteString(fmt.Sprintf("\t\"%s\"\n", "github.com/swaggest/usecase"))
	b.WriteString("\n\t\"")
	b.WriteString(c.StorePackageImport)
	b.WriteString("\"\n")
	b.WriteString(")\n\n")

	b.WriteString(crudlTemplateString)
	b.WriteString("\n")

	for _, interactorConfig := range c.Interactors {
		if interactorConfig.CamelName == "" {
			return fmt.Errorf("interactor %s has no camelName", interactorConfig.PascalName)
		}

		if interactorConfig.PascalName == "" {
			return fmt.Errorf("interactor %s has no pascalName", interactorConfig.PascalName)
		}

		if interactorConfig.KabobPluralName == "" {
			return fmt.Errorf("interactor %s has no kabobPluralName", interactorConfig.PascalName)
		}

		if interactorConfig.MethodsMap == nil || len(interactorConfig.MethodsMap) == 0 {
			interactorConfig.MethodsMap = c.DefaultMethodsMap
		}

		s := ""
		if err := interactorConfig.Render(&s, c.StorePackage); err != nil {
			log.Fatalf("Error rendering interactor: %v", err)
		}

		b.WriteString(s)
		b.WriteString("\n")
	}

	b.WriteString("// UseAll uses all interactors.\n")
	b.WriteString("func UseAll(service *web.Service, store ")
	b.WriteString(c.StorePackage)
	b.WriteString(".Store) {\n")
	for _, interactorConfig := range c.Interactors {
		b.WriteString(fmt.Sprintf("\tUse%s(service, store)\n", interactorConfig.PascalName))
	}
	b.WriteString("}\n")

	*result = b.String()
	return nil
}

// UnmarshalConfig unmarshals the configuration.
func UnmarshalConfig(data []byte, c *Config) error {
	return json.Unmarshal(data, c)
}

// ReadConfigFile reads the configuration file.
func ReadConfigFile(path string, c *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	if err := UnmarshalConfig(data, c); err != nil {
		log.Fatalf("Error unmarshaling configuration: %v", err)
	}

	return nil
}
