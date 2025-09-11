package tdigest

import (
	"sort"
	"math"
)

// Centroid represents a data point in the t-digest.
type Centroid struct {
	Mean  float64
	Count float64
}

// TDigest represents a probabilistic data structure for estimating quantiles.
type TDigest struct {
	centroids []Centroid
	compression float64 // Controls the number of centroids
	maxCentroids int
}

// New creates a new TDigest instance.
// compression: A higher value means more centroids and better accuracy, but more memory.
// Typical values are 20, 50, 100.
func New(compression float64) *TDigest {
	if compression <= 0 {
		compression = 20 // Default compression
	}
	return &TDigest{
		centroids: make([]Centroid, 0),
		compression: compression,
		maxCentroids: int(compression * 5), // Heuristic for max centroids
	}
}

// Add adds a value to the t-digest.
func (td *TDigest) Add(value float64) {
	td.AddWithWeight(value, 1.0)
}

// AddWithWeight adds a value with a given weight to the t-digest.
func (td *TDigest) AddWithWeight(value, weight float64) {
	// Find the closest centroid
	closestIdx := -1
	minDist := math.MaxFloat64
	for i, c := range td.centroids {
		dist := math.Abs(c.Mean - value)
		if dist < minDist {
			minDist = dist
			closestIdx = i
		}
	}

	if closestIdx != -1 && td.canMerge(td.centroids[closestIdx], weight) {
		// Merge with closest centroid
		c := td.centroids[closestIdx]
		newCount := c.Count + weight
		c.Mean = (c.Mean*c.Count + value*weight) / newCount
		c.Count = newCount
		td.centroids[closestIdx] = c
	} else {
		// Add as a new centroid
		td.centroids = append(td.centroids, Centroid{Mean: value, Count: weight})
	}

	// Periodically compress the centroids
	if len(td.centroids) > td.maxCentroids {
		td.Compress()
	}
}

// canMerge determines if a new value can be merged with an existing centroid.
// This is a simplified rule. Real t-digest uses a more complex quantile-based merging.
func (td *TDigest) canMerge(c Centroid, weight float64) bool {
	// Simplified: allow merging if the centroid's count is not too large
	// relative to the compression factor.
	return c.Count < td.compression * 10 // Arbitrary threshold for conceptual example
}

// Compress reduces the number of centroids by merging nearby ones.
// This is a simplified compression. Real t-digest sorts centroids and merges
// based on quantile-based distance.
func (td *TDigest) Compress() {
	if len(td.centroids) <= 1 {
		return
	}

	// Sort centroids by mean for easier merging
	sort.Slice(td.centroids, func(i, j int) bool {
		return td.centroids[i].Mean < td.centroids[j].Mean
	})

	newCentroids := make([]Centroid, 0, len(td.centroids))
	current := td.centroids[0]

	for i := 1; i < len(td.centroids); i++ {
		next := td.centroids[i]
		// Simplified merging logic: merge if close enough
		if math.Abs(current.Mean - next.Mean) < (current.Mean * 0.01) && len(newCentroids) < td.maxCentroids {
			// Merge
			newCount := current.Count + next.Count
			current.Mean = (current.Mean*current.Count + next.Mean*next.Count) / newCount
			current.Count = newCount
		} else {
			newCentroids = append(newCentroids, current)
			current = next
		}
	}
	newCentroids = append(newCentroids, current)
	td.centroids = newCentroids
}

// Quantile estimates the value at a given quantile (0.0 to 1.0).
func (td *TDigest) Quantile(q float64) float64 {
	if len(td.centroids) == 0 {
		return 0.0
	}
	if q < 0.0 { q = 0.0 }
	if q > 1.0 { q = 1.0 }

	// Ensure centroids are sorted for quantile calculation
	sort.Slice(td.centroids, func(i, j int) bool {
		return td.centroids[i].Mean < td.centroids[j].Mean
	})

	totalCount := 0.0
	for _, c := range td.centroids {
		totalCount += c.Count
	}

	if totalCount == 0 {
		return 0.0
	}

	targetCount := q * totalCount
	currentCount := 0.0

	for i, c := range td.centroids {
		currentCount += c.Count
		if currentCount >= targetCount {
			// Found the centroid containing the quantile
			if i == 0 {
				return c.Mean // First centroid
			}
			prev := td.centroids[i-1]
			// Linear interpolation between centroids
			return prev.Mean + (c.Mean - prev.Mean) * ((targetCount - (currentCount - c.Count)) / (c.Count))
		}
	}
	return td.centroids[len(td.centroids)-1].Mean // Last centroid
}