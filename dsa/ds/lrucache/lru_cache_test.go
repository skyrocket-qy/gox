package lrucache_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/dsa/ds/lrucache"
)

func TestLRUCache(t *testing.T) {
	cache := lrucache.New(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	if cache.Get(1) != 1 {
		t.Errorf("Get(1) failed, expected 1, got %d", cache.Get(1))
	}

	cache.Put(3, 3)

	if cache.Get(2) != -1 {
		t.Errorf("Get(2) failed, expected -1, got %d", cache.Get(2))
	}

	cache.Put(4, 4)

	if cache.Get(1) != -1 {
		t.Errorf("Get(1) failed, expected -1, got %d", cache.Get(1))
	}

	if cache.Get(3) != 3 {
		t.Errorf("Get(3) failed, expected 3, got %d", cache.Get(3))
	}

	if cache.Get(4) != 4 {
		t.Errorf("Get(4) failed, expected 4, got %d", cache.Get(4))
	}
}

func TestLRUCache_update(t *testing.T) {
	cache := lrucache.New(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(1, 10)

	if cache.Get(1) != 10 {
		t.Errorf("Get(1) failed, expected 10, got %d", cache.Get(1))
	}
}

func TestLRUCache_zero_capacity(t *testing.T) {
	cache := lrucache.New(0)
	cache.Put(1, 1)

	if cache.Get(1) != -1 {
		t.Errorf("Get(1) failed, expected -1, got %d", cache.Get(1))
	}
}
