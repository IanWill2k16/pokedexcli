package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "      hello    world      ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "      charizard            is               bae",
			expected: []string{"charizard", "is", "bae"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths do not match. Input: %q, expected length: %d, got: %d", actual, len(actual), len(c.expected))
			t.FailNow()
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words do not match. Input: %q, expected: %q, got: %q at index %d", c.input, c.expected[i], actual[i], i)
				t.FailNow()
			}
		}
	}
}
