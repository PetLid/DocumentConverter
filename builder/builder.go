package builder

import (
	"DocumentConverter/element"
	"DocumentConverter/models"
)

func BuildDocumentTree(lineData []models.Line, rootElement element.Root) models.DocumentNode {
	builder := newDocumentBuilder(rootElement)

	for _, line := range lineData {
		builder.addLine(line)
	}

	return builder.root.mapToDocumentNode()
}

type documentBuilder struct {
	root             *builderNode
	lastAddedElement *builderNode
}

func newDocumentBuilder(rootElement element.Root) *documentBuilder {
	root := &builderNode{
		elementDefinition: &rootElement,
		children:          nil,
		parentNode:        nil,
	}

	return &documentBuilder{
		root:             root,
		lastAddedElement: root,
	}
}

func (builder *documentBuilder) elementNodeFromLine(line models.Line) *builderNode {
	elementNode := &builderNode{
		elementDefinition: line.Element,
		attributes:        line.Attributes,
		children:          nil,
	}

	return elementNode
}

func (builder *documentBuilder) addLine(line models.Line) {
	elementNode := builder.elementNodeFromLine(line)
	builder.addToParent(elementNode)
	builder.lastAddedElement = elementNode
}

func (builder *documentBuilder) addToParent(elementNode *builderNode) {
	parentCandidate := builder.lastAddedElement

	for parentCandidate != nil {
		if parentCandidate.elementDefinition.CanBeParentOf(elementNode.elementDefinition) {
			elementNode.parentNode = parentCandidate
			parentCandidate.children = append(parentCandidate.children, elementNode)
			break
		} else {
			parentCandidate = parentCandidate.parentNode
		}
	}
}
