package countingbloomfilter

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Reserve reserves a Counting Bloom filter with a specific error rate and capacity.
// It uses the BF.RESERVE command, which can be used for both regular and counting bloom filters.
func Reserve(
	ctx context.Context,
	rdb *redis.Client,
	key string,
	errorRate float64,
	capacity int64,
) error {
	return rdb.Do(ctx, "BF.RESERVE", key, errorRate, capacity).Err()
}

// Add adds an item to the Counting Bloom filter.
// For a counting bloom filter, adding an item multiple times effectively increments its internal
// count.
func Add(ctx context.Context, rdb *redis.Client, key, value string) error {
	return rdb.Do(ctx, "BF.ADD", key, value).Err()
}

// Exists checks if an item exists in the Counting Bloom filter.
func Exists(ctx context.Context, rdb *redis.Client, key, value string) (bool, error) {
	res, err := rdb.Do(ctx, "BF.EXISTS", key, value).Result()
	if err != nil {
		return false, err
	}

	if val, ok := res.(int64); ok {
		return val == 1, nil
	}

	return false, fmt.Errorf("unexpected type for BF.EXISTS result: %T", res)
}

// Remove attempts to remove an item from the Counting Bloom filter.
// Note: RedisBloom's BF.DEL command is for Cuckoo Filters. For Bloom Filters, removal is not
// directly supported in a way that decrements counts. If true counting bloom filter behavior with
// decrements is needed, a different approach or data structure might be required (e.g., a custom
// implementation or a Count-Min Sketch). This function is included for completeness but its
// effectiveness for true counting bloom filter decrements is limited.
func Remove(ctx context.Context, rdb *redis.Client, key, value string) (bool, error) {
	// RedisBloom's BF.DEL is for Cuckoo Filters. There's no direct decrement for Bloom Filters.
	// If a true decrement is needed, a different approach is required.
	// For now, we'll return false as it's not directly supported by BF commands for Bloom Filters.
	return false, errors.New(
		"BF.DEL command is not supported for Bloom Filters. Use Cuckoo Filters for direct deletion.",
	)
}
