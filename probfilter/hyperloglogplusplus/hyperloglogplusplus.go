// Package hyperloglogplusplus implements the HyperLogLog and HyperLogLog++ cardinality //
// estimation algorithms. // These algorithms are used for accurately estimating the cardinality of
// a // multiset using constant memory. HyperLogLog++ has multiple improvements over // HyperLogLog,
// with a much lower error rate for smaller cardinalities. // // HyperLogLog is described here: //
// http://algo.inria.fr/flajolet/Publications/FlFuGaMe07.pdf // // HyperLogLog++ is described here:
// // http://research.google.com/pubs/pub40671.html
package hyperloglogplusplus

import (
	"bytes"
	"encoding/gob"
	"errors"
	"math"
)

const two32 = 1 << 32

type HyperLogLog struct {
	reg []uint8
	m   uint32
	p   uint8
}

// New returns a new initialized HyperLogLog.
func New(precision uint8) (*HyperLogLog, error) {
	if precision > 16 || precision < 4 {
		return nil, errors.New("precision must be between 4 and 16")
	}

	h := &HyperLogLog{}
	h.p = precision
	h.m = 1 << precision
	h.reg = make([]uint8, h.m)

	return h, nil
}

// Clear sets HyperLogLog h back to its initial state.
func (h *HyperLogLog) Clear() {
	h.reg = make([]uint8, h.m)
}

// Add adds a new item to HyperLogLog h.
func (h *HyperLogLog) Add(item Hash32) {
	x := item.Sum32()
	// Use the top p bits of the hash for the register index.
	// eb32(x, 32, h.p) extracts the top h.p bits from the 32-bit hash x,
	// which is equivalent to x >> (32 - h.p).
	i := eb32(x, 32, h.p) // {x31,...,x32-p}
	w := x<<h.p | 1       // {x32-p-1,...,x0}1

	zeroBits := clz32(w) + 1
	if zeroBits > h.reg[i] {
		h.reg[i] = zeroBits
	}
}

// Merge takes another HyperLogLog and combines it with HyperLogLog h.
func (h *HyperLogLog) Merge(other *HyperLogLog) error {
	if h.p != other.p {
		return errors.New("precisions must be equal")
	}

	for i, v := range other.reg {
		if v > h.reg[i] {
			h.reg[i] = v
		}
	}

	return nil
}

// Count returns the cardinality estimate.
func (h *HyperLogLog) Count() uint64 {
	est := calculateEstimate(h.reg)
	if est <= float64(h.m)*2.5 {
		if v := countZeros(h.reg); v != 0 {
			return uint64(linearCounting(h.m, v))
		}

		return uint64(est)
	} else if est < two32/30 {
		return uint64(est)
	}

	return uint64(-two32 * math.Log(1-est/two32))
}

// Encode HyperLogLog into a gob.
func (h *HyperLogLog) GobEncode() ([]byte, error) {
	buf := bytes.Buffer{}

	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(h.reg); err != nil {
		return nil, err
	}

	if err := enc.Encode(h.m); err != nil {
		return nil, err
	}

	if err := enc.Encode(h.p); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decode gob into a HyperLogLog structure.
func (h *HyperLogLog) GobDecode(b []byte) error {
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	if err := dec.Decode(&h.reg); err != nil {
		return err
	}

	if err := dec.Decode(&h.m); err != nil {
		return err
	}

	if err := dec.Decode(&h.p); err != nil {
		return err
	}

	return nil
}

// Hash32 is the interface that wraps the basic Sum32 method.
type Hash32 interface {
	Sum32() uint32
}

// eb32 extracts bits from a uint32.
// x: the uint32 to extract from
// k: the number of bits in x
// j: the number of bits to extract
func eb32(x uint32, k, j uint8) uint32 {
	return (x >> (k - j))
}

// clz32 counts leading zeros in a uint32.
func clz32(x uint32) uint8 {
	if x == 0 {
		return 32
	}

	var n uint8
	if (x & 0xFFFF0000) == 0 {
		n = n + 16
		x = x << 16
	}

	if (x & 0xFF000000) == 0 {
		n = n + 8
		x = x << 8
	}

	if (x & 0xF0000000) == 0 {
		n = n + 4
		x = x << 4
	}

	if (x & 0xC0000000) == 0 {
		n = n + 2
		x = x << 2
	}

	if (x & 0x80000000) == 0 {
		n = n + 1
	}

	return n
}

// calculateEstimate calculates the estimate from the registers.
func calculateEstimate(reg []uint8) float64 {
	var sum float64
	for _, v := range reg {
		sum += math.Pow(2, -float64(v))
	}

	alpha := 0.7213 / (1 + 1.079/float64(len(reg)))

	return alpha * float64(len(reg)*len(reg)) / sum
}

// linearCounting calculates the linear counting estimate.
func linearCounting(m, v uint32) float64 {
	return float64(m) * math.Log(float64(m)/float64(v))
}

// countZeros counts the number of zero registers.
func countZeros(reg []uint8) uint32 {
	var count uint32

	for _, v := range reg {
		if v == 0 {
			count++
		}
	}

	return count
}
