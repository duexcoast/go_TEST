package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		}, {
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		}, {
			"struct with non-string field",
			Profile{24, "London"},
			[]string{"London"},
		}, {
			"nested fields",
			Person{
				"Chris",
				Profile{24, "London"},
			},
			[]string{"Chris", "London"},
		}, {
			"pointers to things",
			&Person{
				"Chris",
				Profile{33, "London"},
			},
			[]string{"Chris", "London"},
		}, {
			"slices",
			[]Profile{
				{33, "London"},
				{24, "Beirut"},
			},
			[]string{"London", "Beirut"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v expected %v", got, test.ExpectedCalls)
			}
		})
	}
}
