package pkg

import (
	"math"
	"testing"
)

func TestAnalyzeInts(t *testing.T) {
	// Test case 1: Basic test with a simple integer slice
	data1 := []int{1, 2, 3, 4, 5}
	result1 := AnalyzeInts(data1)
	if result1.Count != 5 {
		t.Errorf("Test 1: Expected Count 5, got %d", result1.Count)
	}
	if result1.Min != 1 {
		t.Errorf("Test 1: Expected Min 1, got %d", result1.Min)
	}
	if result1.Max != 5 {
		t.Errorf("Test 1: Expected Max 5, got %d", result1.Max)
	}
	if result1.Average != 3.0 {
		t.Errorf("Test 1: Expected Average 3.0, got %f", result1.Average)
	}
	if result1.Median != 3.0 {
		t.Errorf("Test 1: Expected Median 3.0, got %f", result1.Median)
	}

	// Test case 2: Empty slice
	data2 := []int{}
	result2 := AnalyzeInts(data2)
	if result2.Count != 0 {
		t.Errorf("Test 2: Expected Count 0, got %d", result2.Count)
	}

	// Test case 3: Slice with one element
	data3 := []int{10}
	result3 := AnalyzeInts(data3)
	if result3.Count != 1 {
		t.Errorf("Test 3: Expected Count 1, got %d", result3.Count)
	}
	if result3.Min != 10 || result3.Max != 10 || result3.Average != 10.0 {
		t.Errorf("Test 3: Incorrect values for single-element slice")
	}

	// Test case 4: More complex data
	data4 := []int64{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	result4 := AnalyzeInts(data4)
	if result4.P90 != 91.0 {
		t.Errorf("Test 4: Expected P90 91.0, got %f", result4.P90)
	}
	if result4.P99 != 99.1 {
		t.Errorf("Test 4: Expected P99 99.1, got %f", result4.P99)
	}
}

func TestPercentile(t *testing.T) {
	sortedData := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Test 50th percentile (median)
	p50 := percentile(sortedData, 50)
	if p50 != 5.5 {
		t.Errorf("Expected 50th percentile to be 5.5, got %f", p50)
	}

	// Test 90th percentile
	p90 := percentile(sortedData, 90)
	if p90 != 9.1 {
		t.Errorf("Expected 90th percentile to be 9.1, got %f", p90)
	}

	// Test 100th percentile
	p100 := percentile(sortedData, 100)
	if p100 != 10 {
		t.Errorf("Expected 100th percentile to be 10, got %f", p100)
	}

	// Test with a single element
	singleElement := []int64{100}
	p_single := percentile(singleElement, 50)
	if p_single != 100 {
		t.Errorf("Expected 50th percentile of single element to be 100, got %f", p_single)
	}

	// Test with empty slice
	emptySlice := []int64{}
	p_empty := percentile(emptySlice, 50)
	if p_empty != 0 {
		t.Errorf("Expected percentile of empty slice to be 0, got %f", p_empty)
	}
}

const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestAnalyzeInts_StdDev(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	result := AnalyzeInts(data)
	expectedStdDev := 1.41421356237
	if !almostEqual(result.StdDev, expectedStdDev) {
		t.Errorf("Expected StdDev %f, got %f", expectedStdDev, result.StdDev)
	}
}
