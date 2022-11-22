package main

import "reflect"

// x interface{} is the same as type any
func Walk(x interface{}, fn func(input string)) {
	// reflect.ValueOf returns us the value of a given variable
	val	:= reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String{
			fn(field.String())
		}
	}
}