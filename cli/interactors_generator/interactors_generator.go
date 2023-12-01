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

// CRUDLMethodPrefix is the prefix of a CRUDL method name.
type CRUDLMethodPrefix string

const (
	// CRUDLMethodPrefixCreate is the create method name.
	CRUDLMethodPrefixCreate CRUDLMethodPrefix = "Create"

	// CRUDLMethodPrefixRead is the read method name.
	CRUDLMethodPrefixRead CRUDLMethodPrefix = "Read"

	// CRUDLMethodPrefixUpdate is the update method name.
	CRUDLMethodPrefixUpdate CRUDLMethodPrefix = "Update"

	// CRUDLMethodPrefixDelete is the delete method name.
	CRUDLMethodPrefixDelete CRUDLMethodPrefix = "Delete"

	// CRUDLMethodPrefixList is the list method name.
	CRUDLMethodPrefixList CRUDLMethodPrefix = "List"
)

var (
	crudlMethodPrefixes = []CRUDLMethodPrefix{CRUDLMethodPrefixCreate, CRUDLMethodPrefixRead, CRUDLMethodPrefixUpdate, CRUDLMethodPrefixDelete, CRUDLMethodPrefixList}
)

// Pascal returns the PascalCase representation of the CRUDL method.
func (m CRUDLMethodPrefix) Pascal() string {
	return string(m)
}

// Imperative returns the imperative representation of the CRUDL method.
func (m CRUDLMethodPrefix) Imperative() string {
	switch m {
	case CRUDLMethodPrefixCreate:
		return "creates"
	case CRUDLMethodPrefixRead:
		return "reads"
	case CRUDLMethodPrefixUpdate:
		return "updates"
	case CRUDLMethodPrefixDelete:
		return "deletes"
	case CRUDLMethodPrefixList:
		return "lists"
	default:
		return ""
	}
}

// CRUDLMethodPrefixesMap is a map of CRUDL methods to store methods.
type CRUDLMethodPrefixesMap map[CRUDLMethodPrefix]string

// InteractorConfig is the configuration for an interactor.
type InteractorConfig struct {
	// PascalName is the PascalCase name of the interactor.
	PascalName string `json:"pascalName"`

	// PascalPluralName is the PascalCase plural name of the interactor.
	PascalPluralName string `json:"pascalPluralName"`

	// CamelName is the camelCase name of the interactor.
	CamelName string `json:"camelName"`

	// CamelPluralName is the camelCase plural name of the interactor.
	// CamelPluralName string `json:"camelPluralName"`

	// KabobName is the kabob-case name of the interactor.
	// KabobName string `json:"kabobName"`

	// KabobPluralName is the kabob-case plural name of the interactor.
	KabobPluralName string `json:"kabobPluralName"`

	// MethodPrefixesMap is the optional map of CRUDL method prefixes to store method prefixes.
	MethodPrefixesMap CRUDLMethodPrefixesMap `json:"methodsMap"`
}

func (c *InteractorConfig) Types(methodPrefix, typesPackage string) (string, string, string) {
	method := methodPrefix + c.PascalName
	typePrefix := method
	if typesPackage != "" {
		typePrefix = typesPackage + "." + typePrefix
	}
	requestType, responseType := typePrefix+"Request", typePrefix+"Response"
	return method, requestType, responseType
}

// Render renders the interactor configuration to code.
func (c *InteractorConfig) Render(result *string, typesPackage string) error {
	var b strings.Builder
	b.WriteString("// Use")
	b.WriteString(c.PascalPluralName)
	b.WriteString(" uses a generated ")
	b.WriteString(c.PascalPluralName)
	b.WriteString(" interactor.\n")
	b.WriteString(fmt.Sprintf("func Use%s(service *web.Service, store Store) {\n", c.PascalPluralName))
	b.WriteString("\tuseCRUDL(\n")
	b.WriteString("\t\tservice,\n")
	b.WriteString("\t\twithPrefix(\"/")
	b.WriteString(c.KabobPluralName)
	b.WriteString("\"),\n")

	for _, methodPrefix := range crudlMethodPrefixes {
		storeMethodPrefix, ok := c.MethodPrefixesMap[methodPrefix]
		if !ok {
			continue
		}

		storeMethod, requestType, responseType := c.Types(storeMethodPrefix, typesPackage)
		b.WriteString(fmt.Sprintf(
			"\t\twith%s(usecase.NewInteractor(func(ctx context.Context, request *%s, response *%s) (err error) {\n",
			methodPrefix.Pascal(),
			requestType,
			responseType,
		))
		b.WriteString(fmt.Sprintf(
			"\t\t\t*response, err = store.%s(request)\n",
			storeMethod,
		))
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
	// Package is the package name. Package must have a length greater than 0.
	Package string `json:"package"`

	// TypesPackage is the name of the external package containing the request and response types. Default is "".
	TypesPackage string `json:"typesPackage"`

	// TypesPackageImport is the import path of the external package containing the request and response types. Default is "".
	TypesPackageImport string `json:"typesPackageImport"`

	// DefaultMethodsMap is the default map of CRUDL methods to store methods.
	DefaultMethodPrefixesMap CRUDLMethodPrefixesMap `json:"defaultMethodPrefixesMap"`

	// Interactors is the list of interactors.
	Interactors []InteractorConfig `json:"interactors"`
}

// Render renders the configuration to code.
func (c *Config) Render(result *string) error {
	if c.Package == "" {
		return fmt.Errorf("package must be set")
	}

	if c.Package == c.TypesPackage {
		return fmt.Errorf("package and typesPackage must be different")
	}

	var b strings.Builder
	b.WriteString("// Code is generated. DO NOT EDIT.\n\n")
	b.WriteString("package ")
	b.WriteString(c.Package)
	b.WriteString("\n\n")
	b.WriteString("import (\n")
	b.WriteString("\t\"context\"\n")
	b.WriteString("\t\"io\"\n")
	b.WriteString("\t\"net/http\"\n\n")
	b.WriteString(fmt.Sprintf("\t\"%s\"\n", "github.com/swaggest/rest/nethttp"))
	b.WriteString(fmt.Sprintf("\t\"%s\"\n", "github.com/swaggest/rest/web"))
	b.WriteString(fmt.Sprintf("\t\"%s\"\n", "github.com/swaggest/usecase"))
	if c.TypesPackage != "" {
		if c.TypesPackageImport == "" {
			return fmt.Errorf("typesPackageImport must be set if typesPackage is set")
		}

		b.WriteString(fmt.Sprintf("\t%s \"%s\"\n", c.TypesPackageImport, c.TypesPackage))
	}
	b.WriteString(")\n\n")

	b.WriteString(crudlTemplateString)
	b.WriteString("\n")

	b.WriteString("// ContainsContext can be embedded by any interface to have an overrideable\n")
	b.WriteString("// context.\n")
	b.WriteString("type ContainsContext interface {\n")
	b.WriteString("\tWithContext(context.Context) ContainsContext\n")
	b.WriteString("}\n\n")

	b.WriteString("// Store is the store interface.\n")
	b.WriteString("type Store interface {\n")
	b.WriteString("\tio.Closer\n")
	b.WriteString("\tContainsContext\n\n")
	for _, interactorConfig := range c.Interactors {
		for _, methodPrefix := range crudlMethodPrefixes {
			storeMethodPrefix, ok := interactorConfig.MethodPrefixesMap[methodPrefix]
			if !ok {
				continue
			}

			storeMethod, requestType, responseType := interactorConfig.Types(storeMethodPrefix, c.TypesPackage)
			b.WriteString(fmt.Sprintf(
				"\t// %s %s %s.\n",
				storeMethod,
				methodPrefix.Imperative(),
				interactorConfig.PascalName,
			))
			b.WriteString(fmt.Sprintf(
				"\t%s(r %s) (*%s, error)\n\n",
				storeMethod,
				requestType,
				responseType,
			))
		}
	}
	b.WriteString("}\n\n")

	s := ""
	for _, interactorConfig := range c.Interactors {
		if interactorConfig.CamelName == "" {
			return fmt.Errorf("interactor %s has no camelName", interactorConfig.PascalName)
		}

		if interactorConfig.PascalName == "" {
			return fmt.Errorf("interactor %s has no pascalName", interactorConfig.PascalName)
		}

		if interactorConfig.PascalPluralName == "" {
			return fmt.Errorf("interactor %s has no pascalPluralName", interactorConfig.PascalName)
		}

		if interactorConfig.KabobPluralName == "" {
			return fmt.Errorf("interactor %s has no kabobPluralName", interactorConfig.PascalName)
		}

		if interactorConfig.MethodPrefixesMap == nil || len(interactorConfig.MethodPrefixesMap) == 0 {
			interactorConfig.MethodPrefixesMap = c.DefaultMethodPrefixesMap
		}

		if err := interactorConfig.Render(&s, c.TypesPackage); err != nil {
			log.Fatalf("Error rendering interactor: %v", err)
		}
		b.WriteString(s)
		b.WriteString("\n")
	}

	b.WriteString("// UseAll uses all interactors.\n")
	b.WriteString("func UseAll(service *web.Service, store Store) {\n")
	for _, interactorConfig := range c.Interactors {
		b.WriteString(fmt.Sprintf("\tUse%s(service, store)\n", interactorConfig.PascalPluralName))
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
