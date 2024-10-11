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
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "heLLo wOrLd",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		expected := cs.expected

		if len(actual) != len(expected) {
			t.Errorf("The lengths are not equal -> actual: %v vs expected: %v", len(actual), len(expected))
			continue
		}
		for i := range actual {
			if actual[i] != expected[i] {
				t.Errorf("Words do not macth -> actual: %v vs expected: %v", actual[i], expected[i])
				continue
			}
		}

	}

}
