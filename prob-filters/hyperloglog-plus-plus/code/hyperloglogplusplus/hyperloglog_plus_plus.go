package hyperloglogplusplus

import (
	"fmt"
	"hash/fnv"
	"math"
)

// HyperLogLogPlusPlus represents an enhanced probabilistic data structure for cardinality estimation.
type HyperLogLogPlusPlus struct {
	m         uint32  // Number of registers (must be a power of 2)
	p         uint8   // log2(m)
	registers []uint8 // Stores the maximum number of leading zeros
	// For simplicity, we omit the sparse representation and bias correction
	// that are key features of HLL++. This is a conceptual example.
}

// New creates a new HyperLogLogPlusPlus instance.
// precision: determines the number of registers (m = 2^precision).
// A higher precision leads to more accuracy but uses more memory.
// Typical values for precision are between 4 and 16.
func New(precision uint8) *HyperLogLogPlusPlus {
	if precision < 4 || precision > 16 {
		precision = 14 // Default to a common precision
	}
	m := uint32(1 << precision)
	return &HyperLogLogPlusPlus{
		m:         m,
		p:         precision,
		registers: make([]uint8, m),
	}
}

// Add adds an element to the HyperLogLogPlusPlus structure.
func (hllpp *HyperLogLogPlusPlus) Add(data []byte) {
	hasher := fnv.New64a()
	hasher.Write(data)
	x := hasher.Sum64()

	// Get the register index (first p bits)
	j := x >> (64 - hllpp.p)

	// Get the number of leading zeros after the p bits
	w := x << hllpp.p               // Shift left to remove the first p bits
	rho := countLeadingZeros(w) + 1 // rho(w) + 1

	// Update the register if the new value is greater
	if rho > hllpp.registers[j] {
		hllpp.registers[j] = rho
	}
}

// countLeadingZeros counts the number of leading zeros in a uint64.
func countLeadingZeros(x uint64) uint8 {
	if x == 0 {
		return 64 // All zeros
	}
	var count uint8
	for i := 0; i < 64; i++ {
		if (x>>(63-i))&1 == 0 {
			count++
		} else {
			break
		}
	}
	return count
}

// Estimate returns the estimated cardinality.
// This simplified Estimate does not include the sparse representation
// or the advanced bias correction of a true HLL++.
func (hllpp *HyperLogLogPlusPlus) Estimate() float64 {
	sum := 0.0
	for _, val := range hllpp.registers {
		sum += 1.0 / math.Pow(2.0, float64(val))
	}

	alphaM2 := hllpp.alpha() * float64(hllpp.m) * float64(hllpp.m)
	estimate := alphaM2 / sum

	// Apply corrections for small and large cardinalities (simplified)
	if estimate <= 5.0/2.0*float64(hllpp.m) { // Small range correction
		V := 0 // Count of zero registers
		for _, val := range hllpp.registers {
			if val == 0 {
				V++
			}
		}
		if V != 0 {
			estimate = float64(hllpp.m) * math.Log(float64(hllpp.m)/float64(V))
		}
	} else if estimate > 1.0/30.0*math.Pow(2.0, 64.0) { // Large range correction (for 64-bit hash)
		estimate = -math.Pow(2.0, 64.0) * math.Log(1.0-estimate/math.Pow(2.0, 64.0))
	}

	return estimate
}

// alpha returns the bias correction factor based on the number of registers.
func (hllpp *HyperLogLogPlusPlus) alpha() float64 {
	switch hllpp.m {
	case 16:
		return 0.673
	case 32:
		return 0.697
	case 64:
		return 0.709
	default:
		return 0.7213 / (1.0 + 1.079/float64(hllpp.m)) // For m >= 128
	}
}

// Merge combines two HyperLogLogPlusPlus structures.
// This is a simplified merge, just taking the maximum of corresponding registers.
// A true HLL++ merge would also handle sparse representations.
func (hllpp *HyperLogLogPlusPlus) Merge(other *HyperLogLogPlusPlus) error {
	if hllpp.m != other.m || hllpp.p != other.p {
		return fmt.Errorf("cannot merge HLL++ instances with different precisions")
	}

	for i := range hllpp.registers {
		if other.registers[i] > hllpp.registers[i] {
			hllpp.registers[i] = other.registers[i]
		}
	}
	return nil
}
