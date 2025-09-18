package builder

import (
	"DocumentConverter/element"
	"DocumentConverter/models"
)

type builderNode struct {
	elementDefinition element.Definer
	children          []*builderNode
	attributes        []models.Attribute
	parentNode        *builderNode
}

func (builderNode *builderNode) mapToDocumentNode() models.DocumentNode {
	mappedChildren := make([]models.DocumentNode, 0, len(builderNode.children))

	for _, child := range builderNode.children {
		mappedChildren = append(mappedChildren, child.mapToDocumentNode())
	}

	definition := builderNode.elementDefinition

	return models.DocumentNode{Name: definition.GetIdentifier(), Attributes: builderNode.attributes, Children: mappedChildren}
}
