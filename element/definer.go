package element

type Definer interface {
	GetIdentifier() string
	getLegacyIdentifier() string
	CanBeParentOf(element Definer) bool
	GetAttributeNames() []string
	canBeParent() bool
}
