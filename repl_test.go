package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world",
			expected: []string{"hello", "world"},
		},

		{
			input:    " hello   world",
			expected: []string{"hello", "world"},
		},

		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},

		{
			input:    "",
			expected: []string{},
		},

		{
			input:    "hello\tworld",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLen := len(actual)
		expectedLen := len(c.expected)

		if actualLen != expectedLen {
			t.Errorf("length mismatch: got %d, wanted %d", actualLen, expectedLen)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words don't match: got %v, wanted %v", word, expectedWord)
			}
		}
	}
}
