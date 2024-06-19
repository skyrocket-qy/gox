package Common

import (
	"errors"
	"fmt"
	"runtime"
)

// GetFileAndLine default callerSkip is 1
func GetFileAndLine(callerSkip ...int) (string, error) {
	skip := 1
	if len(callerSkip) == 0 {
		skip = callerSkip[0]
	}
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", errors.New("find failed")
	}

	return fmt.Sprintf("%s %d", file, line), nil
}
