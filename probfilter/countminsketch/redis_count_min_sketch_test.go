package countminsketch

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestInitByDim(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("CMS.INITBYDIM", "mykey", int64(100), int64(5)).SetVal("OK")

	err := InitByDim(ctx, db, "mykey", 100, 5)
	assert.NoError(t, err)

	mock.ExpectDo("CMS.INITBYDIM", "mykey", int64(100), int64(5)).SetErr(errors.New("init failed"))

	err = InitByDim(ctx, db, "mykey", 100, 5)
	assert.Error(t, err)
}

func TestInitByProb(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("CMS.INITBYPROB", "mykey", 0.01, 0.001).SetVal("OK")

	err := InitByProb(ctx, db, "mykey", 0.01, 0.001)
	assert.NoError(t, err)

	mock.ExpectDo("CMS.INITBYPROB", "mykey", 0.01, 0.001).SetErr(errors.New("init failed"))

	err = InitByProb(ctx, db, "mykey", 0.01, 0.001)
	assert.Error(t, err)
}

func TestIncrBy(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("CMS.INCRBY", "mykey", "item1", int64(10)).SetVal([]any{int64(10)})

	counts, err := IncrBy(ctx, db, "mykey", "item1", 10)
	assert.NoError(t, err)
	assert.Equal(t, []int64{10}, counts)

	mock.ExpectDo("CMS.INCRBY", "mykey", "item1", int64(10)).SetErr(errors.New("incrby failed"))

	_, err = IncrBy(ctx, db, "mykey", "item1", 10)
	assert.Error(t, err)

	mock.ExpectDo("CMS.INCRBY", "mykey", "item1", int64(10)).SetVal("not an array")

	_, err = IncrBy(ctx, db, "mykey", "item1", 10)
	assert.Error(t, err)

	mock.ExpectDo("CMS.INCRBY", "mykey", "item1", int64(10)).SetVal([]any{"not an int"})

	_, err = IncrBy(ctx, db, "mykey", "item1", 10)
	assert.Error(t, err)
}

func TestQuery(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("CMS.QUERY", "mykey", "item1").SetVal([]any{int64(10)})

	count, err := Query(ctx, db, "mykey", "item1")
	assert.NoError(t, err)
	assert.Equal(t, int64(10), count)

	mock.ExpectDo("CMS.QUERY", "mykey", "item1").SetErr(errors.New("query failed"))

	_, err = Query(ctx, db, "mykey", "item1")
	assert.Error(t, err)

	mock.ExpectDo("CMS.QUERY", "mykey", "item1").SetVal("not an array")

	_, err = Query(ctx, db, "mykey", "item1")
	assert.Error(t, err)

	mock.ExpectDo("CMS.QUERY", "mykey", "item1").SetVal([]any{})

	_, err = Query(ctx, db, "mykey", "item1")
	assert.Error(t, err)

	mock.ExpectDo("CMS.QUERY", "mykey", "item1").SetVal([]any{"not an int"})

	_, err = Query(ctx, db, "mykey", "item1")
	assert.Error(t, err)
}

func TestMerge(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	sourceKeys := []string{"s1", "s2"}
	mock.ExpectDo("CMS.MERGE", "destkey", len(sourceKeys), "s1", "s2").SetVal("OK")
	err := Merge(ctx, db, "destkey", sourceKeys)
	assert.NoError(t, err)

	mock.ExpectDo("CMS.MERGE", "destkey", len(sourceKeys), "s1", "s2").
		SetErr(errors.New("merge failed"))
	err = Merge(ctx, db, "destkey", sourceKeys)
	assert.Error(t, err)
}
