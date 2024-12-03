package impl

import (
	"reflect"
)

type Tag struct {
	FieldType FieldType
	Name      Name // by default, the name is empty and the field name is used to reference the resource.
}

type DependInfo struct {
	NodeName  Name
	FieldName Name
}

type Field struct {
	Name      Name         // the name of the field
	Type      reflect.Type // the type of the field
	TagString string       // the original string of the tag
	Tag       *Tag         // the parsed version of the tag
}

type Node struct {
	Name           Name                   // the name of the node without path
	Type           reflect.Type           // the type of the node
	IsPtr          bool                   // if the type is of ptr
	PtrDepth       int                    // the depth from the ptr to the underline data
	Fields         []*Field               // all the fields in the node
	Constructor    NodeConstructorWrapper // the constructor to generate a new node
	DependentNodes []DependInfo           // the directly dependent nodes
}
