package Common

import (
	"errors"
	"fmt"
	"runtime"
)

// GetFileAndLine default callerSkip is 1
func GetFileAndLine(callerSkip ...int) (lineInfo string, err error) {
	skip := 1
	if len(callerSkip) > 0 {
		skip = callerSkip[0]
	}
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", errors.New("find failed")
	}

	return fmt.Sprintf("%s %d", file, line), nil
}

func GetCallStack(callerSkip ...int) (stkMsg string) {
	pc := make([]uintptr, 10)
	skip := 2
	if len(callerSkip) > 0 {
		skip = callerSkip[0]
	}
	n := runtime.Callers(skip, pc)

	frames := runtime.CallersFrames(pc[:n])

	stkMsg = "Call stack:"

	for {
		frame, more := frames.Next()
		stkMsg += fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}
	return stkMsg
}
