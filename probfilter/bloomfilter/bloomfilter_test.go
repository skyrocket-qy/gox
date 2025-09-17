package bloomfilter

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	db, mock := redismock.NewClientMock()
	ctx := context.Background()
	key := "test_bloom"
	value := "test_value"

	mock.ExpectDo("BF.ADD", key, value).SetVal(int64(1))

	err := Add(ctx, db, key, value)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExists(t *testing.T) {
	db, mock := redismock.NewClientMock()
	ctx := context.Background()
	key := "test_bloom"
	value := "test_value"

	mock.ExpectDo("BF.EXISTS", key, value).SetVal(int64(1))

	exists, err := Exists(ctx, db, key, value)
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectDo("BF.EXISTS", key, "non_existent_value").SetVal(int64(0))
	exists, err = Exists(ctx, db, key, "non_existent_value")
	assert.NoError(t, err)
	assert.False(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExists_Error(t *testing.T) {
	db, mock := redismock.NewClientMock()
	ctx := context.Background()
	key := "test_bloom"
	value := "test_value"

	mock.ExpectDo("BF.EXISTS", key, value).SetErr(errors.New("redis error"))

	exists, err := Exists(ctx, db, key, value)
	assert.Error(t, err)
	assert.False(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExists_UnexpectedType(t *testing.T) {
	db, mock := redismock.NewClientMock()
	ctx := context.Background()
	key := "test_bloom"
	value := "test_value"

	mock.ExpectDo("BF.EXISTS", key, value).SetVal("not_an_int")

	exists, err := Exists(ctx, db, key, value)
	assert.Error(t, err)
	assert.False(t, exists)
	assert.Contains(t, err.Error(), "unexpected type for BF.EXISTS result")
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestReserve(t *testing.T) {
	db, mock := redismock.NewClientMock()
	ctx := context.Background()
	key := "test_bloom"
	errorRate := 0.01
	capacity := int64(1000)

	mock.ExpectDo("BF.RESERVE", key, errorRate, capacity).SetVal("OK")

	err := Reserve(ctx, db, key, errorRate, capacity)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestReserve_Error(t *testing.T) {
	db, mock := redismock.NewClientMock()
	ctx := context.Background()
	key := "test_bloom"
	errorRate := 0.01
	capacity := int64(1000)

	mock.ExpectDo("BF.RESERVE", key, errorRate, capacity).SetErr(errors.New("redis error"))

	err := Reserve(ctx, db, key, errorRate, capacity)
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
