package pokecache

import (
	"bytes"
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test_NewCache(t *testing.T) {
	expectedDuration := time.Second * 5

	res := NewCache(expectedDuration)

	if res.entries == nil {
		t.Fatalf("Expected entries map to be initialized but it was nil\n")
	}

	if res.interval != expectedDuration {
		t.Fatalf("Expected interval to be: %v but got: %v\n", expectedDuration, res.interval)
	}

	if len(res.entries) != 0 {
		t.Fatalf("Expected entries map to be empty but found %d entries\n", len(res.entries))
	}
}

func Test_Add_KeyExists(t *testing.T) {
	expectedDuration := time.Second * 5
	expectedBytes := []byte{}
	expectedKey := "www.google.com"

	sut := NewCache(expectedDuration)

	sut.Add(expectedKey, expectedBytes)

	if sut.entries == nil {
		t.Fatalf("Map was nil\n")
	}

	if sut.interval != expectedDuration {
		t.Fatalf("Expected interval to be: %v but got: %v\n", expectedDuration, sut.interval)
	}

	if sutVal, exists := sut.entries[expectedKey]; !exists {
		t.Fatalf("Expected key: %v but not found\n", expectedKey)
	} else {
		if !bytes.Equal(expectedBytes, sutVal.val) {
			t.Fatalf("Expected value: '%v' but got: '%v'\n", string(expectedBytes), string(sutVal.val))
		}
	}

	sut.Add(expectedKey, expectedBytes)

	if sutVal, exists := sut.entries[expectedKey]; !exists {
		t.Fatalf("Expected key: %v to still exist after re-adding but it did not\n", expectedKey)
	} else {
		if !bytes.Equal(expectedBytes, sutVal.val) {
			t.Fatalf("After re-adding, expected value: '%v' but got: '%v'\n", string(expectedBytes), string(sutVal.val))
		}
	}
}

func Test_Add_ValueExists(t *testing.T) {
	expectedDuration := time.Second * 5
	expectedBytes := []byte("hallo dies sollten bytes sein")
	expectedKey := "www.google.com"

	sut := NewCache(expectedDuration)
	sut.Add(expectedKey, expectedBytes)

	if sut.entries == nil {
		t.Fatalf("Map was nil\n")
	}

	if sut.interval != expectedDuration {
		t.Fatalf("Expected interval to be: %v but got: %v\n", expectedDuration, sut.interval)
	}

	sutVal, exists := sut.entries[expectedKey]
	if !exists {
		t.Fatalf("Expected key: %v but not found\n", expectedKey)
	}

	if !bytes.Equal(expectedBytes, sutVal.val) {
		t.Fatalf("Expected value: '%v' but got: '%v'\n", string(expectedBytes), string(sutVal.val))
	}
}

func Test_Add_CreatedAtValidity(t *testing.T) {
	expectedDuration := time.Second * 5
	expectedBytes := []byte{}
	expectedKey := "www.google.com"

	sut := NewCache(expectedDuration)
	sut.Add(expectedKey, expectedBytes)

	if sut.entries == nil {
		t.Fatalf("Map was nil\n")
	}

	if sut.interval != expectedDuration {
		t.Fatalf("Expected interval to be: %v but got: %v\n", expectedDuration, sut.interval)
	}

	sutVal, exists := sut.entries[expectedKey]
	if !exists {
		t.Fatalf("Expected key: %v but not found\n", expectedKey)
	}

	now := time.Now()
	if sutVal.createdAt.After(now) || sutVal.createdAt.Before(now.Add(-expectedDuration)) {
		t.Fatalf("Expected `createdAt` to be within the last %v seconds, but got: %v\n", expectedDuration.Seconds(), sutVal.createdAt)
	}

	if !bytes.Equal(expectedBytes, sutVal.val) {
		t.Fatalf("Expected value: '%v' but got: '%v'\n", string(expectedBytes), string(sutVal.val))
	}
}

func Test_Add_ValueAndCreatedAtExists(t *testing.T) {
	expectedDuration := time.Second * 5
	expectedBytes := []byte("hallo dies sollten bytes sein")
	expectedKey := "www.google.com"

	sut := NewCache(expectedDuration)
	sut.Add(expectedKey, expectedBytes)

	if sut.entries == nil {
		t.Fatalf("Map was nil\n")
	}

	if sut.interval != expectedDuration {
		t.Fatalf("Expected interval to be: %v but got: %v\n", expectedDuration, sut.interval)
	}

	sutVal, exists := sut.entries[expectedKey]
	if !exists {
		t.Fatalf("Expected key: %v but not found\n", expectedKey)
	}

	now := time.Now()
	if sutVal.createdAt.After(now) || sutVal.createdAt.Before(now.Add(-expectedDuration)) {
		t.Fatalf("Expected `createdAt` to be recent, but got: %v\n", sutVal.createdAt)
	}

	if !bytes.Equal(expectedBytes, sutVal.val) {
		t.Fatalf("Expected value: '%v' but got: '%v'\n", string(expectedBytes), string(sutVal.val))
	}
}

func Test_Get_ValueAndCreatedAtExists(t *testing.T) {
	expectedDuration := time.Second * 5
	expectedBytes := []byte("hallo dies sollten bytes sein")
	expectedKey := "www.google.com"

	sut := NewCache(expectedDuration)
	sut.Add(expectedKey, expectedBytes)

	actual, exists := sut.Get(expectedKey)
	if !exists {
		t.Fatalf("Expected key: %v to exist but it did not\n", expectedKey)
	}

	if !bytes.Equal(expectedBytes, actual) {
		t.Fatalf("Expected value: '%v' but got: '%v'\n", string(expectedBytes), string(actual))
	}
}

func Test_ReapLoop_ValueAndCreatedAtExists(t *testing.T) {
	expectedDuration := time.Millisecond * 10
	expectedBytes := []byte("===expected==bytes===")
	expectedBytes2 := []byte("===expected==bytes===2")
	expectedKey := "www.google.com"
	expectedKey2 := "www.google-2.com"

	sut := NewCache(expectedDuration)
	sut.Add(expectedKey, expectedBytes)
	sut.Add(expectedKey2, expectedBytes2)

	time.Sleep(expectedDuration * 2)

	sut.ReapLoop()

	if sut.entries == nil {
		t.Fatalf("Map was nil\n")
	}

	if len(sut.entries) != 0 {
		t.Fatalf("Expected len of values to be 0 but got: %v\n", len(sut.entries))
	}

	if _, exists := sut.entries[expectedKey]; exists {
		t.Fatalf("Expected key '%v' to be deleted, but it was found\n", expectedKey)
	}
	if _, exists := sut.entries[expectedKey2]; exists {
		t.Fatalf("Expected key '%v' to be deleted, but it was found\n", expectedKey2)
	}
}

func Test_Add_ConcurrentAccess(t *testing.T) {
	expectedDuration := time.Microsecond * 10
	cache := NewCache(expectedDuration)
	numEntries := 100

	var wg sync.WaitGroup

	for i := 0; i < numEntries; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			time.Sleep(time.Microsecond)
			cache.Add(fmt.Sprintf("key_%d", i), []byte(fmt.Sprintf("vaulue_%d", i)))
		}(i)
	}

	wg.Wait()
	if len(cache.entries) != 100 {
		t.Fatalf("Expected length: %v but it was %v\n", numEntries, len(cache.entries))
	}
}

func Test_AddAndReset_ConcurrentAccess(t *testing.T) {
	expectedDuration := time.Microsecond * 10
	cache := NewCache(expectedDuration)
	numEntries := 100
	var wg sync.WaitGroup

	for i := 0; i < numEntries; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Microsecond)
			cache.Add(fmt.Sprintf("key_%d", i), []byte(fmt.Sprintf("value_%d", i)))
		}(i)
	}

	wg.Wait()

	if len(cache.entries) != numEntries {
		t.Fatalf("Expected length: %v but it was %v\n", numEntries, len(cache.entries))
	}

	cache.Reset()
	if len(cache.entries) != 0 {
		t.Fatalf("Expected length: 0 but it was %v\n", len(cache.entries))
	}
}

func Test_ReapLoop_ConcurrentAccess(t *testing.T) {
	expectedDuration := time.Millisecond * 10
	cache := NewCache(expectedDuration)

	var wg sync.WaitGroup

	numEntries := 100

	for i := 0; i < numEntries; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			cache.Add(fmt.Sprintf("key_%d", i), []byte("value"))
		}(i)
	}

	time.Sleep(time.Millisecond * 20)
	go func() {
		for j := 0; j < 5; j++ {
			cache.ReapLoop()
			time.Sleep(time.Millisecond)
		}
	}()

	wg.Wait()
	time.Sleep(time.Millisecond * 20)

	for key, value := range cache.entries {
		newTime := value.createdAt.Add(cache.interval)
		if newTime.Before(time.Now()) {
			t.Fatalf("Expected key: %v to be removed but it was found\n", key)
		}
	}
}
