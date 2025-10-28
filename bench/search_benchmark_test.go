package main

import (
	"math/rand"
	"testing"
)

// generateSlice generates a slice of n random integers.
func generateSlice(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(n * 2) // Values up to 2n to ensure some misses
	}
	return s
}

// generateMap generates a map of n random integers.
func generateMap(n int) map[int]struct{} {
	m := make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		m[rand.Intn(n*2)] = struct{}{}
	}
	return m
}

// searchSlice searches for a value in a slice.
func searchSlice(s []int, val int) bool {
	for _, v := range s {
		if v == val {
			return true
		}
	}
	return false
}

// searchMap searches for a value in a map.
func searchMap(m map[int]struct{}, val int) bool {
	_, found := m[val]
	return found
}

func benchmarkSearch(b *testing.B, n int, searchFunc func(val int) bool) {
	b.StopTimer()
	// Generate data once per benchmark run
	data := generateSlice(n) // Use slice to generate values for both slice and map
	m := make(map[int]struct{}, n)
	for _, v := range data {
		m[v] = struct{}{}
	}

	// Pick a value to search for (can be present or not)
	searchValue := rand.Intn(n * 2)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		searchFunc(searchValue)
	}
}

const N = 50

func BenchmarkSearchSlice_N25(b *testing.B) {
	s := generateSlice(N)
	benchmarkSearch(b, N, func(val int) bool { return searchSlice(s, val) })
}

func BenchmarkSearchMap_N25(b *testing.B) {
	m := generateMap(N)
	benchmarkSearch(b, N, func(val int) bool { return searchMap(m, val) })
}
