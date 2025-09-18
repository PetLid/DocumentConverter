package printer

import (
	"DocumentConverter/models"
	"fmt"
	"strings"
)

type XMLPrinter struct {
	indentationSpaces int
}

func NewXMLPrinter() *XMLPrinter {
	return &XMLPrinter{indentationSpaces: 4}
}

func (printer *XMLPrinter) Print(root models.DocumentNode) string {
	result := ""
	printer.printNode(root, 0, &result)

	return result
}

func (printer *XMLPrinter) printNode(node models.DocumentNode, indentationLevel int, result *string) {

	*result += fmt.Sprintf("%s<%s>\n", strings.Repeat(" ", indentationLevel*printer.indentationSpaces), node.Name)

	for _, attribute := range node.Attributes {
		*result += fmt.Sprintf("%s<%[2]s>%[3]s</%[2]s>\n", strings.Repeat(" ", (indentationLevel+1)*printer.indentationSpaces), attribute.Name, attribute.Value)
	}

	indentationLevel++

	for _, child := range node.Children {
		printer.printNode(child, indentationLevel, result)
	}

	indentationLevel--
	*result += fmt.Sprintf("%s</%s>\n", strings.Repeat(" ", indentationLevel*printer.indentationSpaces), node.Name)
}
