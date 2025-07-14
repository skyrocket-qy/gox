package main

import (
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

type ProfileResult struct {
	ElapsedNS      int64
	MemAllocStart  uint64
	MemAllocEnd    uint64
	MemAllocDelta  int64
	NumGC          uint32
	CPUProfileFile string
}

func ProfileFunc(
	fn func(),
	enableCPU bool,
	cpuProfilePath string,
	repeat int,
) (*ProfileResult, error) {
	var result ProfileResult

	// CPU profile
	var stopCPU func()
	var err error
	if enableCPU {
		stopCPU, err = StartCPUProfile(cpuProfilePath)
		if err != nil {
			return nil, err
		}
		defer stopCPU()
		result.CPUProfileFile = cpuProfilePath
	}

	memStart := SnapshotMemStats()

	elapsed := MeasureTime(func() {
		for i := 0; i < repeat; i++ {
			fn()
		}
	})
	result.ElapsedNS = elapsed

	memEnd := SnapshotMemStats()
	result.MemAllocStart = memStart.Alloc
	result.MemAllocEnd = memEnd.Alloc
	result.MemAllocDelta = int64(memEnd.Alloc) - int64(memStart.Alloc)
	result.NumGC = memEnd.NumGC - memStart.NumGC

	return &result, nil
}

func StartCPUProfile(path string) (func(), error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	pprof.StartCPUProfile(f)
	return func() {
		pprof.StopCPUProfile()
		f.Close()
	}, nil
}

type MemStatsSnapshot struct {
	Alloc uint64
	NumGC uint32
}

func SnapshotMemStats() MemStatsSnapshot {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return MemStatsSnapshot{Alloc: m.Alloc, NumGC: m.NumGC}
}

func MeasureTime(fn func()) int64 {
	start := time.Now()
	fn()
	return time.Since(start).Nanoseconds()
}
