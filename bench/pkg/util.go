package pkg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"
)

type ProfileResult struct {
	ElapsedNS      int64
	MemAllocStart  uint64
	MemAllocEnd    uint64
	MemAllocDelta  uint64
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
	var (
		stopCPU func()
		err     error
	)

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
		for range repeat {
			fn()
		}
	})
	result.ElapsedNS = elapsed

	memEnd := SnapshotMemStats()
	result.MemAllocStart = memStart.Alloc
	result.MemAllocEnd = memEnd.Alloc
	result.MemAllocDelta = memEnd.Alloc - memStart.Alloc
	result.NumGC = memEnd.NumGC - memStart.NumGC

	return &result, nil
}

func StartCPUProfile(path string) (func(), error) {
	// G304 (CWE-22): Ensure path is within the 'tmp' directory.
	tmpDir, err := filepath.Abs("tmp")
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for tmp directory: %w", err)
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get absolute path for CPU profile file: %w", err)
	}

	if !strings.HasPrefix(absPath, tmpDir) {
		return nil, fmt.Errorf(
			"CPU profile file path %s is outside of the allowed directory %s",
			absPath,
			tmpDir,
		)
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		if cerr := f.Close(); cerr != nil {
			log.Printf("Error closing file after failed CPU profile start: %v", cerr)
		}

		return nil, err
	}

	return func() {
		pprof.StopCPUProfile()

		if err := f.Close(); err != nil {
			// Log the error, as we can't return it from a deferred function
			// In a real application, you might use a proper logging library
			log.Printf("Error closing CPU profile file: %v", err)
		}
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
