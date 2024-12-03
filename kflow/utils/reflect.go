package utils

import (
	"fmt"
	"reflect"
)

func DeferenceToNonePtr(t reflect.Type) (reflect.Type, int) {
	if t.Kind() != reflect.Pointer {
		return t, 0
	}
	i := 0
	for t.Kind() == reflect.Pointer {
		i++
		if i == 11 {
			panic(fmt.Sprintf(ErrorMessage(PointerTooDeep), 10))
		}
		t = t.Elem()
	}
	return t, i
}
