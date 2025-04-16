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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Bulbasaur CharManDer Squirtle",
			expected: []string{"bulbasaur", "charmander", "squirtle"},
		},
		{
			input:    "   Pikachu  EKANS  Mankey   ",
			expected: []string{"pikachu", "ekans", "mankey"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Input: %q\nExpected length: %d, got: %d", c.input, len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Input: %q\nWord %d mismatch: expected %q, got %q", c.input, i, expectedWord, word)
			}
		}
	}
}
