//tests
package main

import (
	"testing"
)

func TestCleanInput (t *testing.T){
	cases := []struct {
		input	string
		expected []string
	}{
		{
			input:	"  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:	"world of words",
			expected: []string{"world", "of", "words"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len of slices does not match")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word %v does not match expected word %v", word, expectedWord)
			}
		}

	}
}