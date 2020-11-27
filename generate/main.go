package main

import (
	"bufio"
	"flag"
	"github.com/ghodss/yaml"
	"go/parser"
	"go/token"
	"regexp"
	"text/template"
	"io/ioutil"
	"log"
	"os"
)

type DelegateRegistration struct {
	Filename string
	Function string
	Export string
}

type DelegateConfig struct {
	Name string
	Version string
	Description string
	Module string
	Delegates []DelegateRegistration
}

type DelegateTemplateValue struct {
	Name string
	Handler string
	Package string
}

type TemplateInformation struct {
	Version string
	Description string
	Module string
	Package string
	Delegates []DelegateTemplateValue
}

func main() {
	var configFilePath= ""
	var outputPath = ""

	flag.StringVar(&configFilePath, "config", "","the delegate configuration to use")
	flag.StringVar(&outputPath, "output", "server.go","the output to write to")
	flag.Parse()

	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("Could not open file '%v': %v", configFilePath, err)
	}

	delegateConfig := DelegateConfig{}
	err = yaml.Unmarshal(configFile, &delegateConfig)
	if err != nil {
		log.Fatalf("Could not parse config yaml: %v", err)
	}
	tplValues := TemplateInformation{
		Version: delegateConfig.Version,
		Description: delegateConfig.Description,
		Delegates: make([]DelegateTemplateValue, 0),
	}
	for _, delegate := range delegateConfig.Delegates {
		info := getDelegateInformation(delegate)
		tplValues.Delegates = append(tplValues.Delegates, info)
		tplValues.Module = delegateConfig.Module
		tplValues.Package = info.Package
	}
	templateContent, err := template.New("server.go").Parse(TEMPLATE)
	if err != nil {
		log.Fatalf("Could not parse template: %v",  err)
	}
	outFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Could not create output '%v': %v", outputPath, err)
	}
	err = templateContent.Execute(outFile, tplValues)
	if err != nil {
		log.Fatalf("Error writing main file '%v': %v", outputPath, err)
	}

}

func getDelegateInformation(delegate DelegateRegistration) DelegateTemplateValue {
	result := DelegateTemplateValue{}

	fset := token.NewFileSet()
	src, err := os.Open(delegate.Filename)
	if err != nil {
		log.Fatalf("Could not open referenced source file %v", delegate);
	}
	reader := bufio.NewScanner(src)
	f, err := parser.ParseFile(fset, delegate.Filename, nil, parser.PackageClauseOnly)
	if err != nil {
		log.Fatalf("Error parsing file %v: %v", delegate.Filename, err)
	}
	line := fset.Position(f.Package).Line - 1
	for reader.Scan() {
		if line == 0 {
			re, _ := regexp.Compile("package *(\\w*)")
			result.Package  = re.FindStringSubmatch(reader.Text())[1]
			break
		}
		line--
	}
	result.Handler = delegate.Function
	result.Name = delegate.Export
	return result
}