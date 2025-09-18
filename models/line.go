package models

import "DocumentConverter/element"

type Line struct {
	Element    element.Definer
	Attributes []Attribute
}
