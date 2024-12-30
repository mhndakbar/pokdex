package main

import (
	"testing"
)

// cleanInput method is expected to return slice of strings
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " hello ",
			expected: []string{"hello"},
		},
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Hello World ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length does not match")
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("actual: %s, expected: %s", actual[i], c.expected[i])
			}
		}
	}
}
