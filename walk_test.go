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
		}, {
			"arrays",
			[2]Profile{
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
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Name": "Chris",
			"City": "London",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Chris")
		assertContains(t, got, "London")
	})
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{11, "Osaka"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Osaka"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("with functions", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{11, "Osaka"}
		}
		var got []string
		want := []string{"Berlin", "Osaka"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q, but it didn't", haystack, needle)
	}
}
