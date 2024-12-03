package kflow

import (
	"context"
	"fmt"
	"kflow/kflow/impl"
	"kflow/kflow/utils"
	"reflect"
	"strings"
)

type FlowBuilder[T any] struct {
	config *Config

	nodes map[impl.Name]*impl.Node
}

func NewFlowBuilder[T any]() *FlowBuilder[T] {
	return &FlowBuilder[T]{
		nodes: make(map[impl.Name]*impl.Node),
	}
}

func Add[T INode, R any](f *FlowBuilder[R], constructor impl.NodeConstructor[T]) {
	var output T
	nodeType, ptrDepth := utils.DeferenceToNonePtr(reflect.TypeOf(output))
	constructorWrapper := func(ctx context.Context) INode {
		return constructor(ctx)
	}
	f.addNodeConstructor(nodeType, ptrDepth, constructorWrapper)
}

func (f *FlowBuilder[T]) SetConfig(config *Config) *FlowBuilder[T] {
	f.config = config
	return f
}

func (f *FlowBuilder[T]) Prepare() error {

	return nil
}

func (f *FlowBuilder[T]) Build(ctx context.Context, data T) *Flow[T] {
	return &Flow[T]{
		ctx:  ctx,
		data: data,
	}
}

func (f *FlowBuilder[T]) addNodeConstructor(nodeType reflect.Type, ptrDepth int, constructor impl.NodeConstructorWrapper) {
	// parse node to extract the information for later use.
	parsedNode := f.parseNodeConstructor(nodeType, constructor)
	parsedNode.IsPtr = ptrDepth != 0
	parsedNode.PtrDepth = ptrDepth
	f.nodes[parsedNode.Name] = parsedNode

}

func (f *FlowBuilder[T]) parseNodeConstructor(nodeType reflect.Type, constructor impl.NodeConstructorWrapper) *impl.Node {
	nodeName := impl.Name(nodeType.Name())
	fields := make([]*impl.Field, 0, nodeType.NumField())
	for i := 0; i < nodeType.NumField(); i++ {
		fieldType := nodeType.Field(i)
		if _, ok := fieldType.Tag.Lookup("k"); !ok {
			continue
		}
		field := f.parseField(string(nodeName), fieldType)
		fields = append(fields, field)
	}

	node := &impl.Node{
		Name:           nodeName,
		Type:           nodeType,
		Fields:         fields,
		Constructor:    constructor,
		DependentNodes: []impl.DependInfo{},
	}

	return node
}

func (f *FlowBuilder[T]) parseField(nodeName string, fieldType reflect.StructField) *impl.Field {
	fieldName := fieldType.Name

	if !fieldType.IsExported() {
		panic(fmt.Sprintf("%v.%v is not a public field", nodeName, fieldName))
	}

	tagString := fieldType.Tag.Get("k")

	tag := f.parseTag(nodeName, fieldName, tagString)

	field := &impl.Field{
		Name:      impl.Name(fieldName),
		Type:      fieldType.Type,
		TagString: tagString,
		Tag:       tag,
	}
	return field
}

// parseTag parses the tag input impl.Tag. For example, the tag is "type:auto; name:hello; "
func (f *FlowBuilder[T]) parseTag(nodeName, fieldName, value string) *impl.Tag {
	t := new(impl.Tag)
	items := strings.Split(value, ";")
	allTags := make(map[string]string)
	for _, item := range items {
		parts := strings.SplitN(item, ":", 2)
		if len(parts) != 2 {
			panic(fmt.Sprintf("field (%v) in node (%v) has invalid tag (%v)", fieldName, nodeName, value))
		}
		allTags[strings.Trim(parts[0], " ")] = strings.Trim(parts[1], " ")
	}

	tagType := allTags["type"]
	if tagType == "" {
		panic(fmt.Sprintf("%v.%v misses type in tag", nodeName, fieldName))
	}
	switch tagType {
	case "auto":
		f.fillTagForAuto(t, allTags)
	case "input":
		f.fillTagForInput(t, allTags)
	case "delay":
		f.fillTagForDelay(t, allTags)
	default:
		panic(fmt.Sprintf("%v.%v has unsupported type %v in tag", nodeName, fieldName, tagType))
	}

	return t
}

func (f *FlowBuilder[T]) fillTagForAuto(t *impl.Tag, allTags map[string]string) {
	t.FieldType = impl.Auto
	t.Name = impl.Name(allTags["name"])
}

func (f *FlowBuilder[T]) fillTagForInput(t *impl.Tag, allTags map[string]string) {
	t.FieldType = impl.Input
	t.Name = impl.Name(allTags["name"])
}

func (f *FlowBuilder[T]) fillTagForDelay(t *impl.Tag, allTags map[string]string) {
	t.FieldType = impl.Delay
	t.Name = impl.Name(allTags["name"])
}
