package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	/* cases as slices of struct */
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{"John"},
			ExpectedCalls: []string{"John"},
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{Name: "John", City: "London"},
			ExpectedCalls: []string{"John", "London"},
		},
		{
			Name: "struct with non string fild",
			Input: struct {
				Name string
				Age  int
			}{Name: "Doe", Age: 31},
			ExpectedCalls: []string{"Doe"},
		},
		{
			Name: "nested fields",
			Input: Person{
				Name:    "John",
				Profile: Profile{Age: 31, City: "London"},
			},
			ExpectedCalls: []string{"John", "London"},
		},
		{
			Name: "pointers to things",
			Input: &Person{
				Name:    "John",
				Profile: Profile{Age: 31, City: "London"},
			},
			ExpectedCalls: []string{"John", "London"},
		},
		{
			Name: "slices too",
			Input: []Profile{
				{Age: 31, City: "London"},
				{Age: 33, City: "Birmingham"},
			},
			ExpectedCalls: []string{"London", "Birmingham"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{Age: 31, City: "London"},
				{Age: 33, City: "Birmingham"},
			},
			ExpectedCalls: []string{"London", "Birmingham"},
		},
		{
			Name: "maps",
			Input: map[string]string{
				"Cow":   "Moo",
				"Sheep": "Beeek",
			},
			ExpectedCalls: []string{"Moo", "Beeek"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
