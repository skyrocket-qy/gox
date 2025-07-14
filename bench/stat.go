package main

import (
	"math"
	"slices"
)

type StatResult struct {
	Count   int
	Min     int64
	Max     int64
	Average float64
	Median  float64
	P90     float64
	P95     float64
	P99     float64
	P999    float64
	StdDev  float64
}

func AnalyzeInts[T ~int | ~int64 | ~int32 | ~uint | ~uint64 | ~uint32](data []T) StatResult {
	n := len(data)
	if n == 0 {
		return StatResult{}
	}

	// Convert to []int64 for ease of processing
	values := make([]int64, n)
	var sum int64
	for i, v := range data {
		values[i] = int64(v)
		sum += values[i]
	}

	slices.Sort(values)

	avg := float64(sum) / float64(n)
	median := percentile(values, 50)
	p90 := percentile(values, 90)
	p95 := percentile(values, 95)
	p99 := percentile(values, 99)
	p999 := percentile(values, 99.9)

	stddev := func() float64 {
		var sqsum float64
		for _, v := range values {
			d := float64(v) - avg
			sqsum += d * d
		}
		return math.Sqrt(sqsum / float64(n))
	}()

	return StatResult{
		Count:   n,
		Min:     values[0],
		Max:     values[n-1],
		Average: avg,
		Median:  median,
		P90:     p90,
		P95:     p95,
		P99:     p99,
		P999:    p999,
		StdDev:  stddev,
	}
}

func percentile(sorted []int64, p float64) float64 {
	if len(sorted) == 0 {
		return 0
	}
	rank := p / 100 * float64(len(sorted)-1)
	lower := int(math.Floor(rank))
	upper := int(math.Ceil(rank))
	if lower == upper {
		return float64(sorted[lower])
	}
	weight := rank - float64(lower)
	return float64(sorted[lower])*(1-weight) + float64(sorted[upper])*weight
}
