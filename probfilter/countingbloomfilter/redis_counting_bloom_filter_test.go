package countingbloomfilter

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestReserve(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("BF.RESERVE", "mykey", 0.01, int64(1000)).SetVal("OK")

	err := Reserve(ctx, db, "mykey", 0.01, 1000)
	assert.NoError(t, err)

	mock.ExpectDo("BF.RESERVE", "mykey", 0.01, int64(1000)).SetErr(errors.New("reserve failed"))

	err = Reserve(ctx, db, "mykey", 0.01, 1000)
	assert.Error(t, err)
}

func TestAdd(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("BF.ADD", "mykey", "item1").SetVal("OK")

	err := Add(ctx, db, "mykey", "item1")
	assert.NoError(t, err)

	mock.ExpectDo("BF.ADD", "mykey", "item1").SetErr(errors.New("add failed"))

	err = Add(ctx, db, "mykey", "item1")
	assert.Error(t, err)
}

func TestExists(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("BF.EXISTS", "mykey", "item1").SetVal(int64(1))

	exists, err := Exists(ctx, db, "mykey", "item1")
	assert.NoError(t, err)
	assert.True(t, exists)

	mock.ExpectDo("BF.EXISTS", "mykey", "item2").SetVal(int64(0))

	exists, err = Exists(ctx, db, "mykey", "item2")
	assert.NoError(t, err)
	assert.False(t, exists)

	mock.ExpectDo("BF.EXISTS", "mykey", "item3").SetErr(errors.New("exists failed"))

	_, err = Exists(ctx, db, "mykey", "item3")
	assert.Error(t, err)

	mock.ExpectDo("BF.EXISTS", "mykey", "item4").SetVal("not an int")

	_, err = Exists(ctx, db, "mykey", "item4")
	assert.Error(t, err)
}

func TestRemove(t *testing.T) {
	ctx := context.Background()
	db, _ := redismock.NewClientMock()

	removed, err := Remove(ctx, db, "mykey", "item1")
	assert.Error(t, err)
	assert.False(t, removed)
	assert.Equal(
		t,
		"BF.DEL command is not supported for Bloom Filters. Use Cuckoo Filters for direct deletion.",
		err.Error(),
	)
}
