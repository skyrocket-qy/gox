package hyperloglogplusplus_test

import (
	"strconv"
	"testing"

	"github.com/skyrocket-qy/gox/probfilter/hyperloglogplusplus"
	"github.com/stretchr/testify/assert"
)

// a simple hash function for testing
type stringHash string

func (s stringHash) Sum32() uint32 {
	h := uint32(2166136261)
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

func TestNew(t *testing.T) {
	_, err := hyperloglogplusplus.New(3)
	assert.Error(t, err)

	_, err = hyperloglogplusplus.New(17)
	assert.Error(t, err)

	hll, err := hyperloglogplusplus.New(14)
	assert.NoError(t, err)
	assert.NotNil(t, hll)
}

func TestAddAndCount(t *testing.T) {
	hll, _ := hyperloglogplusplus.New(14)

	for i := 0; i < 100000; i++ {
		hll.Add(stringHash(strconv.Itoa(i)))
	}

	count := hll.Count()
	// Allow for a 10% error margin
	errorMargin := float64(100000) * 0.1
	assert.InDelta(t, 100000, count, errorMargin, "Expected count to be around 100000")
}

func TestClear(t *testing.T) {
	hll, _ := hyperloglogplusplus.New(14)
	hll.Add(stringHash("1"))
	hll.Add(stringHash("2"))
	hll.Clear()
	assert.Equal(t, uint64(0), hll.Count())
}

func TestMerge(t *testing.T) {
	hll1, _ := hyperloglogplusplus.New(14)
	hll2, _ := hyperloglogplusplus.New(14)
	hll3, _ := hyperloglogplusplus.New(15)

	for i := 0; i < 50000; i++ {
		hll1.Add(stringHash(strconv.Itoa(i)))
	}

	for i := 50000; i < 100000; i++ {
		hll2.Add(stringHash(strconv.Itoa(i)))
	}

	err := hll1.Merge(hll3)
	assert.Error(t, err, "Expected error when merging HLLs with different precisions")

	err = hll1.Merge(hll2)
	assert.NoError(t, err)

	count := hll1.Count()
	errorMargin := float64(100000) * 0.1
	assert.InDelta(t, 100000, count, errorMargin, "Expected count to be around 100000")
}

func TestGobEncodeDecode(t *testing.T) {
	hll, _ := hyperloglogplusplus.New(14)
	for i := 0; i < 10000; i++ {
		hll.Add(stringHash(strconv.Itoa(i)))
	}

	buf, err := hll.GobEncode()
	assert.NoError(t, err)

	hll2 := &hyperloglogplusplus.HyperLogLog{}
	err = hll2.GobDecode(buf)
	assert.NoError(t, err)

	assert.Equal(t, hll.Count(), hll2.Count())
}
