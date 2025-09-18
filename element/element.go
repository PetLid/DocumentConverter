package element

import "slices"

type Normal struct {
	legacyIdentifier       string
	newIdentifier          string
	attributeNames         []string
	childLegacyIdentifiers []string
}

func NewElement(legacyIdentifier string, identifier string, attributeNames []string, childIdentifiers []string) *Normal {
	return &Normal{legacyIdentifier: legacyIdentifier, newIdentifier: identifier, attributeNames: attributeNames, childLegacyIdentifiers: childIdentifiers}
}

func (element *Normal) getLegacyIdentifier() string {
	return element.legacyIdentifier
}

func (element *Normal) GetIdentifier() string {
	return element.newIdentifier
}

func (element *Normal) CanBeParentOf(other Definer) bool {
	if element.childLegacyIdentifiers == nil {
		return false
	}

	return slices.Contains(element.childLegacyIdentifiers, other.getLegacyIdentifier())
}

func (element *Normal) GetAttributeNames() []string {
	return element.attributeNames
}

func (element *Normal) canBeParent() bool {
	if element.childLegacyIdentifiers == nil {
		return false
	}

	return len(element.childLegacyIdentifiers) > 0
}

type Root struct {
	identifier string
}

func NewRoot(identifier string) *Root {
	return &Root{identifier: identifier}
}

func (element *Root) GetIdentifier() string {
	return element.identifier
}

func (element *Root) getLegacyIdentifier() string {
	return ""
}

func (element *Root) CanBeParentOf(_ Definer) bool {
	return true
}

func (element *Root) GetAttributeNames() []string {
	return nil
}

func (element *Root) canBeParent() bool {
	return true
}
