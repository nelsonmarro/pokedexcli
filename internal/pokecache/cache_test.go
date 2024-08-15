package pokecache

import (
	"fmt"
	"testing"
	"time"
)

// test add to cache and get from cache
func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://testapi.com",
			val: []byte("test data"),
		},
		{
			key: "https://testapi.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			// Arrange
			cache := NewCache(interval)

			// Act
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)

			// Assert
			if !ok {
				t.Error("Expected to find the key")
				return
			}

			if string(val) != string(c.val) {
				t.Error("Expected to find a correct value")
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const basetime = 5 * time.Microsecond
	const waitTime = basetime + 5*time.Microsecond
	cache := NewCache(basetime)
	cache.Add("https://testapi.com", []byte("test data"))

	_, ok := cache.Get("https://testapi.com")
	if !ok {
		t.Error("Expected to find the key")
		return
	}

	time.Sleep(waitTime)
	_, ok = cache.Get("https://testapi.com")
	if ok {
		t.Error("Expected not to find a key")
		return
	}
}
