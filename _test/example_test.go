package _test

import (
	"DocumentConverter/converter"
	"DocumentConverter/element"
	"DocumentConverter/parser"
	"DocumentConverter/printer"
	"bufio"
	"strings"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestExample(t *testing.T) {
	exampleDefinitions := map[string]element.Definer{
		"P": element.NewElement("P", "person", []string{"firstname", "lastname"}, []string{"T", "A", "F"}),
		"A": element.NewElement("A", "address", []string{"street", "city", "zip"}, nil),
		"T": element.NewElement("T", "phone", []string{"mobile", "landline"}, nil),
		"F": element.NewElement("F", "family", []string{"name", "born"}, []string{"T", "A"}),
	}

	exampleInput := `P|Elof|Sundin
T|073-101801|018-101801
A|S:t Johannesgatan 16|Uppsala|75330
F|Hans|1967
A|Frodegatan 13B|Uppsala|75325
F|Anna|1969
T|073-101802|08-101802
P|Boris|Johnson
A|10 Downing Street|London`

	converter := converter.Converter{
		LineParser: &parser.SeparatorLineParser{Separator: "|", ElementDefinitions: exampleDefinitions},
		Printer:    printer.NewXMLPrinter(),
	}
	scanner := bufio.NewScanner(strings.NewReader(exampleInput))

	result, _ := converter.ConvertDocument(*element.NewRoot("people"), scanner)

	approvals.VerifyString(t, result)
}
