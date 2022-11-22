package main

import (
	"reflect"
)

// x interface{} is the same as type any
// we are looking for strings, when we find them we call fn (arg 2)
// if we have data types that contain strings within, we deal with them recursively
func Walk(x interface{}, fn func(input string)) {

	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field

	case reflect.Slice:
		numberOfValues = val.Len()
		getField = val.Index
	}
	for i := 0; i < numberOfValues; i++ {
		Walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// you can use NumField on a pointer, so we have to use
	// an if statement outside of the loop
	if val.Kind() == reflect.Pointer {
		// Elem() returns the value than an interface contains or a pointer points to
		val = val.Elem()
	}
	return val
}
