package topk

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

	mock.ExpectDo("TOPK.RESERVE", "mykey", int64(10), int64(100), int64(99), int64(1000)).SetVal("OK")
	err := Reserve(ctx, db, "mykey", 10, 100, 99, 1000)
	assert.NoError(t, err)

	mock.ExpectDo("TOPK.RESERVE", "mykey", int64(10), int64(100), int64(99), int64(1000)).SetErr(errors.New("reserve failed"))
	err = Reserve(ctx, db, "mykey", 10, 100, 99, 1000)
	assert.Error(t, err)
}

func TestAdd(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TOPK.ADD", "mykey", "item1", "item2").SetVal([]any{"item3"})
	removed, err := Add(ctx, db, "mykey", "item1", "item2")
	assert.NoError(t, err)
	assert.Equal(t, []string{"item3"}, removed)
}

func TestQuery(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TOPK.QUERY", "mykey", "item1", "item2").SetVal([]any{int64(1), int64(0)})

	exists, err := Query(ctx, db, "mykey", "item1", "item2")
	assert.NoError(t, err)
	assert.Equal(t, []bool{true, false}, exists)
}

func TestCount(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TOPK.COUNT", "mykey", "item1", "item2").SetVal([]any{int64(5), int64(3)})
	counts, err := Count(ctx, db, "mykey", "item1", "item2")
	assert.NoError(t, err)
	assert.Equal(t, []int64{5, 3}, counts)
}

func TestList(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TOPK.LIST", "mykey").SetVal([]any{"item1", "item2"})
	items, err := List(ctx, db, "mykey")
	assert.NoError(t, err)
	assert.Equal(t, []string{"item1", "item2"}, items)
}

func TestInfo(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	infoResult := []any{
		"k", int64(10),
		"width", int64(100),
		"decay", int64(99),
		"period", int64(1000),
	}
	mock.ExpectDo("TOPK.INFO", "mykey").SetVal(infoResult)
	info, err := Info(ctx, db, "mykey")
	assert.NoError(t, err)
	assert.Equal(t, int64(10), info["k"])
	assert.Equal(t, int64(100), info["width"])
	assert.Equal(t, int64(99), info["decay"])
	assert.Equal(t, int64(1000), info["period"])
}
