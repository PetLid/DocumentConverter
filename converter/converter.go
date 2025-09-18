package converter

import (
	"DocumentConverter/builder"
	"DocumentConverter/element"
	"DocumentConverter/models"
	"DocumentConverter/parser"
	"DocumentConverter/printer"
	"bufio"
)

type Converter struct {
	LineParser parser.LineParser
	Printer    printer.Printer
}

func (converter Converter) ConvertDocument(rootElement element.Root, scanner *bufio.Scanner) (string, error) {
	var lineData []models.Line

	for scanner.Scan() {
		data, err := converter.LineParser.ParseLine(scanner.Text())

		if err != nil {
			return "", err
		}

		lineData = append(lineData, *data)
	}

	documentRoot := builder.BuildDocumentTree(lineData, rootElement)

	return converter.Printer.Print(documentRoot), nil
}
