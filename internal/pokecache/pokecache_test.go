package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	if cache.cache == nil {
		t.Error("Cache is Nil")
	}

}

func TestAddGetCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Add("1", []byte("val1"))
	actual, ok := cache.Get("1")
	if !ok {
		t.Error("Key not found")
	}
	if string(actual) != string([]byte("val1")) {
		t.Error("Val don't match")
	}

}

func TestReapCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Add("1", []byte("val1"))

	time.Sleep(interval + time.Millisecond)
	_, ok := cache.Get("1")

	if ok {
		t.Error("ReapFailed")
	}

}
