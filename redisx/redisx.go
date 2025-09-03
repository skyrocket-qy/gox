package redisx

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

// ExecuteOnce ensures that the provided job function is executed only once
// across all services for a given unique jobKey.
// It returns true if the job was executed by this call, false otherwise.
func ExecuteOnce(ctx context.Context, rdb *redis.Client, jobKey string, lockTTL time.Duration, job func() error) (bool, error) {
	// Attempt to acquire the lock.
	// SETNX is atomic. It returns true if the key was set (we got the lock).
	// The value "1" is arbitrary; its presence is what matters.
	wasSet, err := rdb.SetNX(ctx, jobKey, 1, lockTTL).Result()
	if err != nil {
		return false, fmt.Errorf("failed to acquire lock: %w", err)
	}

	// If wasSet is false, another instance has the lock.
	// We simply log and exit gracefully.
	if !wasSet {
		log.Printf("Job '%s' is already locked by another instance. Skipping.", jobKey)
		return false, nil
	}

	// We acquired the lock. Defer the release of the lock to ensure it runs
	// even if the job panics.
	log.Printf("Lock acquired for job '%s'. Executing...", jobKey)
	defer func() {
		if err := rdb.Del(ctx, jobKey).Err(); err != nil {
			log.Printf("CRITICAL: failed to release lock for key '%s': %v", jobKey, err)
		} else {
			log.Printf("Lock released for job '%s'.", jobKey)
		}
	}()

	// Execute the actual job.
	if err := job(); err != nil {
		return true, fmt.Errorf("job execution failed: %w", err)
	}

	return true, nil
}
