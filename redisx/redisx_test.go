package redisx

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestBloomFilter(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	key := "test_bloom_filter"
	value := "test_value"

	// Clean up before test
	rdb.Del(ctx, key)

	// 1. Reserve the filter
	err := BloomFilterReserve(ctx, rdb, key, 0.01, 1000)
	if err != nil {
		t.Fatalf("BloomFilterReserve failed: %v", err)
	}

	// 2. Add a value to the filter
	err = BloomFilterAdd(ctx, rdb, key, value)
	if err != nil {
		t.Fatalf("BloomFilterAdd failed: %v", err)
	}

	// 3. Check if the value exists
	exists, err := BloomFilterExists(ctx, rdb, key, value)
	if err != nil {
		t.Fatalf("BloomFilterExists failed: %v", err)
	}
	if !exists {
		t.Errorf("BloomFilterExists should return true for a value that was added")
	}

	// 4. Check for a non-existent value
	exists, err = BloomFilterExists(ctx, rdb, key, "non_existent_value")
	if err != nil {
		t.Fatalf("BloomFilterExists failed for non-existent value: %v", err)
	}
	if exists {
		t.Errorf("BloomFilterExists should return false for a value that was not added")
	}

	// Clean up after test
	rdb.Del(ctx, key)
}
