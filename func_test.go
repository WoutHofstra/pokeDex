package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "   hello        world   ",
			expected: []string{"hello", "world"},
		},
		{
			input: " i  am   a  teapot ",
			expected: []string{"i", "am", "a", "teapot"},
		},
		{
			input: "LowerCaSeS ALSO",
			expected: []string{"lowercases", "also"},
		},
	}


	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test failed! Expected: %v, got: %v", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Test failed! Expected: %v, got: %v", expectedWord, word)
			}
		}
	}
}
