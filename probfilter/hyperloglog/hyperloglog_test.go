package hyperloglog

import (
	"context"
	"errors"
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

	mock.ExpectDo("PFADD", "mykey", "a", "b", "c").SetErr(errors.New("add failed"))
	_, err = Add(ctx, db, "mykey", "a", "b", "c")
	assert.Error(t, err)

	mock.ExpectDo("PFADD", "mykey", "a", "b", "c").SetVal("not an int")
	val, err = Add(ctx, db, "mykey", "a", "b", "c")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), val)
}

func TestCount(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("PFCOUNT", "key1", "key2").SetVal(int64(12345))
	count, err := Count(ctx, db, "key1", "key2")
	assert.NoError(t, err)
	assert.Equal(t, int64(12345), count)

	mock.ExpectDo("PFCOUNT", "key1", "key2").SetErr(errors.New("count failed"))
	_, err = Count(ctx, db, "key1", "key2")
	assert.Error(t, err)

	mock.ExpectDo("PFCOUNT", "key1", "key2").SetVal("not an int")
	count, err = Count(ctx, db, "key1", "key2")
	assert.NoError(t, err)
	assert.Equal(t, int64(0), count)
}

func TestMerge(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("PFMERGE", "dest", "key1", "key2").SetVal("OK")
	err := Merge(ctx, db, "dest", "key1", "key2")
	assert.NoError(t, err)

	mock.ExpectDo("PFMERGE", "dest", "key1", "key2").SetErr(errors.New("merge failed"))
	err = Merge(ctx, db, "dest", "key1", "key2")
	assert.Error(t, err)
}
