package utils

import "reflect"

func DeferenceToNonePtr(t reflect.Type) (reflect.Type, int) {
	if t.Kind() != reflect.Pointer {
		return t, 0
	}
	i := 0
	for t.Kind() == reflect.Pointer {
		i++
		if i == 10 {
			panic("pointer is too deep")
		}
		t = t.Elem()
	}
	return t, i
}
