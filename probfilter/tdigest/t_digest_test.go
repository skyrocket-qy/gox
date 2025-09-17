package tdigest

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.CREATE", "mykey", 100.0).SetVal("OK")

	err := Create(ctx, db, "mykey", 100.0)
	assert.NoError(t, err)

	mock.ExpectDo("TDIGEST.CREATE", "mykey", 100.0).SetErr(errors.New("create failed"))

	err = Create(ctx, db, "mykey", 100.0)
	assert.Error(t, err)
}

func TestAdd(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	items := map[float64]int64{1.0: 10, 2.0: 20}
	mock.ExpectDo("TDIGEST.ADD", "mykey", 1.0, int64(10), 2.0, int64(20)).SetVal("OK")

	err := Add(ctx, db, "mykey", items)
	assert.NoError(t, err)

	mock.ExpectDo("TDIGEST.ADD", "mykey", 1.0, int64(10), 2.0, int64(20)).
		SetErr(errors.New("add failed"))

	err = Add(ctx, db, "mykey", items)
	assert.Error(t, err)
}

func TestMerge(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.MERGE", "destkey", "s1", "s2").SetVal("OK")

	err := Merge(ctx, db, "destkey", "s1", "s2")
	assert.NoError(t, err)

	mock.ExpectDo("TDIGEST.MERGE", "destkey", "s1", "s2").SetErr(errors.New("merge failed"))

	err = Merge(ctx, db, "destkey", "s1", "s2")
	assert.Error(t, err)
}

func TestMin(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.MIN", "mykey").SetVal(1.23)

	val, err := Min(ctx, db, "mykey")
	assert.NoError(t, err)
	assert.Equal(t, 1.23, val)

	mock.ExpectDo("TDIGEST.MIN", "mykey").SetErr(errors.New("min failed"))

	_, err = Min(ctx, db, "mykey")
	assert.Error(t, err)

	mock.ExpectDo("TDIGEST.MIN", "mykey").SetVal("not a float")

	_, err = Min(ctx, db, "mykey")
	assert.Error(t, err)
}

func TestMax(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.MAX", "mykey").SetVal(4.56)

	val, err := Max(ctx, db, "mykey")
	assert.NoError(t, err)
	assert.Equal(t, 4.56, val)

	mock.ExpectDo("TDIGEST.MAX", "mykey").SetErr(errors.New("max failed"))

	_, err = Max(ctx, db, "mykey")
	assert.Error(t, err)

	mock.ExpectDo("TDIGEST.MAX", "mykey").SetVal("not a float")

	_, err = Max(ctx, db, "mykey")
	assert.Error(t, err)
}

func TestQuantile(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.QUANTILE", "mykey", 0.5).SetVal(7.89)

	val, err := Quantile(ctx, db, "mykey", 0.5)
	assert.NoError(t, err)
	assert.Equal(t, 7.89, val)

	mock.ExpectDo("TDIGEST.QUANTILE", "mykey", 0.5).SetErr(errors.New("quantile failed"))

	_, err = Quantile(ctx, db, "mykey", 0.5)
	assert.Error(t, err)

	mock.ExpectDo("TDIGEST.QUANTILE", "mykey", 0.5).SetVal("not a float")

	_, err = Quantile(ctx, db, "mykey", 0.5)
	assert.Error(t, err)
}

func TestCDF(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.CDF", "mykey", 10.0).SetVal(0.99)

	val, err := CDF(ctx, db, "mykey", 10.0)
	assert.NoError(t, err)
	assert.Equal(t, 0.99, val)

	mock.ExpectDo("TDIGEST.CDF", "mykey", 10.0).SetErr(errors.New("cdf failed"))

	_, err = CDF(ctx, db, "mykey", 10.0)
	assert.Error(t, err)

	mock.ExpectDo("TDIGEST.CDF", "mykey", 10.0).SetVal("not a float")

	_, err = CDF(ctx, db, "mykey", 10.0)
	assert.Error(t, err)
}
