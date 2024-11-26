package kflow

import "reflect"

type NodeField struct {
	name string
	fieldType reflect.Type
	tagString string
	
}

type Node struct{
	name string
	nodeType reflect.Type
	fields map[string]NodeField
}