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

func TestExecuteExactlyOnce_ContextCanceled(t *testing.T) {
	db, _ := redismock.NewClientMock()
	baseKey := "my-job-5"
	lockTTL := 1 * time.Hour
	job := func() error { return nil }

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := ExecuteExactlyOnce(ctx, db, baseKey, lockTTL, job)
	if !errors.Is(err, context.Canceled) {
		t.Errorf("Expected context.Canceled error, got %v", err)
	}
}

func TestExecuteExactlyOnce_RedisFailsOnSetNX(t *testing.T) {
	db, mock := redismock.NewClientMock()
	baseKey := "my-job-6"
	lockKey := "job:lock:" + baseKey
	lockTTL := 1 * time.Hour
	job := func() error { return nil }

	mock.ExpectSetNX(lockKey, 1, lockTTL).SetErr(errors.New("redis error"))
	mock.ExpectSetNX(lockKey, 1, lockTTL).SetVal(true)
	mock.ExpectGet("job:status:"+baseKey).RedisNil()
	mock.ExpectTxPipeline()
	mock.ExpectSet("job:status:"+baseKey, JobStatusCompleted, 30*24*time.Hour).SetVal("OK")
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

func TestExecuteExactlyOnce_RedisFailsOnGet(t *testing.T) {
	db, mock := redismock.NewClientMock()
	baseKey := "my-job-7"
	lockKey := "job:lock:" + baseKey
	statusKey := "job:status:" + baseKey
	lockTTL := 1 * time.Hour
	job := func() error { return nil }

	mock.ExpectSetNX(lockKey, 1, lockTTL).SetVal(true)
	mock.ExpectGet(statusKey).SetErr(errors.New("redis error"))
	mock.ExpectDel(lockKey).SetVal(1) // from defer

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

func TestExecuteExactlyOnce_RedisFailsOnTx(t *testing.T) {
	db, mock := redismock.NewClientMock()
	baseKey := "my-job-8"
	lockKey := "job:lock:" + baseKey
	statusKey := "job:status:" + baseKey
	lockTTL := 1 * time.Hour
	job := func() error { return nil }

	mock.ExpectSetNX(lockKey, 1, lockTTL).SetVal(true)
	mock.ExpectGet(statusKey).RedisNil()
	mock.ExpectTxPipeline()
	mock.ExpectSet(statusKey, JobStatusCompleted, 30*24*time.Hour).SetVal("OK")
	mock.ExpectDel(lockKey).SetVal(1)
	mock.ExpectTxPipelineExec().SetErr(errors.New("redis error"))
	mock.ExpectDel(lockKey).SetVal(1) // from defer

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

func TestRenewLock(t *testing.T) {
	db, mock := redismock.NewClientMock()
	key := "my-lock"
	ttl := 100 * time.Millisecond

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mock.ExpectPExpire(key, ttl).SetVal(true)
	mock.ExpectPExpire(key, ttl).SetVal(true)

	go renewLock(ctx, db, key, ttl)

	time.Sleep(250 * time.Millisecond) // Allow time for a few renewals

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
