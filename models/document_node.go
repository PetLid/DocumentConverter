package models

type DocumentNode struct {
	Name       string
	Attributes []Attribute
	Children   []DocumentNode
}
