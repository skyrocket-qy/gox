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

// ExecuteExactlyOnce tries to execute a job until it is successfully marked as complete,
// gracefully handling context cancellation. It returns an error if the context

// is canceled or if the job fails repeatedly.
func ExecuteExactlyOnce(
	ctx context.Context,
	rdb *redis.Client,
	baseKey string,
	lockTTL time.Duration,
	job func() error,
) error {
	jobRetryDelay := 1 * time.Second

	const (
		maxJobRetryDelay = 30 * time.Second
		lockPollInterval = 3 * time.Second
	)

	for {
		// Always check for context cancellation at the start of each attempt.
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// isFinished indicates the job is marked as completed in Redis.
		// jobErr contains errors from the job function itself.
		// attemptErr contains errors from Redis operations (e.g., connection issues).
		isFinished, jobErr, attemptErr := tryJobExecution(ctx, rdb, baseKey, lockTTL, job)

		// Case 1: The job is confirmed to be completed.
		if isFinished {
			return nil
		}

		// Case 2: A Redis or infrastructure error occurred.
		// We can retry this after a short delay.
		if attemptErr != nil {
			// In a real application, you would log this error.
			// log.Printf("Attempt failed with infrastructure error: %v. Retrying...", attemptErr)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(1 * time.Second): // Fixed short retry for Redis errors.
				continue
			}
		}

		// Case 3: The job function itself failed.
		// We retry with exponential backoff.
		if jobErr != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(jobRetryDelay):
			}

			jobRetryDelay *= 2
			if jobRetryDelay > maxJobRetryDelay {
				jobRetryDelay = maxJobRetryDelay
			}

			continue
		}

		// Case 4: The lock was held by another worker.
		// We wait for the polling interval before trying again.
		jobRetryDelay = 1 * time.Second // Reset job retry delay after a successful Redis op.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(lockPollInterval):
		}
	}
}

// tryJobExecution performs a single attempt to acquire the lock and run the job.
// It returns three values:
// 1. isFinished: true if the job status is "completed".
// 2. jobErr: An error returned by the job function itself.
// 3. attemptErr: An error from Redis operations.
func tryJobExecution(
	ctx context.Context,
	rdb *redis.Client,
	baseKey string,
	lockTTL time.Duration,
	job func() error,
) (isFinished bool, jobErr, attemptErr error) {
	statusKey := "job:status:" + baseKey
	lockKey := "job:lock:" + baseKey

	// Attempt to acquire the lock.
	wasLockSet, err := rdb.SetNX(ctx, lockKey, 1, lockTTL).Result()
	if err != nil {
		return false, nil, fmt.Errorf("failed to acquire lock: %w", err)
	}

	if !wasLockSet {
		// Lock is held by another worker. This is not an error.
		return false, nil, nil
	}

	// --- Lock Acquired ---
	// Use defer to guarantee the lock is released on all exit paths.
	defer rdb.Del(context.Background(), lockKey).Err() // Use background context for cleanup.

	// Double-check job status after acquiring the lock.
	// This handles the edge case where a worker completed the job but crashed before releasing the
	// lock.
	status, err := rdb.Get(ctx, statusKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return false, nil, fmt.Errorf("failed to check job status: %w", err)
	}

	if status == JobStatusCompleted {
		return true, nil, nil // Job already done.
	}

	// Start a background goroutine to renew the lock (heartbeat).
	// This prevents the lock from expiring during a long-running job.
	renewalCtx, cancelRenewal := context.WithCancel(ctx)
	defer cancelRenewal()

	go renewLock(renewalCtx, rdb, lockKey, lockTTL)

	// Execute the actual job.
	jobErr = job()
	if jobErr != nil {
		// Job failed, we will release the lock via defer and the caller can retry.
		return false, jobErr, nil
	}

	// --- Job Succeeded ---
	// Stop the lock renewal before the final atomic operation.
	cancelRenewal()

	// Atomically mark the job as complete and release the lock.
	// We release the lock here explicitly as part of the transaction,
	// so the deferred Del() will have no effect.
	pipe := rdb.TxPipeline()
	pipe.Set(ctx, statusKey, JobStatusCompleted, 30*24*time.Hour) // Set long-term status
	pipe.Del(ctx, lockKey)

	if _, err := pipe.Exec(ctx); err != nil {
		return false, nil, fmt.Errorf("CRITICAL: failed to set status and release lock: %w", err)
	}

	return true, nil, nil // Success!
}

// renewLock periodically extends the lock's TTL until the context is canceled.
func renewLock(ctx context.Context, rdb *redis.Client, key string, ttl time.Duration) {
	// Renew at a fraction of the TTL to ensure it doesn't expire.
	ticker := time.NewTicker(ttl / 2)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// PEXPIRE is used to update the TTL without changing the value.
			// We ignore the result, as the defer in the caller will handle cleanup
			// if the key expires for any reason.
			rdb.PExpire(ctx, key, ttl)
		case <-ctx.Done():
			return
		}
	}
}
