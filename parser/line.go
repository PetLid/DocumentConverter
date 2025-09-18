package parser

import (
	"DocumentConverter/element"
	"DocumentConverter/models"
	"errors"
	"fmt"
	"log"
	"strings"
)

type LineParser interface {
	ParseLine(line string) (*models.Line, error)
}

type SeparatorLineParser struct {
	Separator          string
	ElementDefinitions map[string]element.Definer
}

func (parser *SeparatorLineParser) ParseLine(line string) (*models.Line, error) {
	parts := strings.Split(line, parser.Separator)

	identifier := parts[0]
	elementDefinition := parser.ElementDefinitions[identifier]
	if elementDefinition == nil {
		return nil, errors.New(fmt.Sprintf("Element definition not found for identifier \"%s\".", identifier))
	}

	attributeValues := parts[1:]
	attributes := parser.parseAttributes(line, attributeValues, elementDefinition)

	return &models.Line{
		Element:    elementDefinition,
		Attributes: attributes,
	}, nil
}

func (parser *SeparatorLineParser) parseAttributes(line string, attributeValues []string, elementDefinition element.Definer) []models.Attribute {
	attributeNames := elementDefinition.GetAttributeNames()

	if len(attributeValues) < len(attributeNames) {
		log.Printf("Line \"%s\" contained fewer attributes values than definition (\"%s\")", line, strings.Join(attributeNames[len(attributeValues):], "\", \""))
	}

	if len(attributeValues) > len(attributeNames) {
		log.Printf("Line \"%s\" contained additional attribute value(s) not named in definition (\"%s\"), check definitions", line, strings.Join(attributeValues[len(attributeNames):], "\", \""))
	}

	nAttributes := min(len(attributeNames), len(attributeValues))

	attributes := make([]models.Attribute, 0, nAttributes)

	for i := range nAttributes {
		attributes = append(attributes, models.Attribute{
			Name:  attributeNames[i],
			Value: attributeValues[i],
		})
	}
	return attributes
}
