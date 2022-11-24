package main

import (
	"reflect"
)

// x interface{} is the same as type any
// we are looking for strings, when we find them we call fn (arg 2)
// if we have data types that contain strings within, we deal with them recursively
func Walk(x interface{}, fn func(input string)) {

	val := getValue(x)

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
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
