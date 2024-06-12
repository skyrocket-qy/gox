/*
TimeCheck waits for the condition defined by the `checkFunc` to be met within a specified timeout duration.
It uses the `getCurrentStatus` function to get the current status for checking.
The function sleeps for a fixed interval between checks.
If the condition is met within the timeout, the function returns nil. If not, it returns an error with the message "timeout".

Parameters:
  - getCurrentStatus: A function that retrieves the current status. It should return the current status and an error (if any).
  - targetStatus: The target status that the function is waiting for.
  - checkFunc: A function that checks if the current status meets the target status.
  - interval: The interval duration to wait between status checks.
  - timeout: The maximum duration to wait for the condition to be met.

Returns:
  - An error indicating the result of the operation. Returns nil if the condition is met within the timeout.

Example Usage:

	getCurrentStatus := func() (int, error) {
	    // Implementation of retrieving current status
	}
	targetStatus := 10
	checkFunc := func(cur, target int) bool {
	    return cur == target
	}
	interval := 1 * time.Second
	timeout := 5 * time.Second
	err := FixedTimeCheck(getCurrentStatus, targetStatus, checkFunc, interval, timeout)
	if err != nil {
	    log.Printf("Error: %s\n", err.Error())
	} else {
	    log.Println("Condition met within timeout")
	}
*/

package timecheck

import (
	"errors"
	"log"
	"time"
)

func FixedTimeCheck(
	getCurrentStatus func() (int, error),
	targetStatus int,
	checkFunc func(int, int) bool,
	interval time.Duration,
	timeout time.Duration,
) error {
	startTime := time.Now()

	for {
		curStatus, err := getCurrentStatus()
		if err != nil {
			return err
		}
		if checkFunc(curStatus, targetStatus) {
			break
		}
		if time.Since(startTime) >= timeout {
			return errors.New("timeout")
		}
		log.Printf("Current status is %d, waiting to %d\n", curStatus, targetStatus)
		time.Sleep(interval * time.Second)
	}

	return nil
}

func ExponentialTimeCheck(
	getCurrentStatus func() (int, error),
	targetStatus int,
	checkFunc func(int, int) bool,
	startInterval time.Duration,
	timeout time.Duration,
) error {
	startTime := time.Now()
	interval := startInterval
	for {
		curStatus, err := getCurrentStatus()
		if err != nil {
			return err
		}
		if checkFunc(curStatus, targetStatus) {
			break
		}
		if time.Since(startTime) >= timeout {
			return errors.New("timeout")
		}
		log.Printf("Current status is %d, waiting for %d\n", curStatus, targetStatus)
		time.Sleep(interval * time.Second)
		interval *= 2
	}

	return nil
}

func DiffTimeCheck(
	getCurrentStatus func() (int, error),
	targetStatus int,
	checkFunc func(int, int) bool,
	diffInterval, maxInterval time.Duration,
	timeout time.Duration,
) error {
	startTime := time.Now()

	for {
		curStatus, err := getCurrentStatus()
		if err != nil {
			return err
		}
		if checkFunc(curStatus, targetStatus) {
			break
		}
		if time.Since(startTime) >= timeout {
			return errors.New("timeout")
		}
		log.Printf("Current status is %d, waiting for %d\n", curStatus, targetStatus)
		time.Sleep(
			min(
				time.Duration((curStatus-targetStatus))*diffInterval,
				maxInterval,
			) * time.Second,
		)
	}

	return nil
}

func SelfAdaptiveTimeCheck(
	getCurrentStatus func() (int, error),
	targetStatus int,
	checkFunc func(int, int) bool,
	startInterval, maxInterval time.Duration,
	timeout time.Duration,
) error {
	startTime := time.Now()
	startInterval = min(startInterval, maxInterval)
	var preDiff, preInterval time.Duration
	for {
		curStatus, err := getCurrentStatus()
		if err != nil {
			return err
		}
		if checkFunc(curStatus, targetStatus) {
			break
		}
		if time.Since(startTime) >= timeout {
			return errors.New("timeout")
		}
		log.Printf("Current status is %d, waiting for %d\n", curStatus, targetStatus)
		if preDiff == 0 {
			time.Sleep(startInterval * time.Second)
			preDiff, preInterval = time.Duration(abs(curStatus-targetStatus)), startInterval
		} else {
			diff := time.Duration(abs(curStatus - targetStatus))
			unitInterval := preInterval / time.Duration(abs(int(diff-preDiff)))
			interval := min(unitInterval*diff, maxInterval)
			time.Sleep(interval * time.Second)
			preDiff, preInterval = diff, interval
		}
	}

	return nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}
