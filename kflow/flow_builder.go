package kflow

import "context"

type FlowBuilder[T any] struct {
	
}

func NewFlowBuilder[T any]() *FlowBuilder[T] {
	return &FlowBuilder[T]{}
}

func (f *FlowBuilder[T]) NewFlow(ctx context.Context, data T) *Flow[T] {
	return &Flow[T]{
		ctx:  ctx,
		data: data,
	}
}

func (f *FlowBuilder[T]) With(node INode) *FlowBuilder[T] {
	f.addNode(node)
	return f
}

func (f *FlowBuilder[T]) addNode(node INode) {
	panic("unimplemented")
}
