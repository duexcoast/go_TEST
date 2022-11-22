package main

import "reflect"

// x interface{} is the same as type any
func Walk(x interface{}, fn func(input string)) {
	// reflect.ValueOf returns us the value of a given variable
	val	:= reflect.ValueOf(x)

	// you can use NumField on a pointer value, so we have to use 
	// an if statement outside of the loop
	if val.Kind() == reflect.Pointer {
		// Elem() returns the value than an interface contains or a pointer points to
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Walk(field.Interface(), fn)
		}
	}
}