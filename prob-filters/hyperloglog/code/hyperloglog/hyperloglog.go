package hyperloglog

import (
	"hash/fnv"
	"math"
)

// HyperLogLog represents a probabilistic data structure for cardinality estimation.
type HyperLogLog struct {
	m        uint32   // Number of registers (must be a power of 2)
	p        uint8    // log2(m)
	registers []uint8 // Stores the maximum number of leading zeros
}

// New creates a new HyperLogLog instance.
// precision: determines the number of registers (m = 2^precision).
// A higher precision leads to more accuracy but uses more memory.
// Typical values for precision are between 4 and 16.
func New(precision uint8) *HyperLogLog {
	if precision < 4 || precision > 16 {
		// HLL typically uses precision between 4 and 16
		// For simplicity, we'll enforce this range.
		precision = 14 // Default to a common precision
	}
	m := uint32(1 << precision)
	return &HyperLogLog{
		m:        m,
		p:        precision,
		registers: make([]uint8, m),
	}
}

// Add adds an element to the HyperLogLog structure.
func (hll *HyperLogLog) Add(data []byte) {
	hasher := fnv.New64a()
	hasher.Write(data)
	x := hasher.Sum64()

	// Get the register index (first p bits)
	j := x >> (64 - hll.p)

	// Get the number of leading zeros after the p bits
	w := x << hll.p // Shift left to remove the first p bits
	rho := countLeadingZeros(w) + 1 // rho(w) + 1

	// Update the register if the new value is greater
	if rho > hll.registers[j] {
		hll.registers[j] = rho
	}
}

// countLeadingZeros counts the number of leading zeros in a uint64.
func countLeadingZeros(x uint64) uint8 {
	if x == 0 {
		return 64 // All zeros
	}
	var count uint8
	for i := 0; i < 64; i++ {
		if (x >> (63 - i)) & 1 == 0 {
			count++
		} else {
			break
		}
	}
	return count
}

// Estimate returns the estimated cardinality.
func (hll *HyperLogLog) Estimate() float64 {
	sum := 0.0
	for _, val := range hll.registers {
		sum += 1.0 / math.Pow(2.0, float64(val))
	}

	alphaM2 := hll.alpha() * float64(hll.m) * float64(hll.m)
	estimate := alphaM2 / sum

	// Apply corrections for small and large cardinalities
	if estimate <= 5.0/2.0*float64(hll.m) { // Small range correction
		V := 0 // Count of zero registers
		for _, val := range hll.registers {
			if val == 0 {
				V++
			}
		}
		if V != 0 {
			estimate = float64(hll.m) * math.Log(float64(hll.m)/float64(V))
		}
	} else if estimate > 1.0/30.0*math.Pow(2.0, 64.0) { // Large range correction (for 64-bit hash)
		estimate = -math.Pow(2.0, 64.0) * math.Log(1.0 - estimate/math.Pow(2.0, 64.0))
	}

	return estimate
}

// alpha returns the bias correction factor based on the number of registers.
func (hll *HyperLogLog) alpha() float64 {
	switch hll.m {
	case 16:
		return 0.673
	case 32:
		return 0.697
	case 64:
		return 0.709
	default:
		return 0.7213 / (1.0 + 1.079/float64(hll.m)) // For m >= 128
	}
}