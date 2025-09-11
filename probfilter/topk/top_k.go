package topk

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Reserve creates a new TopK sketch.
// k: The number of top items to keep.
// width: The total number of counters to maintain in the sketch.
// decay: The decay factor for the counters. Items that are not seen for a while will have their counts reduced.
// period: The number of additions between decay operations.
func Reserve(ctx context.Context, rdb *redis.Client, key string, k, width, decay, period int64) error {
	return rdb.Do(ctx, "TOPK.RESERVE", key, k, width, decay, period).Err()
}

// Add adds items to the TopK sketch.
func Add(ctx context.Context, rdb *redis.Client, key string, items ...string) ([]string, error) {
	args := []interface{}{"TOPK.ADD", key}
	for _, item := range items {
		args = append(args, item)
	}
	res, err := rdb.Do(ctx, args...).Result()
	if err != nil {
		return nil, err
	}

	if arr, ok := res.([]interface{}); ok {
		removed := make([]string, len(arr))
		for i, v := range arr {
			if s, ok := v.(string); ok {
				removed[i] = s
			} else {
				return nil, fmt.Errorf("unexpected type for TOPK.ADD result at index %d: %T", i, v)
			}
		}
		return removed, nil
	}

	return nil, fmt.Errorf("unexpected type for TOPK.ADD result: %T", res)
}

// Query checks if items are among the top K items.
func Query(ctx context.Context, rdb *redis.Client, key string, items ...string) ([]bool, error) {
	args := []interface{}{"TOPK.QUERY", key}
	for _, item := range items {
		args = append(args, item)
	}
	res, err := rdb.Do(ctx, args...).Result()
	if err != nil {
		return nil, err
	}

	if arr, ok := res.([]interface{}); ok {
		results := make([]bool, len(arr))
		for i, v := range arr {
			if val, ok := v.(int64); ok {
				results[i] = val == 1
			} else {
				return nil, fmt.Errorf("unexpected type for TOPK.QUERY result at index %d: %T", i, v)
			}
		}
		return results, nil
	}

	return nil, fmt.Errorf("unexpected type for TOPK.QUERY result: %T", res)
}

// Count returns the count of specified items.
func Count(ctx context.Context, rdb *redis.Client, key string, items ...string) ([]int64, error) {
	args := []interface{}{"TOPK.COUNT", key}
	for _, item := range items {
		args = append(args, item)
	}
	res, err := rdb.Do(ctx, args...).Result()
	if err != nil {
		return nil, err
	}

	if arr, ok := res.([]interface{}); ok {
		counts := make([]int64, len(arr))
		for i, v := range arr {
			if count, ok := v.(int64); ok {
				counts[i] = count
			} else {
				return nil, fmt.Errorf("unexpected type for TOPK.COUNT result at index %d: %T", i, v)
			}
		}
		return counts, nil
	}

	return nil, fmt.Errorf("unexpected type for TOPK.COUNT result: %T", res)
}

// List returns the current top K items.
func List(ctx context.Context, rdb *redis.Client, key string) ([]string, error) {
	res, err := rdb.Do(ctx, "TOPK.LIST", key).Result()
	if err != nil {
		return nil, err
	}

	if arr, ok := res.([]interface{}); ok {
		items := make([]string, len(arr))
		for i, v := range arr {
			if s, ok := v.(string); ok {
				items[i] = s
			} else {
				return nil, fmt.Errorf("unexpected type for TOPK.LIST result at index %d: %T", i, v)
			}
		}
		return items, nil
	}

	return nil, fmt.Errorf("unexpected type for TOPK.LIST result: %T", res)
}

// Info returns information about the TopK sketch.
func Info(ctx context.Context, rdb *redis.Client, key string) (map[string]interface{}, error) {
	res, err := rdb.Do(ctx, "TOPK.INFO", key).Result()
	if err != nil {
		return nil, err
	}

	if arr, ok := res.([]interface{}); ok && len(arr) == 8 {
		info := make(map[string]interface{})
		info["k"] = arr[1]
		info["width"] = arr[3]
		info["decay"] = arr[5]
		info["period"] = arr[7]
		return info, nil
	}

	return nil, fmt.Errorf("unexpected type or length for TOPK.INFO result: %T", res)
}
