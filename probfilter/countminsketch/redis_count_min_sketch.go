package countminsketch

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// InitByDim initializes a Count-Min Sketch with specified dimensions.
// width: The number of counters in each row.
// depth: The number of rows (hash functions).
func InitByDim(ctx context.Context, rdb *redis.Client, key string, width, depth int64) error {
	return rdb.Do(ctx, "CMS.INITBYDIM", key, width, depth).Err()
}

// InitByProb initializes a Count-Min Sketch with specified error rate and probability.
// errorRate: The desired error rate (epsilon).
// probability: The desired probability of error (delta).
func InitByProb(ctx context.Context, rdb *redis.Client, key string, errorRate, probability float64) error {
	return rdb.Do(ctx, "CMS.INITBYPROB", key, errorRate, probability).Err()
}

// IncrBy increments the count of an item by a specified increment.
func IncrBy(ctx context.Context, rdb *redis.Client, key, item string, increment int64) ([]int64, error) {
	res, err := rdb.Do(ctx, "CMS.INCRBY", key, item, increment).Result()
	if err != nil {
		return nil, err
	}

	if arr, ok := res.([]interface{}); ok {
		counts := make([]int64, len(arr))
		for i, v := range arr {
			if count, ok := v.(int64); ok {
				counts[i] = count
			} else {
				return nil, fmt.Errorf("unexpected type for increment result at index %d: %T", i, v)
			}
		}
		return counts, nil
	}

	return nil, fmt.Errorf("unexpected type for CMS.INCRBY result: %T", res)
}

// Query returns the estimated count of an item.
func Query(ctx context.Context, rdb *redis.Client, key, item string) (int64, error) {
	res, err := rdb.Do(ctx, "CMS.QUERY", key, item).Result()
	if err != nil {
		return 0, err
	}

	if arr, ok := res.([]interface{}); ok && len(arr) > 0 {
		if count, ok := arr[0].(int64); ok {
			return count, nil
		}
		return 0, fmt.Errorf("unexpected type for CMS.QUERY result element: %T", arr[0])
	}

	return 0, fmt.Errorf("unexpected type or empty result for CMS.QUERY: %T", res)
}

// Merge merges multiple Count-Min Sketches into a destination sketch.
func Merge(ctx context.Context, rdb *redis.Client, destKey string, sourceKeys []string) error {
	args := []interface{}{"CMS.MERGE", destKey, len(sourceKeys)}
	for _, k := range sourceKeys {
		args = append(args, k)
	}
	return rdb.Do(ctx, args...).Err()
}
