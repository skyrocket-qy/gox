package bloomfilter_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/skyrocket-qy/gox/probfilter/bloomfilter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	db, mock := redismock.NewClientMock()
	ctx := context.Background()
	key := "test_bloom"
	value := "test_value"

	mock.ExpectDo("BF.ADD", key, value).SetVal(int64(1))

	err := bloomfilter.Add(ctx, db, key, value) // Added bloomfilter.
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExists(t *testing.T) {
	db, mock := redismock.NewClientMock()
	ctx := context.Background()
	key := "test_bloom"
	value := "test_value"

	mock.ExpectDo("BF.EXISTS", key, value).SetVal(int64(1))

	exists, err := bloomfilter.Exists(ctx, db, key, value) // Added bloomfilter.
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectDo("BF.EXISTS", key, "non_existent_value").SetVal(int64(0))
	exists, err = bloomfilter.Exists(ctx, db, key, "non_existent_value") // Added bloomfilter.
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

	exists, err := bloomfilter.Exists(ctx, db, key, value) // Added bloomfilter.
	require.Error(t, err)
	assert.False(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExists_UnexpectedType(t *testing.T) {
	db, mock := redismock.NewClientMock()
	ctx := context.Background()
	key := "test_bloom"
	value := "test_value"

	mock.ExpectDo("BF.EXISTS", key, value).SetVal("not_an_int")

	exists, err := bloomfilter.Exists(ctx, db, key, value) // Already had bloomfilter.
	require.Error(t, err)
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

	err := bloomfilter.Reserve(ctx, db, key, errorRate, capacity) // Added bloomfilter.
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

	err := bloomfilter.Reserve(ctx, db, key, errorRate, capacity) // Added bloomfilter.
	require.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
