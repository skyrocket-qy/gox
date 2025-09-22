package pkg_test

import (
	"os"
	"testing"
	"time"

	"github.com/skyrocket-qy/gox/bench/pkg"
)

func TestMeasureTime(t *testing.T) {
	sleepDuration := 10 * time.Millisecond
	elapsed := pkg.MeasureTime(func() {
		time.Sleep(sleepDuration)
	})

	if elapsed < sleepDuration.Nanoseconds() {
		t.Errorf(
			"Expected elapsed time to be at least %d ns, got %d ns",
			sleepDuration.Nanoseconds(),
			elapsed,
		)
	}

	// Check if it's within a reasonable upper bound (e.g., 2x the sleep time)
	// to catch any significant delays.
	if elapsed > 2*sleepDuration.Nanoseconds() {
		t.Errorf(
			"Elapsed time %d ns is much higher than expected %d ns",
			elapsed,
			sleepDuration.Nanoseconds(),
		)
	}
}

func TestProfileFunc(t *testing.T) {
	// Test without CPU profiling
	result, err := pkg.ProfileFunc(func() { time.Sleep(1 * time.Millisecond) }, false, "", 1)
	if err != nil {
		t.Fatalf("ProfileFunc failed without CPU profiling: %v", err)
	}

	if result.ElapsedNS <= 0 {
		t.Errorf("Expected positive elapsed time, got %d", result.ElapsedNS)
	}

	if result.CPUProfileFile != "" {
		t.Errorf("Expected no CPU profile file, got %s", result.CPUProfileFile)
	}

	// Test with CPU profiling
	// Create a temporary directory for CPU profiles within the package's tmp directory.
	cpuProfileDir := "tmp"
	if err := os.MkdirAll(cpuProfileDir, 0o750); err != nil {
		t.Fatalf("Failed to create CPU profile directory: %v", err)
	}
	defer os.RemoveAll(cpuProfileDir)

	tmpfile, err := os.CreateTemp(cpuProfileDir, "cpuprofile_*.prof")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if err := tmpfile.Close(); err != nil {
		t.Errorf("Error closing temporary file: %v", err)
	} // Close the file so ProfileFunc can create it.

	result2, err := pkg.ProfileFunc(
		func() { time.Sleep(1 * time.Millisecond) },
		true,
		tmpfile.Name(),
		1,
	)
	if err != nil {
		t.Fatalf("ProfileFunc failed with CPU profiling: %v", err)
	}

	if result2.ElapsedNS <= 0 {
		t.Errorf("Expected positive elapsed time with profiling, got %d", result2.ElapsedNS)
	}

	if result2.CPUProfileFile != tmpfile.Name() {
		t.Errorf("Expected CPU profile file '%s', got '%s'", tmpfile.Name(), result2.CPUProfileFile)
	}

	// Check if the profile file was created and is not empty
	info, err := os.Stat(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to stat profile file: %v", err)
	}

	if info.Size() == 0 {
		t.Errorf("Expected profile file to be non-empty, but it has size 0")
	}
}
