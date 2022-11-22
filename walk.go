package main

import (
	"reflect"
)

// x interface{} is the same as type any
func Walk(x interface{}, fn func(input string)) {

	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			Walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
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
