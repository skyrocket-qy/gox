package redisx

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// CuckooFilterReserve reserves a cuckoo filter with a specific capacity.
func CuckooFilterReserve(ctx context.Context, rdb *redis.Client, key string, capacity int64) error {
	return rdb.Do(ctx, "CF.RESERVE", key, capacity).Err()
}

// CuckooFilterAdd adds an item to the cuckoo filter.
func CuckooFilterAdd(ctx context.Context, rdb *redis.Client, key string, value string) (bool, error) {
	res, err := rdb.Do(ctx, "CF.ADD", key, value).Result()
	if err != nil {
		return false, err
	}
	return res.(bool), nil
}

// CuckooFilterAddNX adds an item to the cuckoo filter if it does not exist.
func CuckooFilterAddNX(ctx context.Context, rdb *redis.Client, key string, value string) (bool, error) {
	res, err := rdb.Do(ctx, "CF.ADDNX", key, value).Result()
	if err != nil {
		return false, err
	}
	return res.(bool), nil
}

// CuckooFilterExists checks if an item exists in the cuckoo filter.
func CuckooFilterExists(ctx context.Context, rdb *redis.Client, key string, value string) (bool, error) {
	res, err := rdb.Do(ctx, "CF.EXISTS", key, value).Result()
	if err != nil {
		return false, err
	}
	return res.(bool), nil
}

// CuckooFilterDel deletes an item from the cuckoo filter.
func CuckooFilterDel(ctx context.Context, rdb *redis.Client, key string, value string) (bool, error) {
	res, err := rdb.Do(ctx, "CF.DEL", key, value).Result()
	if err != nil {
		return false, err
	}
	return res.(bool), nil
}

// CuckooFilterCount returns the number of items in a cuckoo filter.
func CuckooFilterCount(ctx context.Context, rdb *redis.Client, key string, value string) (int64, error) {
	res, err := rdb.Do(ctx, "CF.COUNT", key, value).Result()
	if err != nil {
		return 0, err
	}
	return res.(int64), nil
}
