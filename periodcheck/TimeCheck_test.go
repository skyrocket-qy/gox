package periodcheck

import (
	"errors"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Helper to capture log output
func captureOutput(f func()) string {
	var buf strings.Builder
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	f()
	return buf.String()
}

func TestFixedTimeCheck(t *testing.T) {
	// Test case 1: Condition met within timeout
	t.Run("ConditionMet", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++
			if callCount == 2 {
				return 10, nil
			}
			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}
	
		err := FixedTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 100*time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 2, callCount)
	})

	// Test case 2: Timeout reached
	t.Run("Timeout", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++
			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := FixedTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 10*time.Millisecond)
		assert.Error(t, err)
		assert.EqualError(t, err, "timeout")
		assert.GreaterOrEqual(t, callCount, 2)
	})

	// Test case 3: getCurrentStatus returns an error
	t.Run("GetCurrentStatusError", func(t *testing.T) {
		expectedErr := errors.New("failed to get status")
		getCurrentStatus := func() (int, error) {
			return 0, expectedErr
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := FixedTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 100*time.Millisecond)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestExponentialTimeCheck(t *testing.T) {
	// Test case 1: Condition met within timeout
	t.Run("ConditionMet", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++
			if callCount == 3 {
				return 10, nil
			}
			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := ExponentialTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 100*time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 3, callCount)
	})

	// Test case 2: Timeout reached
	t.Run("Timeout", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++
			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := ExponentialTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 10*time.Millisecond)
		assert.Error(t, err)
		assert.EqualError(t, err, "timeout")
		assert.GreaterOrEqual(t, callCount, 2)
	})

	// Test case 3: getCurrentStatus returns an error
	t.Run("GetCurrentStatusError", func(t *testing.T) {
		expectedErr := errors.New("failed to get status")
		getCurrentStatus := func() (int, error) {
			return 0, expectedErr
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := ExponentialTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 100*time.Millisecond)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestDiffTimeCheck(t *testing.T) {
	// Test case 1: Condition met within timeout
	t.Run("ConditionMet", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++
			if callCount == 2 {
				return 10, nil
			}
			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := DiffTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 10*time.Millisecond, 100*time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 2, callCount)
	})

	// Test case 2: Timeout reached
	t.Run("Timeout", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++
			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := DiffTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 10*time.Millisecond, 10*time.Millisecond)
		assert.Error(t, err)
		assert.EqualError(t, err, "timeout")
		assert.GreaterOrEqual(t, callCount, 2)
	})

	// Test case 3: getCurrentStatus returns an error
	t.Run("GetCurrentStatusError", func(t *testing.T) {
		expectedErr := errors.New("failed to get status")
		getCurrentStatus := func() (int, error) {
			return 0, expectedErr
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := DiffTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 10*time.Millisecond, 100*time.Millisecond)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestSelfAdaptiveTimeCheck(t *testing.T) {
	// Test case 1: Condition met within timeout
	t.Run("ConditionMet", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++
			if callCount == 2 {
				return 10, nil
			}
			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := SelfAdaptiveTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 10*time.Millisecond, 100*time.Millisecond)
		assert.NoError(t, err)
		assert.Equal(t, 2, callCount)
	})

	// Test case 2: Timeout reached
	t.Run("Timeout", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++
			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := SelfAdaptiveTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 10*time.Millisecond, 10*time.Millisecond)
		assert.Error(t, err)
		assert.EqualError(t, err, "timeout")
		assert.GreaterOrEqual(t, callCount, 2)
	})

	// Test case 3: getCurrentStatus returns an error
	t.Run("GetCurrentStatusError", func(t *testing.T) {
		expectedErr := errors.New("failed to get status")
		getCurrentStatus := func() (int, error) {
			return 0, expectedErr
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := SelfAdaptiveTimeCheck(getCurrentStatus, 10, checkFunc, 1*time.Millisecond, 10*time.Millisecond, 100*time.Millisecond)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 5, abs(5))
	assert.Equal(t, 5, abs(-5))
	assert.Equal(t, 0, abs(0))
}
