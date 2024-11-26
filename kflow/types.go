package kflow

type INode interface {
	Name() string	
}

type IMiddleware interface {
	Before() error
	After() error
}

type IRunnable interface {
	Run() error
}

type ILoader interface {
	Load() error
}