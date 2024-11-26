package kflow

import "context"

type Flow[T any] struct {
	ctx context.Context

	data T
}

func (* Flow[T]) Run() error {
	panic("unimplement")
}