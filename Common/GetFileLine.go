package Common

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/google/uuid"
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

type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}

// Log error stack into log server with matched uuid
func WithTrace(err error) error {
	// Fool-proof
	if err == nil {
		return &CustomError{}
	}

	// if != customError means not logged, so log it
	if _, ok := err.(*CustomError); !ok {
		ud := uuid.New().String()
		callStkMsg := getCallStack(3)

		fmt.Sprintf("uuid: %s\n", ud)
		fmt.Sprintf("error: %s\n", err.Error())
		fmt.Sprintf("stack: %s\n", callStkMsg)

		return &CustomError{
			msg: fmt.Sprintf("uuid: %s \nerror: %s", ud, err.Error()),
		}
	}

	return err
}

func getCallStack(callerSkip ...int) (stkMsg string) {
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
