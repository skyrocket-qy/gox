package periodcheck_test

import (
	"errors"
	"testing"
	"time"

	"github.com/skyrocket-qy/gox/periodcheck"
	"github.com/stretchr/testify/assert"
)

// Common test runner for TimeCheck functions with 5 parameters.
func runCommonTimeCheckTests(
	t *testing.T,
	timeCheckFunc func(
		getCurrentStatus func() (int, error),
		target int,
		checkFunc func(cur, target int) bool,
		interval, timeout time.Duration,
	) error,
	expectedConditionMetCallCount int,
) {
	t.Run("ConditionMet", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++
			if callCount == expectedConditionMetCallCount {
				return 10, nil
			}

			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := timeCheckFunc(
			getCurrentStatus,
			10,
			checkFunc,
			1*time.Millisecond,
			100*time.Millisecond,
		)
		assert.NoError(t, err)
		assert.Equal(t, expectedConditionMetCallCount, callCount)
	})

	t.Run("Timeout", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++

			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := timeCheckFunc(
			getCurrentStatus,
			10,
			checkFunc,
			1*time.Millisecond,
			10*time.Millisecond,
		)
		assert.Error(t, err)
		assert.EqualError(t, err, "timeout")
		assert.GreaterOrEqual(t, callCount, 2)
	})

	t.Run("GetCurrentStatusError", func(t *testing.T) {
		expectedErr := errors.New("failed to get status")
		getCurrentStatus := func() (int, error) {
			return 0, expectedErr
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := timeCheckFunc(
			getCurrentStatus,
			10,
			checkFunc,
			1*time.Millisecond,
			100*time.Millisecond,
		)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

// Common test runner for TimeCheck functions with 6 parameters (DiffTimeCheck,
// SelfAdaptiveTimeCheck).
func runCommonTimeCheckTestsWithMaxInterval(
	t *testing.T,
	timeCheckFunc func(
		getCurrentStatus func() (int, error),
		target int,
		checkFunc func(cur, target int) bool,
		interval, minInterval, maxInterval time.Duration,
	) error,
	expectedConditionMetCallCount int,
) {
	t.Run("ConditionMet", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++
			if callCount == expectedConditionMetCallCount {
				return 10, nil
			}

			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := timeCheckFunc(
			getCurrentStatus,
			10,
			checkFunc,
			1*time.Millisecond,
			10*time.Millisecond,
			100*time.Millisecond,
		)
		assert.NoError(t, err)
		assert.Equal(t, expectedConditionMetCallCount, callCount)
	})

	t.Run("Timeout", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++

			return 0, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := timeCheckFunc(
			getCurrentStatus,
			10,
			checkFunc,
			1*time.Millisecond,
			10*time.Millisecond,
			10*time.Millisecond,
		)
		assert.Error(t, err)
		assert.EqualError(t, err, "timeout")
		assert.GreaterOrEqual(t, callCount, 2)
	})

	t.Run("GetCurrentStatusError", func(t *testing.T) {
		expectedErr := errors.New("failed to get status")
		getCurrentStatus := func() (int, error) {
			return 0, expectedErr
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := timeCheckFunc(
			getCurrentStatus,
			10,
			checkFunc,
			1*time.Millisecond,
			10*time.Millisecond,
			100*time.Millisecond,
		)
		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestFixedTimeCheck(t *testing.T) {
	runCommonTimeCheckTests(t, periodcheck.FixedTimeCheck, 2)
}

func TestExponentialTimeCheck(t *testing.T) {
	runCommonTimeCheckTests(t, periodcheck.ExponentialTimeCheck, 3)
}

func TestDiffTimeCheck(t *testing.T) {
	runCommonTimeCheckTestsWithMaxInterval(t, periodcheck.DiffTimeCheck, 2)
}

func TestSelfAdaptiveTimeCheck(t *testing.T) {
	runCommonTimeCheckTestsWithMaxInterval(t, periodcheck.SelfAdaptiveTimeCheck, 2)

	t.Run("DenominatorIsZero", func(t *testing.T) {
		callCount := 0
		getCurrentStatus := func() (int, error) {
			callCount++
			if callCount == 3 {
				return 10, nil
			}
			// This will make the diff the same for the first two calls
			return 5, nil
		}
		checkFunc := func(cur, target int) bool {
			return cur == target
		}

		err := periodcheck.SelfAdaptiveTimeCheck( // Added periodcheck.
			getCurrentStatus,
			10,
			checkFunc,
			1*time.Millisecond,
			10*time.Millisecond,
			100*time.Millisecond,
		)
		assert.NoError(t, err)
		assert.Equal(t, 3, callCount)
	})
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 5, periodcheck.Abs(5))  // Changed to periodcheck.Abs
	assert.Equal(t, 5, periodcheck.Abs(-5)) // Changed to periodcheck.Abs
	assert.Equal(t, 0, periodcheck.Abs(0))  // Changed to periodcheck.Abs
}
