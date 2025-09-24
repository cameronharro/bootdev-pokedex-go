package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    " hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello World!",
			expected: []string{"hello", "world!"},
		},
	}

	for _, c := range cases {
		actual := cleanUserInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("len mismatch: actual %d, expected %d from input %#v", len(actual), len(c.expected), c.input)
			return
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]
			if word != expected {
				t.Errorf("Failure: actual %v does not equal expected %v", word, expected)
			}
		}
	}
}
