package cuckoofilter

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// CuckooFilterReserve reserves a cuckoo filter with a specific capacity.
func Reserve(ctx context.Context, rdb *redis.Client, key string, capacity int64) error {
	return rdb.Do(ctx, "CF.RESERVE", key, capacity).Err()
}

// CuckooFilterAdd adds an item to the cuckoo filter.
func Add(ctx context.Context, rdb *redis.Client, key, value string) (bool, error) {
	res, err := rdb.Do(ctx, "CF.ADD", key, value).Result()
	if err != nil {
		return false, err
	}

	if val, ok := res.(bool); ok {
		return val, nil
	}

	return false, fmt.Errorf("unexpected type for result: %T", res)
}

// CuckooFilterAddNX adds an item to the cuckoo filter if it does not exist.
func AddNX(ctx context.Context, rdb *redis.Client, key, value string) (bool, error) {
	res, err := rdb.Do(ctx, "CF.ADDNX", key, value).Result()
	if err != nil {
		return false, err
	}

	if val, ok := res.(bool); ok {
		return val, nil
	}

	return false, fmt.Errorf("unexpected type for result: %T", res)
}

// CuckooFilterExists checks if an item exists in the cuckoo filter.
func Exists(ctx context.Context, rdb *redis.Client, key, value string) (bool, error) {
	res, err := rdb.Do(ctx, "CF.EXISTS", key, value).Result()
	if err != nil {
		return false, err
	}

	if val, ok := res.(bool); ok {
		return val, nil
	}

	return false, fmt.Errorf("unexpected type for result: %T", res)
}

// CuckooFilterDel deletes an item from the cuckoo filter.
func Del(ctx context.Context, rdb *redis.Client, key, value string) (bool, error) {
	res, err := rdb.Do(ctx, "CF.DEL", key, value).Result()
	if err != nil {
		return false, err
	}

	if val, ok := res.(bool); ok {
		return val, nil
	}

	return false, fmt.Errorf("unexpected type for result: %T", res)
}

// CuckooFilterCount returns the number of items in a cuckoo filter.
func Count(ctx context.Context, rdb *redis.Client, key, value string) (int64, error) {
	res, err := rdb.Do(ctx, "CF.COUNT", key, value).Result()
	if err != nil {
		return 0, err
	}

	if val, ok := res.(int64); ok {
		return val, nil
	}

	return 0, fmt.Errorf("unexpected type for result: %T", res)
}
