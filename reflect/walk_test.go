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
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			Name:     "string test",
			Input:    "Peter",
			Expected: []string{"Peter"},
		},

		{
			Name: "array test",
			Input: [2]Profile{
				{21, "Melbourn"},
				{22, "Sydney"},
			},
			Expected: []string{"Melbourn", "Sydney"},
		},

		{
			Name: "map test",
			Input: map[string]string{
				"China":  "Beijing",
				"Russia": "Moscow",
			},
			Expected: []string{"Beijing", "Moscow"},
		},

		{
			Name: "struct test case with only one string field",
			Input: struct {
				Name string
			}{"Chris"},
			Expected: []string{"Chris"},
		},

		{
			Name: "struct test case with two string fields",
			Input: struct {
				Name1 string
				Age   int
			}{Name1: "Jack", Age: 23},
			Expected: []string{"Jack"},
		},

		{
			Name: "Struct test case with nested fields",
			Input: Person{
				Name: "Mike",
				Profile: Profile{
					Age:  24,
					City: "L.A.",
				},
			},
			Expected: []string{"Mike", "L.A."},
		},

		{
			Name: "Pointer Struct test case",
			Input: &Person{
				Name: "Miller",
				Profile: Profile{
					Age:  25,
					City: "Paris",
				},
			},
			Expected: []string{"Miller", "Paris"},
		},

		{
			Name: "Slice test case, actually a slice of struct",
			Input: []Profile{
				{26, "London"},
				{27, "Boston"},
			},
			Expected: []string{"London", "Boston"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.Expected) {
				t.Errorf("got %v want %v", got, test.Expected)
			}
		})
	}
}
