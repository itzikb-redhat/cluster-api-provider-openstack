package main

import (
	_ "embed"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	defaultYear       = "2024"
	defaultAPIVersion = "v1alpha1"
)

//go:embed data/api.template
var api_template string

type specExtraValidation struct {
	Rule    string
	Message string
}

type templateFields struct {
	APIVersion           string
	Year                 string
	Name                 string
	SpecExtraType        string
	StatusExtraType      string
	SpecExtraValidations []specExtraValidation
}

var allResources []templateFields = []templateFields{
	{
		Name: "Image",
		SpecExtraValidations: []specExtraValidation{
			{
				Rule:    "!has(self.__import__) ? has(self.resource.content) : true",
				Message: "resource content must be specified when not importing",
			},
		},
		StatusExtraType: "ImageStatusExtra",
	},
	{
		Name:       "Flavor",
		APIVersion: "v1alpha1",
	},
	{
		Name:       "Network",
		APIVersion: "v1alpha1",
	},
	{
		Name:          "Subnet",
		SpecExtraType: "SubnetRefs",
	},
	{
		Name: "Router",
	},
	{
		Name:          "Port",
		SpecExtraType: "PortRefs",
	},
	{
		Name: "SecurityGroup",
	},
	{
		Name: "Server",
	},
}

func main() {
	apiTemplate := template.Must(template.New("api").Parse(api_template))

	for i := range allResources {
		resource := &allResources[i]

		if resource.Year == "" {
			resource.Year = defaultYear
		}

		if resource.APIVersion == "" {
			resource.APIVersion = defaultAPIVersion
		}

		resourceLower := strings.ToLower(resource.Name)

		apiPath := filepath.Join("api", resource.APIVersion, "zz_generated."+resourceLower+"-resource.go")
		apiFile, err := os.Create(apiPath)
		if err != nil {
			panic(err)
		}
		defer func() {
			err := apiFile.Close()
			if err != nil {
				panic(err)
			}
		}()

		err = writeAutogeneratedHeader(apiFile)
		if err != nil {
			panic(err)
		}

		err = apiTemplate.Execute(apiFile, resource)
		if err != nil {
			panic(err)
		}
	}
}

func writeAutogeneratedHeader(f *os.File) error {
	_, err := f.WriteString("// Code generated by resource-generator. DO NOT EDIT.\n")

	return err
}
