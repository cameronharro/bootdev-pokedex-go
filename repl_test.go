package main

import (
	"testing"
)

func TestCleanUserInput(t *testing.T) {
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

func TestGetFirstWord(t *testing.T) {
	cases := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{},
			expected: "help",
		},
		{
			input:    []string{"map", "mapb"},
			expected: "map",
		},
	}

	for _, c := range cases {
		actual := getFirstWord(c.input)
		if c.expected != actual {
			t.Errorf("Failure: actual %s does not equal expected %v", actual, c.expected)
		}
	}
}
