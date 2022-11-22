package main

import "reflect"

// x interface{} is the same as type any
func Walk(x interface{}, fn func(input string)) {
	// reflect.ValueOf returns us the value of a given variable
	val	:= reflect.ValueOf(x)
	// we look at the first field. if there weren't any fields, this would cause a panic
	field := val.Field(0)
	// this would be wrong if the field were something other than a string
	fn(field.String())
}