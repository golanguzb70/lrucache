package lrucache

import (
	"fmt"
	"testing"
	"time"
)

func TestLRUCache(t *testing.T) {
	cache := New[string, string](10, 0)

	cache.Put("key1", "value1")
	cache.Put("key2", "value2")
	cache.Put("key3", "value3")

	// Test Get existing key
	value, found := cache.Get("key1")
	if !found || value != "value1" {
		t.Errorf("Expected 'value1', got '%v'", value)
	}

	// Test Get non-existing key
	_, found = cache.Get("key4")
	if found {
		t.Errorf("Expected 'false', got 'true'")
	}

	// Test Clear
	cache.Clear()
	_, found = cache.Get("key1")
	if found {
		t.Errorf("Expected 'false', got 'true'")
	}

	// Test timeout
	cache = New[string, string](12, 1) // 1-second timeout

	cache.Put("key5", "value5")
	fmt.Println("Test is here", time.Now())
	time.Sleep(2 * time.Second)
	fmt.Println("Test is here", time.Now())
	_, found = cache.Get("key5")
	if found {
		t.Errorf("Expected 'false', got 'true'")
	}
}
