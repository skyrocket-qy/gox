package redisx

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// BloomFilterAdd adds an item to the bloom filter.
func BloomFilterAdd(ctx context.Context, rdb *redis.Client, key, value string) error {
	return rdb.Do(ctx, "BF.ADD", key, value).Err()
}

// BloomFilterExists checks if an item exists in the bloom filter.
func BloomFilterExists(ctx context.Context, rdb *redis.Client, key, value string) (bool, error) {
	res, err := rdb.Do(ctx, "BF.EXISTS", key, value).Result()
	if err != nil {
		return false, err
	}

	return res.(int64) == 1, nil
}

// BloomFilterReserve reserves a bloom filter with a specific error rate and capacity.
// The errorRate is the desired probability of false positives. A lower value means
// fewer false positives, but requires more memory. The capacity is the expected
// number of items that will be added to the filter.
// Choose the limit of the data capacity.
func BloomFilterReserve(
	ctx context.Context,
	rdb *redis.Client,
	key string,
	errorRate float64,
	capacity int64,
) error {
	return rdb.Do(ctx, "BF.RESERVE", key, errorRate, capacity).Err()
}
