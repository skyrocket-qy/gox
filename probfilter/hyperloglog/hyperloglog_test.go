package hyperloglog_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/skyrocket-qy/gox/probfilter/hyperloglog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("PFADD", "mykey", "a", "b", "c").SetVal(int64(1))

	val, err := hyperloglog.Add(ctx, db, "mykey", "a", "b", "c") // Added hyperloglog.
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)

	mock.ExpectDo("PFADD", "mykey", "a", "b", "c").SetErr(errors.New("add failed"))

	_, err = hyperloglog.Add(ctx, db, "mykey", "a", "b", "c") // Added hyperloglog.
	require.Error(t, err)

	mock.ExpectDo("PFADD", "mykey", "a", "b", "c").SetVal("not an int")

	val, err = hyperloglog.Add(ctx, db, "mykey", "a", "b", "c") // Added hyperloglog.
	assert.NoError(t, err)
	assert.Equal(t, int64(0), val)
}

func TestCount(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("PFCOUNT", "key1", "key2").SetVal(int64(12345))

	count, err := hyperloglog.Count(ctx, db, "key1", "key2") // Added hyperloglog.
	assert.NoError(t, err)
	assert.Equal(t, int64(12345), count)

	mock.ExpectDo("PFCOUNT", "key1", "key2").SetErr(errors.New("count failed"))

	_, err = hyperloglog.Count(ctx, db, "key1", "key2") // Added hyperloglog.
	require.Error(t, err)

	mock.ExpectDo("PFCOUNT", "key1", "key2").SetVal("not an int")

	count, err = hyperloglog.Count(ctx, db, "key1", "key2") // Already had hyperloglog.
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
}

func TestMerge(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("PFMERGE", "dest", "key1", "key2").SetVal("OK")

	err := hyperloglog.Merge(ctx, db, "dest", "key1", "key2") // Added hyperloglog.
	assert.NoError(t, err)

	mock.ExpectDo("PFMERGE", "dest", "key1", "key2").SetErr(errors.New("merge failed"))

	err = hyperloglog.Merge(ctx, db, "dest", "key1", "key2") // Added hyperloglog.
	require.Error(t, err)
}
