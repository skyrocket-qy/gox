package hyperloglog

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Add adds all the element arguments to the HyperLogLog data structure.
func Add(ctx context.Context, rdb *redis.Client, key string, elements ...any) (int64, error) {
	args := []any{"PFADD", key}
	args = append(args, elements...)

	res, err := rdb.Do(ctx, args...).Result()
	if err != nil {
		return 0, err
	}

	if val, ok := res.(int64); ok {
		return val, nil
	}

	return 0, nil // Should not happen
}

// Count returns the approximated cardinality of the set observed by the HyperLogLog at key(s).
func Count(ctx context.Context, rdb *redis.Client, keys ...string) (int64, error) {
	args := []any{"PFCOUNT"}
	for _, key := range keys {
		args = append(args, key)
	}

	res, err := rdb.Do(ctx, args...).Result()
	if err != nil {
		return 0, err
	}

	if val, ok := res.(int64); ok {
		return val, nil
	}

	return 0, nil // Should not happen
}

// Merge merges multiple HyperLogLog values into a single value.
func Merge(ctx context.Context, rdb *redis.Client, destKey string, sourceKeys ...string) error {
	args := []any{"PFMERGE", destKey}
	for _, key := range sourceKeys {
		args = append(args, key)
	}

	return rdb.Do(ctx, args...).Err()
}
