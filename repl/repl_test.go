package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "  ",
			expected: []string{},
		},
	}

	for _, tst := range cases {
		res := CleanCommand(tst.input)
		if len(res) != len(tst.expected) {
			t.Errorf("Test Error %s, %v", res, tst.expected)
		}
	}
}
