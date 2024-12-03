package kflow_test

import (
	"context"
	"kflow/kflow"
	"testing"
)

type TheNode struct {
	Val   string `k:"type:auto;name:hey"`
	Value string `k:"type:input"`
}

func (t *TheNode) Name() string {
	return "the_node"
}

func (t *TheNode) Run() error {
	return nil
}

func NewTheNode(ctx context.Context) *TheNode {
	return &TheNode{}
}

func TestFlowBuilder(t *testing.T) {
	flow := kflow.NewFlowBuilder[int]()
	kflow.Add(flow, NewTheNode)
}
