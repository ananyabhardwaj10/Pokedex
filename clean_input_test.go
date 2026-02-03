package main
import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	} {
		{
			input: " Hello World ",
			expected: []string{"hello", "world"},
		},

		{
			input: "This is my project work ",
			expected: []string{"this", "is", "my", "project", "work"},
		},

		{
			input: "  Pokemon is my favorite cartoon     ",
			expected: []string{"pokemon", "is", "my", "favorite", "cartoon"},
		},

		{
			input: "WE ARE BUILDING POKEDEX",
			expected: []string{"we", "are", "building", "pokedex"},
		},
	}

	for _, c := range cases {
		actual:= cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Expected length of slice is not equal to actual length of slice. Test Failed")
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Mismatch in actual words and expected words. Test Failed")
			}
		}
	}
}