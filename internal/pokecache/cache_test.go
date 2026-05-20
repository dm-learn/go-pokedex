package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheAddGet(t *testing.T) {
	testInterval := 5 * time.Second

	tests := map[string]struct {
		key string
		val []byte
	} {
		"base": {key: "https://example.com", val: []byte("testdata")},
		"path": {key: "https://example.com/path", val: []byte("moretestdata")},
	}

	for name, c := range tests {
		t.Run(
			fmt.Sprintf("Running case %s", name),
			func (t *testing.T) {
				cache := NewCache(testInterval)
				cache.Add(c.key, c.val)
				gotVal, ok := cache.Get(c.key)
				if !ok {
					t.Errorf("Expected to find key: %v", c.key)
					return
				}
				if string(gotVal) != string(c.val) {
					t.Errorf(
						"Expected to find value %v. Found: %v",
						string(c.val),
						string(gotVal),
					)
					return
				}
			},
		)
	}
}

// func TestCacheReapLoop(t *testing.T) {

// }
