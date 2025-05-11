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
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello     world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello\tworld",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "       ",
			expected: []string{},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("returned %d words, expected %d", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("word %d: got %q, expected %q", i, actual[i], c.expected[i])
			}
		}
	}
}
