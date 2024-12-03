package kflow

import (
	"context"
	"kflow/kflow/utils"
)

type Flow[T any] struct {
	ctx context.Context

	data T
}

func (* Flow[T]) Run() error {
	panic("unimplement")
}

func ReportErrorInEnglish(){
	utils.SetLanguageIndex(2)
}

func ReportErrorInChinese(){
	utils.SetLanguageIndex(1)
}