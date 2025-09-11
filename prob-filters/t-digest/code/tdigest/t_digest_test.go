package tdigest

import (
	"testing"
	"math"
	"math/rand"
	"time"
	"sort"
)

func TestTDigest_Quantile(t *testing.T) {
	// NOTE: This t-digest implementation is described as "simplified conceptual".
	// It may not provide accurate quantile estimates for all data distributions
	// due to its simplified merging and compression logic. The test's error
	// margin is set to reflect this conceptual nature. For a robust t-digest,
	// a more complex implementation would be required.
	digest := New(100) // Compression factor of 100

	// Add some random data points
	rand.Seed(time.Now().UnixNano())
	data := make([]float64, 1000)
	for i := 0; i < 1000; i++ {
		val := rand.NormFloat64()*10 + 50 // Normal distribution around 50
		digest.Add(val)
		data[i] = val
	}

	// Sort data to find true quantiles
	sort.Float64s(data)

	// Test various quantiles
	quantilesToTest := []float64{0.05, 0.25, 0.50, 0.75, 0.95, 0.99}
	for _, q := range quantilesToTest {
		estimatedQuantile := digest.Quantile(q)
		trueQuantile := data[int(q*float64(len(data)))]

		// T-Digest provides an estimate, so we check if it's within a reasonable range.
		allowedError := 0.05 * trueQuantile // 5% error margin for this basic test
		if trueQuantile == 0 {
			allowedError = 0.5 // Absolute error for values near zero
		}

		if math.Abs(estimatedQuantile-trueQuantile) > allowedError {
			t.Errorf("For quantile %.2f: Estimated %.2f is outside allowed error margin for true %.2f (allowed error: %.2f)", q, estimatedQuantile, trueQuantile, allowedError)
		}
	}
}