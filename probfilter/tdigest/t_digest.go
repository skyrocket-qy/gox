package tdigest

import (
	"context"
	"fmt"
	"sort"

	"github.com/redis/go-redis/v9"
)

// Create creates a new T-Digest sketch.
// compression: A parameter that controls the accuracy and size of the sketch. Higher values mean
// more accuracy and larger size.
func Create(ctx context.Context, rdb *redis.Client, key string, compression float64) error {
	return rdb.Do(ctx, "TDIGEST.CREATE", key, compression).Err()
}

// Add adds observations to the T-Digest sketch.
// items: A slice of item-weight pairs. Each item is a float64, and its weight is an int64.
func Add(ctx context.Context, rdb *redis.Client, key string, items map[float64]int64) error {
	args := []any{"TDIGEST.ADD", key}

	keys := make([]float64, 0, len(items))
	for k := range items {
		keys = append(keys, k)
	}

	sort.Float64s(keys)

	for _, k := range keys {
		args = append(args, k, items[k])
	}

	return rdb.Do(ctx, args...).Err()
}

// Merge merges multiple T-Digest sketches into a destination sketch.
func Merge(ctx context.Context, rdb *redis.Client, destKey string, sourceKeys ...string) error {
	args := []any{"TDIGEST.MERGE", destKey}
	for _, k := range sourceKeys {
		args = append(args, k)
	}

	return rdb.Do(ctx, args...).Err()
}

// Min returns the minimum value observed by the T-Digest sketch.
func Min(ctx context.Context, rdb *redis.Client, key string) (float64, error) {
	res, err := rdb.Do(ctx, "TDIGEST.MIN", key).Result()
	if err != nil {
		return 0, err
	}

	if val, ok := res.(float64); ok {
		return val, nil
	}

	return 0, fmt.Errorf("unexpected type for TDIGEST.MIN result: %T", res)
}

// Max returns the maximum value observed by the T-Digest sketch.
func Max(ctx context.Context, rdb *redis.Client, key string) (float64, error) {
	res, err := rdb.Do(ctx, "TDIGEST.MAX", key).Result()
	if err != nil {
		return 0, err
	}

	if val, ok := res.(float64); ok {
		return val, nil
	}

	return 0, fmt.Errorf("unexpected type for TDIGEST.MAX result: %T", res)
}

// Quantile returns the estimated value at a specific quantile.
// quantile: A float64 between 0 and 1.
func Quantile(
	ctx context.Context,
	rdb *redis.Client,
	key string,
	quantile float64,
) (float64, error) {
	res, err := rdb.Do(ctx, "TDIGEST.QUANTILE", key, quantile).Result()
	if err != nil {
		return 0, err
	}

	if val, ok := res.(float64); ok {
		return val, nil
	}

	return 0, fmt.Errorf("unexpected type for TDIGEST.QUANTILE result: %T", res)
}

// CDF returns the estimated cumulative distribution function (CDF) of a value.
// value: The value for which to estimate the CDF.
func CDF(ctx context.Context, rdb *redis.Client, key string, value float64) (float64, error) {
	res, err := rdb.Do(ctx, "TDIGEST.CDF", key, value).Result()
	if err != nil {
		return 0, err
	}

	if val, ok := res.(float64); ok {
		return val, nil
	}

	return 0, fmt.Errorf("unexpected type for TDIGEST.CDF result: %T", res)
}
