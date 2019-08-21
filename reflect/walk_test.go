package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
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
