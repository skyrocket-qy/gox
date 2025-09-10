package main

import (
	"runtime"
	"time"
)

func CollectTimings(fn func(), repeat int) []int64 {
	timings := make([]int64, 0, repeat)
	for range repeat {
		start := time.Now()

		fn()

		elapsed := time.Since(start).Nanoseconds()
		timings = append(timings, elapsed)
	}

	return timings
}

// Measuring per-run heap usage
// Seeing how much memory is retained after a function.
func CollectMems(fn func(), repeat int) []uint64 {
	allocs := make([]uint64, 0, repeat)
	for range repeat {
		allocStart := SnapshotMemStats().Alloc

		fn()

		allocs = append(allocs, SnapshotMemStats().Alloc-allocStart)
	}

	return allocs
}

func CollectNumGCs(fn func(), repeat int) []uint32 {
	numGCs := make([]uint32, 0, repeat)
	for range repeat {
		st := SnapshotMemStats().NumGC

		fn()

		numGCs = append(numGCs, SnapshotMemStats().NumGC-st)
	}

	return numGCs
}

// How much memory was allocated during a benchmark
// GC pressure or allocation rate over time.
func CollectAllocs(fn func(), repeat int) []uint64 {
	allocs := make([]uint64, 0, repeat)

	var start, end runtime.MemStats
	for range repeat {
		runtime.ReadMemStats(&start)
		fn()
		runtime.ReadMemStats(&end)
		allocs = append(allocs, end.TotalAlloc-start.TotalAlloc)
	}

	return allocs
}

func CollectPauseNs(fn func(), repeat int) []uint64 {
	pauses := make([]uint64, 0, repeat)

	var start, end runtime.MemStats

	for range repeat {
		runtime.ReadMemStats(&start)
		fn()
		runtime.ReadMemStats(&end)

		if end.NumGC > start.NumGC {
			idx := (end.NumGC - 1) % uint32(len(end.PauseNs))
			pauses = append(pauses, end.PauseNs[idx])
		} else {
			pauses = append(pauses, 0)
		}
	}

	return pauses
}

func CollectGoroutines(fn func(), repeat int) []int {
	counts := make([]int, 0, repeat)
	for range repeat {
		before := runtime.NumGoroutine()

		fn()

		after := runtime.NumGoroutine()
		counts = append(counts, after-before)
	}

	return counts
}
