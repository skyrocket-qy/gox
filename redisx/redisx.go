package redisx

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	JobStatusCompleted = "completed"
)

// Try attempts to execute a job until it is successfully marked as complete,
// gracefully handling context cancellation. It returns an error if the context
// is canceled before the job can be completed.
func Try(ctx context.Context, rdb *redis.Client, baseKey string, lockTTL time.Duration,
	job func() error,
) error {
	errorRetryDelay := 1 * time.Second
	const maxErrorRetryDelay = 30 * time.Second
	const lockPollInterval = 3 * time.Second
	const maxRetry = 5
	retryCnt := 0

	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		isFinished, err := ExecuteExactlyOnce(ctx, rdb, baseKey, lockTTL, job)

		if err != nil {
			if retryCnt >= maxRetry {
				return fmt.Errorf("maximum number of retries reached: %w", err)
			}

			retryCnt++

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(errorRetryDelay):
			}

			errorRetryDelay *= 2
			if errorRetryDelay > maxErrorRetryDelay {
				errorRetryDelay = maxErrorRetryDelay
			}
			continue
		}

		errorRetryDelay = 1 * time.Second

		if isFinished {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(lockPollInterval):
		}
	}
}

func ExecuteExactlyOnce(ctx context.Context, rdb *redis.Client, baseKey string, lockTTL time.Duration,
	job func() error) (isFinished bool, err error,
) {
	statusKey := fmt.Sprintf("job:status:%s", baseKey)
	lockKey := fmt.Sprintf("job:lock:%s", baseKey)

	// get lock
	wasLockSet, err := rdb.SetNX(ctx, lockKey, 1, lockTTL).Result()
	if err != nil {
		return false, fmt.Errorf("failed to acquire lock: %w", err)
	}
	if !wasLockSet {
		return false, nil
	}

	// check is finshied
	status, err := rdb.Get(ctx, statusKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return false, fmt.Errorf("failed to check job status: %w", err)
	}
	if status == JobStatusCompleted {
		return true, nil
	}

	jobErr := job()

	// Mark as complete and release the lock atomically.
	// We use a Redis pipeline (MULTI/EXEC) to ensure these two operations
	// happen together without interruption.
	pipe := rdb.TxPipeline()
	if jobErr == nil {
		pipe.Set(ctx, statusKey, JobStatusCompleted, 30*24*time.Hour)
	} else {
	}

	pipe.Del(ctx, lockKey)

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
