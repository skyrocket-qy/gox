package redisx

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	JobStatusCompleted = "completed"
)

// ExecuteExactlyOnce ensures a job is performed only once across all services.
// It uses a permanent status key to check for prior completion and a temporary
// lock to prevent concurrent execution.
// It returns true if the job was executed by this call, false otherwise.
func ExecuteExactlyOnce(ctx context.Context, rdb *redis.Client, baseKey string, lockTTL time.Duration, job func() error) (bool, error) {
	statusKey := fmt.Sprintf("job:status:%s", baseKey)
	lockKey := fmt.Sprintf("job:lock:%s", baseKey)

	// 1. Check if the job is already marked as completed.
	// This is a fast path to exit if the work is already done.
	status, err := rdb.Get(ctx, statusKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return false, fmt.Errorf("failed to check job status: %w", err)
	}
	if status == JobStatusCompleted {
		log.Printf("Job '%s' was already completed. Skipping.", baseKey)
		return false, nil
	}

	// 2. Attempt to acquire the temporary lock.
	wasLockSet, err := rdb.SetNX(ctx, lockKey, 1, lockTTL).Result()
	if err != nil {
		return false, fmt.Errorf("failed to acquire lock: %w", err)
	}

	// Another instance is currently processing this job.
	if !wasLockSet {
		log.Printf("Job '%s' is currently being processed by another instance. Skipping.", baseKey)
		return false, nil
	}
	log.Printf("Lock acquired for job '%s'. Executing...", baseKey)

	// 3. Execute the actual job.
	jobErr := job()

	// 4. Mark as complete and release the lock atomically.
	// We use a Redis pipeline (MULTI/EXEC) to ensure these two operations
	// happen together without interruption.
	pipe := rdb.TxPipeline()
	if jobErr == nil {
		// If job was successful, set the permanent completion status.
		// We can give it a long TTL (e.g., 30 days) to eventually clean up.
		pipe.Set(ctx, statusKey, JobStatusCompleted, 30*24*time.Hour)
		log.Printf("Job '%s' completed successfully. Marking as done.", baseKey)
	} else {
		// If the job failed, we don't set the completion key, allowing a retry.
		// We still need to release the lock.
		log.Printf("Job '%s' failed. The lock will be released for a retry.", baseKey)
	}

	// Always release the lock, whether the job succeeded or failed.
	pipe.Del(ctx, lockKey)

	// Execute the pipeline.
	if _, err := pipe.Exec(ctx); err != nil {
		// This is a critical failure. The lock will eventually expire via TTL,
		// but the completion status may be inconsistent.
		return jobErr == nil, fmt.Errorf("CRITICAL: failed to execute final Redis pipeline for job '%s': %w", baseKey, err)
	}

	if jobErr != nil {
		return false, fmt.Errorf("job execution failed: %w", jobErr)
	}

	return true, nil
}
