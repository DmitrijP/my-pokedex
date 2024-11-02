package pokecache

import (
	"testing"
	"time"
)

func Test_NewCache(t *testing.T) {
	expected := make(map[string]CacheEntry)
	expectedDuration := time.Second * 5
	res := NewCache(expectedDuration)
	if res.entries == nil {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), 4)
	}
	if res.interval != expectedDuration {
		t.Fatalf("Expected array lenght: %v but got: %v", len(expected), 4)
	}
	for key, value := range res.entries {
		if val, exists := expected[key]; !exists {
			t.Fatalf("Expected array lenght: %v but got: %v", len(expected), 4)
		} else {
			if val.createdAt != value.createdAt || string(val.val) != string(value.val) {
				t.Fatalf("Expected array lenght: %v but got: %v", len(expected), 4)
			}
		}
	}
}
