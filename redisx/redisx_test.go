package redisx

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
)

func TestExecuteExactlyOnce_SuccessFirstTry(t *testing.T) {
	db, mock := redismock.NewClientMock()
	baseKey := "my-job"
	lockKey := "job:lock:" + baseKey
	statusKey := "job:status:" + baseKey
	lockTTL := 1 * time.Hour

	job := func() error { return nil }

	mock.ExpectSetNX(lockKey, 1, lockTTL).SetVal(true)
	mock.ExpectGet(statusKey).RedisNil()
	mock.ExpectTxPipeline()
	mock.ExpectSet(statusKey, JobStatusCompleted, 30*24*time.Hour).SetVal("OK")
	mock.ExpectDel(lockKey).SetVal(1)
	mock.ExpectTxPipelineExec()
	mock.ExpectDel(lockKey).SetVal(0)

	err := ExecuteExactlyOnce(context.Background(), db, baseKey, lockTTL, job)
	if err != nil {
		t.Errorf("ExecuteExactlyOnce returned an error: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestExecuteExactlyOnce_JobFailsThenSucceeds(t *testing.T) {
	db, mock := redismock.NewClientMock()
	baseKey := "my-job-2"
	lockKey := "job:lock:" + baseKey
	statusKey := "job:status:" + baseKey
	lockTTL := 1 * time.Hour

	var jobExecutionCount int
	job := func() error {
		jobExecutionCount++
		if jobExecutionCount == 1 {
			return errors.New("job failed")
		}
		return nil
	}

	// First attempt (job fails)
	mock.ExpectSetNX(lockKey, 1, lockTTL).SetVal(true)
	mock.ExpectGet(statusKey).RedisNil()
	mock.ExpectDel(lockKey).SetVal(1)

	// Second attempt (job succeeds)
	mock.ExpectSetNX(lockKey, 1, lockTTL).SetVal(true)
	mock.ExpectGet(statusKey).RedisNil()
	mock.ExpectTxPipeline()
	mock.ExpectSet(statusKey, JobStatusCompleted, 30*24*time.Hour).SetVal("OK")
	mock.ExpectDel(lockKey).SetVal(1)
	mock.ExpectTxPipelineExec()
	mock.ExpectDel(lockKey).SetVal(0)

	err := ExecuteExactlyOnce(context.Background(), db, baseKey, lockTTL, job)
	if err != nil {
		t.Errorf("ExecuteExactlyOnce returned an error: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestExecuteExactlyOnce_LockHeld(t *testing.T) {
	db, mock := redismock.NewClientMock()
	baseKey := "my-job-3"
	lockKey := "job:lock:" + baseKey
	statusKey := "job:status:" + baseKey
	lockTTL := 1 * time.Hour

	job := func() error { return nil }

	// First attempt (lock is held)
	mock.ExpectSetNX(lockKey, 1, lockTTL).SetVal(false)

	// Second attempt (lock is acquired and job succeeds)
	mock.ExpectSetNX(lockKey, 1, lockTTL).SetVal(true)
	mock.ExpectGet(statusKey).RedisNil()
	mock.ExpectTxPipeline()
	mock.ExpectSet(statusKey, JobStatusCompleted, 30*24*time.Hour).SetVal("OK")
	mock.ExpectDel(lockKey).SetVal(1)
	mock.ExpectTxPipelineExec()
	mock.ExpectDel(lockKey).SetVal(0)

	err := ExecuteExactlyOnce(context.Background(), db, baseKey, lockTTL, job)
	if err != nil {
		t.Errorf("ExecuteExactlyOnce returned an error: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestExecuteExactlyOnce_JobAlreadyCompleted(t *testing.T) {
	db, mock := redismock.NewClientMock()
	baseKey := "my-job-4"
	lockKey := "job:lock:" + baseKey
	statusKey := "job:status:" + baseKey
	lockTTL := 1 * time.Hour

	job := func() error {
		t.Error("job should not be executed")
		return nil
	}

	mock.ExpectSetNX(lockKey, 1, lockTTL).SetVal(true)
	mock.ExpectGet(statusKey).SetVal(JobStatusCompleted)
	mock.ExpectDel(lockKey).SetVal(1)

	err := ExecuteExactlyOnce(context.Background(), db, baseKey, lockTTL, job)
	if err != nil {
		t.Errorf("ExecuteExactlyOnce returned an error: %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
