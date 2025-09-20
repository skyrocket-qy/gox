package tdigest_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/skyrocket-qy/gox/probfilter/tdigest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.CREATE", "mykey", 100.0).SetVal("OK")

	err := tdigest.Create(ctx, db, "mykey", 100.0)
	assert.NoError(t, err)

	mock.ExpectDo("TDIGEST.CREATE", "mykey", 100.0).SetErr(errors.New("create failed"))

	err = tdigest.Create(ctx, db, "mykey", 100.0)
	require.Error(t, err)
}

func TestAdd(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	items := map[float64]int64{1.0: 10, 2.0: 20}
	mock.ExpectDo("TDIGEST.ADD", "mykey", 1.0, int64(10), 2.0, int64(20)).SetVal("OK")

	err := tdigest.Add(ctx, db, "mykey", items)
	assert.NoError(t, err)

	mock.ExpectDo("TDIGEST.ADD", "mykey", 1.0, int64(10), 2.0, int64(20)).
		SetErr(errors.New("add failed"))

	err = tdigest.Add(ctx, db, "mykey", items)
	require.Error(t, err)
}

func TestMerge(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.MERGE", "destkey", "s1", "s2").SetVal("OK")

	err := tdigest.Merge(ctx, db, "destkey", "s1", "s2")
	assert.NoError(t, err)

	mock.ExpectDo("TDIGEST.MERGE", "destkey", "s1", "s2").SetErr(errors.New("merge failed"))

	err = tdigest.Merge(ctx, db, "destkey", "s1", "s2")
	require.Error(t, err)
}

func TestMin(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.MIN", "mykey").SetVal(1.23)

	val, err := tdigest.Min(ctx, db, "mykey")
	assert.NoError(t, err)
	assert.Equal(t, 1.23, val)

	mock.ExpectDo("TDIGEST.MIN", "mykey").SetErr(errors.New("min failed"))

	_, err = tdigest.Min(ctx, db, "mykey")
	require.Error(t, err)

	mock.ExpectDo("TDIGEST.MIN", "mykey").SetVal("not a float")

	_, err = tdigest.Min(ctx, db, "mykey")
	require.Error(t, err)
}

func TestMax(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.MAX", "mykey").SetVal(4.56)

	val, err := tdigest.Max(ctx, db, "mykey")
	assert.NoError(t, err)
	assert.Equal(t, 4.56, val)

	mock.ExpectDo("TDIGEST.MAX", "mykey").SetErr(errors.New("max failed"))

	_, err = tdigest.Max(ctx, db, "mykey")
	require.Error(t, err)

	mock.ExpectDo("TDIGEST.MAX", "mykey").SetVal("not a float")

	_, err = tdigest.Max(ctx, db, "mykey")
	require.Error(t, err)
}

func TestQuantile(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.QUANTILE", "mykey", 0.5).SetVal(7.89)

	val, err := tdigest.Quantile(ctx, db, "mykey", 0.5)
	assert.NoError(t, err)
	assert.Equal(t, 7.89, val)

	mock.ExpectDo("TDIGEST.QUANTILE", "mykey", 0.5).SetErr(errors.New("quantile failed"))

	_, err = tdigest.Quantile(ctx, db, "mykey", 0.5)
	require.Error(t, err)

	mock.ExpectDo("TDIGEST.QUANTILE", "mykey", 0.5).SetVal("not a float")

	_, err = tdigest.Quantile(ctx, db, "mykey", 0.5)
	require.Error(t, err)
}

func TestCDF(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()

	mock.ExpectDo("TDIGEST.CDF", "mykey", 10.0).SetVal(0.99)

	val, err := tdigest.CDF(ctx, db, "mykey", 10.0)
	assert.NoError(t, err)
	assert.Equal(t, 0.99, val)

	mock.ExpectDo("TDIGEST.CDF", "mykey", 10.0).SetErr(errors.New("cdf failed"))

	_, err = tdigest.CDF(ctx, db, "mykey", 10.0)
	require.Error(t, err)

	mock.ExpectDo("TDIGEST.CDF", "mykey", 10.0).SetVal("not a float")

	_, err = tdigest.CDF(ctx, db, "mykey", 10.0)
	require.Error(t, err)
}
