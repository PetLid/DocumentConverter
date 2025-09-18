package main

import (
	"DocumentConverter/converter"
	"DocumentConverter/element"
	"DocumentConverter/parser"
	"DocumentConverter/printer"
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

func main() {
	inputFileName, definitionsFileName, outputFileName := parseFlags()

	definitions := loadDefinitions(definitionsFileName)

	err, result := convertFile(inputFileName, definitions)

	if err != nil {
		log.Fatalf("Could not convert document: %s", err)
	}

	outputDocument(outputFileName, result)
}

func parseFlags() (string, string, string) {
	var inputFileName string
	flag.StringVar(&inputFileName, "input-file", "input", "file to convert")

	var definitionsFileName string
	flag.StringVar(&definitionsFileName, "definitions-file", "definitions.yaml", "JSON file with element definitions")

	var outputFileName string
	flag.StringVar(&outputFileName, "output-file", "output.xml", "name of generated XML file")

	var rootName string
	flag.StringVar(&rootName, "root-element", "people", "name of the root element")

	flag.Parse()

	return inputFileName, definitionsFileName, outputFileName
}

func loadDefinitions(definitionsFile string) map[string]element.Definer {
	definitionsYaml, err := os.ReadFile(definitionsFile)
	if err != nil {
		log.Fatalf("Could not read definition file %s: %s", definitionsFile, err)
	}

	var definitions []struct {
		LegacyIdentifier       string   `yaml:"LegacyIdentifier"`
		NewIdentifier          string   `yaml:"NewIdentifier"`
		AttributeNames         []string `yaml:"AttributeNames"`
		ChildLegacyIdentifiers []string `yaml:"ChildLegacyIdentifiers"`
	}

	if err := yaml.Unmarshal(definitionsYaml, &definitions); err != nil {
		log.Fatalf("Could not unmarshall definitions in file %s: %s", definitionsFile, err)
	}

	mappedDefinitions := make(map[string]element.Definer, len(definitions))
	for _, definition := range definitions {
		mappedDefinitions[definition.LegacyIdentifier] = element.NewElement(definition.LegacyIdentifier, definition.NewIdentifier, definition.AttributeNames, definition.ChildLegacyIdentifiers)
	}

	return mappedDefinitions
}

func convertFile(inputFileName string, definitions map[string]element.Definer) (error, string) {
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("Failed to open file %s, %s", "inputFile", err)
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			log.Fatalf("Failed to close file %s, %s", inputFileName, err)
		}
	}(inputFile)

	documentConverter := converter.Converter{
		LineParser: &parser.SeparatorLineParser{Separator: "|", ElementDefinitions: definitions},
		Printer:    printer.NewXMLPrinter(),
	}
	scanner := bufio.NewScanner(inputFile)
	result, err := documentConverter.ConvertDocument(*element.NewRoot("people"), scanner)

	return err, result
}

func outputDocument(outputFileFlagValue string, outputString string) {
	err := os.WriteFile(outputFileFlagValue, []byte(outputString), 0770)
	if err != nil {
		log.Fatalf("Failed to write outputfile: %s", err)
	}
}
