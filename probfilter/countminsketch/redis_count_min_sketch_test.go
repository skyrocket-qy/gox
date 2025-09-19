package countminsketch_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/skyrocket-qy/gox/probfilter/countminsketch"
	"github.com/stretchr/testify/assert"
)

func TestInitByDim(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("CMS.INITBYDIM", "mykey", int64(100), int64(5)).SetVal("OK")

	err := countminsketch.InitByDim(ctx, db, "mykey", 100, 5)
	assert.NoError(t, err)

	mock.ExpectDo("CMS.INITBYDIM", "mykey", int64(100), int64(5)).SetErr(errors.New("init failed"))

	err = countminsketch.InitByDim(ctx, db, "mykey", 100, 5)
	assert.Error(t, err)
}

func TestInitByProb(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("CMS.INITBYPROB", "mykey", 0.01, 0.001).SetVal("OK")

	err := countminsketch.InitByProb(ctx, db, "mykey", 0.01, 0.001)
	assert.NoError(t, err)

	mock.ExpectDo("CMS.INITBYPROB", "mykey", 0.01, 0.001).SetErr(errors.New("init failed"))

	err = countminsketch.InitByProb(ctx, db, "mykey", 0.01, 0.001)
	assert.Error(t, err)
}

func TestIncrBy(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("CMS.INCRBY", "mykey", "item1", int64(10)).SetVal([]any{int64(10)})

	counts, err := countminsketch.IncrBy(ctx, db, "mykey", "item1", 10)
	assert.NoError(t, err)
	assert.Equal(t, []int64{10}, counts)

	mock.ExpectDo("CMS.INCRBY", "mykey", "item1", int64(10)).SetErr(errors.New("incrby failed"))

	_, err = countminsketch.IncrBy(ctx, db, "mykey", "item1", 10)
	assert.Error(t, err)

	mock.ExpectDo("CMS.INCRBY", "mykey", "item1", int64(10)).SetVal("not an array")

	_, err = countminsketch.IncrBy(ctx, db, "mykey", "item1", 10)
	assert.Error(t, err)

	mock.ExpectDo("CMS.INCRBY", "mykey", "item1", int64(10)).SetVal([]any{"not an int"})

	_, err = countminsketch.IncrBy(ctx, db, "mykey", "item1", 10)
	assert.Error(t, err)
}

func TestQuery(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("CMS.QUERY", "mykey", "item1").SetVal([]any{int64(10)})

	count, err := countminsketch.Query(ctx, db, "mykey", "item1")
	assert.NoError(t, err)
	assert.Equal(t, int64(10), count)

	mock.ExpectDo("CMS.QUERY", "mykey", "item1").SetErr(errors.New("query failed"))

	_, err = countminsketch.Query(ctx, db, "mykey", "item1")
	assert.Error(t, err)

	mock.ExpectDo("CMS.QUERY", "mykey", "item1").SetVal("not an array")

	_, err = countminsketch.Query(ctx, db, "mykey", "item1")
	assert.Error(t, err)

	mock.ExpectDo("CMS.QUERY", "mykey", "item1").SetVal([]any{})

	_, err = countminsketch.Query(ctx, db, "mykey", "item1")
	assert.Error(t, err)

	mock.ExpectDo("CMS.QUERY", "mykey", "item1").SetVal([]any{"not an int"})

	_, err = countminsketch.Query(ctx, db, "mykey", "item1")
	assert.Error(t, err)
}

func TestMerge(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	sourceKeys := []string{"s1", "s2"}
	mock.ExpectDo("CMS.MERGE", "destkey", len(sourceKeys), "s1", "s2").SetVal("OK")
	err := countminsketch.Merge(ctx, db, "destkey", sourceKeys)
	assert.NoError(t, err)

	mock.ExpectDo("CMS.MERGE", "destkey", len(sourceKeys), "s1", "s2").
		SetErr(errors.New("merge failed"))
	err = countminsketch.Merge(ctx, db, "destkey", sourceKeys)
	assert.Error(t, err)
}
