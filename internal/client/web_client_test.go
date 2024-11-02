package client

import (
	"testing"
)

func Test_EmptyString(t *testing.T) {
	expected := []string{}

	res := RequestLocations("")
	t.Fatalf("Expected array lenght: %v but got: %v", len(expected), res)
}
