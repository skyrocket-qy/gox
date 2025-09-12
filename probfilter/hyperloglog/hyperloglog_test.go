package hyperloglog

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("PFADD", "mykey", "a", "b", "c").SetVal(int64(1))
	val, err := Add(ctx, db, "mykey", "a", "b", "c")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), val)
}

func TestCount(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("PFCOUNT", "key1", "key2").SetVal(int64(12345))
	count, err := Count(ctx, db, "key1", "key2")
	assert.NoError(t, err)
	assert.Equal(t, int64(12345), count)
}

func TestMerge(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("PFMERGE", "dest", "key1", "key2").SetVal("OK")
	err := Merge(ctx, db, "dest", "key1", "key2")
	assert.NoError(t, err)
}
