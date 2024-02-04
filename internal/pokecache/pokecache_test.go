package pokecache

import (
	"testing"
	"time"
)

// run all tests with: go test ./...

func TestCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Errorf("cache.cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	cache := NewCache(time.Millisecond * 100)
	key := "the-key"
	value := []byte("the-value")

	cache.Add(key, value)
	actual, exists := cache.Get(key)

	if !exists {
		t.Errorf("key '%v' not found", key)
	}
	if string(actual) != string(value) {
		t.Errorf("value '%v' not found", value)
	}
}

func TestReap(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)
	key := "the-key"
	value := []byte("the-value")

	cache.Add(key, value)
	time.Sleep(time.Millisecond * 20)
	_, exists := cache.Get(key)

	if exists {
		t.Error("key should have been reaped")
	}
}

func TestNoReap(t *testing.T) {
	cache := NewCache(time.Millisecond * 50)
	key := "the-key"
	value := []byte("the-value")

	cache.Add(key, value)
	time.Sleep(time.Millisecond * 10)
	_, exists := cache.Get(key)

	if !exists {
		t.Error("key should not have been reaped")
	}
}
