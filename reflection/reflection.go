package main

import (
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Map:
		for _, key := range val.MapKeys() {
			Walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.String:
		fn(val.String())
	}

	for i := 0; i < numberOfValues; i++ {
		Walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
