package hyperloglogplusplus_test

import (
	"bytes"
	"encoding/gob"
	"strconv"
	"testing"

	"github.com/skyrocket-qy/gox/probfilter/hyperloglogplusplus"
	"github.com/stretchr/testify/assert"
)

// a simple hash function for testing.
type stringHash string

func (s stringHash) Sum32() uint32 {
	h := uint32(2166136261)
	for i := range len(s) {
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

	for i := range 100000 {
		hll.Add(stringHash(strconv.Itoa(i)))
	}

	count := hll.Count()
	// Allow for a 10% error margin
	errorMargin := float64(100000) * 0.1
	assert.InDelta(t, 100000, count, errorMargin, "Expected count to be around 100000")
}

func TestAddAndCount_SmallCardinality(t *testing.T) {
	hll, _ := hyperloglogplusplus.New(14)

	hll.Add(stringHash("a"))
	hll.Add(stringHash("b"))
	hll.Add(stringHash("c"))

	count := hll.Count()
	assert.InDelta(t, 3, count, 1, "Expected count to be around 3")
}

func TestAddAndCount_NoZeroRegisters(t *testing.T) {
	hll, _ := hyperloglogplusplus.New(4) // p=4, m=16

	// Add enough items to make it likely that all registers are non-zero.
	for i := range 100 {
		hll.Add(stringHash(strconv.Itoa(i)))
	}

	// This test is to cover the branch where countZeros is 0.
	// We don't need to assert the count, just that the code runs.
	hll.Count()
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

	for i := range 50000 {
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
	for i := range 10000 {
		hll.Add(stringHash(strconv.Itoa(i)))
	}

	buf, err := hll.GobEncode()
	assert.NoError(t, err)

	hll2 := &hyperloglogplusplus.HyperLogLog{}
	err = hll2.GobDecode(buf)
	assert.NoError(t, err)

	assert.Equal(t, hll.Count(), hll2.Count())
}

func TestGobDecodeError(t *testing.T) {
	hll := &hyperloglogplusplus.HyperLogLog{}

	// Test with a corrupted buffer that causes an error in the first Decode
	err := hll.GobDecode([]byte{0x01, 0x02, 0x03})
	assert.Error(t, err)

	// Test with a buffer that causes an error in the second Decode
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	enc.Encode([]uint8{1, 2, 3}) // reg

	err = hll.GobDecode(buf.Bytes())
	assert.Error(t, err)

	// Test with a buffer that causes an error in the third Decode
	buf.Reset()
	enc.Encode([]uint8{1, 2, 3}) // reg
	enc.Encode(uint32(16384))    // m

	err = hll.GobDecode(buf.Bytes())
	assert.Error(t, err)
}
