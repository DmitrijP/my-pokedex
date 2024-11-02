package commands

import (
	"testing"
)

func Test_EmptyString(t *testing.T) {
	input := ""
	expected := []string{}

	res := CleanCommand(input)
	if len(res) != len(expected) {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), len(res))
	}
}

func Test_WhitespaceString(t *testing.T) {
	input := "  "
	expected := []string{}

	res := CleanCommand(input)
	if len(res) != len(expected) {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), len(res))
	}
}

func Test_SingleCommand(t *testing.T) {
	input := "help"
	expected := []string{"help"}

	res := CleanCommand(input)
	if len(res) != len(expected) {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), len(res))
	}
	for i, r := range res {
		if r != expected[i] {
			t.Fatalf("Expected %v but got %v", expected, res)
		}
	}
}

func Test_SingleCommandCaseInsensitive(t *testing.T) {
	input := "HeLp"
	expected := []string{"help"}

	res := CleanCommand(input)
	if len(res) != len(expected) {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), len(res))
	}
	for i, r := range res {
		if r != expected[i] {
			t.Fatalf("Expected %v but got %v", expected, res)
		}
	}
}

func Test_LeadingWhitespaceSingleCommand(t *testing.T) {
	input := "   help"
	expected := []string{"help"}

	res := CleanCommand(input)
	if len(res) != len(expected) {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), len(res))
	}
	for i, r := range res {
		if r != expected[i] {
			t.Fatalf("Expected %v but got %v", expected, res)
		}
	}
}

func Test_TrailingWhitespaceSingleCommand(t *testing.T) {
	input := "help    "
	expected := []string{"help"}

	res := CleanCommand(input)
	if len(res) != len(expected) {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), len(res))
	}
	for i, r := range res {
		if r != expected[i] {
			t.Fatalf("Expected %v but got %v", expected, res)
		}
	}
}

func Test_LeadingAndTrailingWhitespaceSingleCommand(t *testing.T) {
	input := "   help    "
	expected := []string{"help"}

	res := CleanCommand(input)
	if len(res) != len(expected) {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), len(res))
	}
	for i, r := range res {
		if r != expected[i] {
			t.Fatalf("Expected %v but got %v", expected, res)
		}
	}
}

func Test_MultipleCommand(t *testing.T) {
	input := "help test hello"
	expected := []string{"help", "test", "hello"}

	res := CleanCommand(input)
	if len(res) != len(expected) {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), len(res))
	}
	for i, r := range res {
		if r != expected[i] {
			t.Fatalf("Expected %v but got %v", expected, res)
		}
	}
}

func Test_LeadingWhitespaceMultipleCommand(t *testing.T) {
	input := "      help test hello"
	expected := []string{"help", "test", "hello"}

	res := CleanCommand(input)
	if len(res) != len(expected) {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), len(res))
	}
	for i, r := range res {
		if r != expected[i] {
			t.Fatalf("Expected %v but got %v", expected, res)
		}
	}
}

func Test_LeadingAndTrailingWhitespaceMultipleCommand(t *testing.T) {
	input := "      help test hello    "
	expected := []string{"help", "test", "hello"}

	res := CleanCommand(input)
	if len(res) != len(expected) {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), len(res))
	}
	for i, r := range res {
		if r != expected[i] {
			t.Fatalf("Expected %v but got %v", expected, res)
		}
	}
}

func Test_TooManyWhitespaceMultipleCommand(t *testing.T) {
	input := "      help     test      hello    "
	expected := []string{"help", "test", "hello"}

	res := CleanCommand(input)
	if len(res) != len(expected) {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), len(res))
	}
	for i, r := range res {
		if r != expected[i] {
			t.Fatalf("Expected %v but got %v", expected, res)
		}
	}
}
