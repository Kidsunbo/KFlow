package impl

import "context"

type Name string

type FieldType int

type NodeConstructor[T INode] func(ctx context.Context) T
type NodeConstructorWrapper func(ctx context.Context) INode

type INode interface {
	Name() string
	Run() error
}

type IMiddleware interface {
	Before() error
	After() error
}

const (
	// Auto field type means that the dependency will be injected automatically
	Auto FieldType = iota

	// Input field type means that the dependency should be given in the configuration files
	Input

	// Delay field type means that the dependency is injected later by a manually call
	Delay

	// Output field hold the result that one node produce
	Output
)
