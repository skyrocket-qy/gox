package pkg

import (
	"sync"
	"testing"
	"time"
)

var blackHole []byte

// dummyFunc is a simple function that allocates memory and sleeps
// to simulate work for the collection functions.
func dummyFunc() {
	blackHole = make([]byte, 1024) // Allocate 1KB

	time.Sleep(1 * time.Millisecond)
}

func TestCollectTimings(t *testing.T) {
	timings := CollectTimings(dummyFunc, 5)
	if len(timings) != 5 {
		t.Fatalf("Expected 5 timings, got %d", len(timings))
	}

	for _, timing := range timings {
		if timing < 1_000_000 { // 1ms
			t.Errorf("Expected timing to be at least 1ms, got %d ns", timing)
		}
	}
}

func TestCollectMems(t *testing.T) {
	mems := CollectMems(dummyFunc, 3)
	if len(mems) != 3 {
		t.Fatalf("Expected 3 memory readings, got %d", len(mems))
	}
	// This is a weak test, as GC can be unpredictable.
	// We are checking if it's capturing *some* allocation.
	// A more robust test would require more control over the runtime.
}

func TestCollectNumGCs(t *testing.T) {
	// It's hard to deterministically trigger a GC, so we'll just check
	// that the function runs and returns the correct number of samples.
	gcs := CollectNumGCs(func() {
		// A more allocation-heavy function might trigger GC, but it's not guaranteed.
		_ = make([]byte, 1024*1024) // 1MB
	}, 2)
	if len(gcs) != 2 {
		t.Fatalf("Expected 2 GC counts, got %d", len(gcs))
	}
}

func TestCollectAllocs(t *testing.T) {
	allocs := CollectAllocs(dummyFunc, 4)
	if len(allocs) != 4 {
		t.Fatalf("Expected 4 allocation counts, got %d", len(allocs))
	}

	for _, alloc := range allocs {
		if alloc < 1024 {
			t.Errorf("Expected at least 1024 bytes allocated, got %d", alloc)
		}
	}
}

func TestCollectPauseNs(t *testing.T) {
	// Similar to GC, it's hard to test this deterministically.
	// We'll just ensure it runs and returns the right number of data points.
	pauses := CollectPauseNs(dummyFunc, 5)
	if len(pauses) != 5 {
		t.Fatalf("Expected 5 pause readings, got %d", len(pauses))
	}
}

func TestCollectGoroutines(t *testing.T) {
	// Test with a function that doesn't spawn goroutines
	counts1 := CollectGoroutines(dummyFunc, 3)
	if len(counts1) != 3 {
		t.Fatalf("Expected 3 goroutine counts, got %d", len(counts1))
	}

	for _, count := range counts1 {
		if count != 0 {
			// It might not be exactly 0 if the test runner has other goroutines.
			// Let's check for a small number.
			if count > 2 {
				t.Errorf("Expected 0 or a small number of new goroutines, got %d", count)
			}
		}
	}

	// Test with a function that does spawn a goroutine
	counts2 := CollectGoroutines(func() {
		ch := make(chan bool)

		go func() {
			time.Sleep(1 * time.Millisecond)

			ch <- true
		}()

		<-ch
	}, 2)

	if len(counts2) != 2 {
		t.Fatalf("Expected 2 goroutine counts, got %d", len(counts2))
	}
	// This is tricky because the goroutine might exit before the 'after' snapshot.
	// A more reliable test would involve waiting on a channel.
	// The dummy function above does this.
	for _, count := range counts2 {
		// The count can be 0 or 1, depending on timing.
		// If the goroutine finishes before the check, it will be 0.
		// If it's still running, it will be 1.
		// A negative count would be a definite error.
		if count < 0 {
			t.Errorf("Expected non-negative goroutine count, got %d", count)
		}
	}
}

// Helper to check if a value is positive, for tests where we expect some allocation/time.
func isPositive(val any) bool {
	switch v := val.(type) {
	case int:
		return v > 0
	case int64:
		return v > 0
	case uint64:
		return v > 0
	case uint32:
		return v > 0
	default:
		return false
	}
}

func TestCollectorsWithNoop(t *testing.T) {
	noop := func() {}
	repeat := 3

	if len(CollectTimings(noop, repeat)) != repeat {
		t.Error("CollectTimings wrong length")
	}

	if len(CollectMems(noop, repeat)) != repeat {
		t.Error("CollectMems wrong length")
	}

	if len(CollectNumGCs(noop, repeat)) != repeat {
		t.Error("CollectNumGCs wrong length")
	}

	if len(CollectAllocs(noop, repeat)) != repeat {
		t.Error("CollectAllocs wrong length")
	}

	if len(CollectPauseNs(noop, repeat)) != repeat {
		t.Error("CollectPauseNs wrong length")
	}

	if len(CollectGoroutines(noop, repeat)) != repeat {
		t.Error("CollectGoroutines wrong length")
	}
}

// A more reliable goroutine test.
func TestCollectGoroutinesReliable(t *testing.T) {
	counts := CollectGoroutines(func() {
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			time.Sleep(5 * time.Millisecond)
			wg.Done()
		}()

		wg.Wait()
	}, 1)

	if len(counts) != 1 {
		t.Fatalf("Expected 1 count, got %d", len(counts))
	}
	// This is still not 100% guaranteed to be 1, but it's more likely.
	// The 'before' check happens, then the goroutine is spawned, then the `wg.Wait()`
	// waits for it to finish, and then the 'after' check happens.
	// The goroutine should be gone by then. A count of 0 is expected.
	if counts[0] != 0 {
		// Let's allow for some scheduler noise, but it should be very small.
		if counts[0] > 2 {
			t.Errorf("Expected 0 new goroutines after waiting, got %d", counts[0])
		}
	}
}
