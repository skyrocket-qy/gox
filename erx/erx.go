package erx

import (
	"errors"
	"fmt"
	"runtime"
)

// W wraps the given error with a call stack and optional additional context.
func W(err error, msgs ...string) *CtxErr {
	if err == nil {
		return nil
	}

	var ctxErr *CtxErr
	if !errors.As(err, &ctxErr) {
		ctxErr = &CtxErr{
			CallerInfos: getCallStack(3),
			Code:        ErrToCode(err),
		}

		firstCaller := ctxErr.CallerInfos[0]
		firstCaller.Msg = fmt.Sprintf("3rd party error: %v", err.Error())

		if len(msgs) > 0 {
			firstCaller.Msg += " " + msgs[0]
		}
	} else {
		if len(msgs) > 0 {
			msg := msgs[0]
			pc, file, line, ok := runtime.Caller(2)
			if !ok {
				return ctxErr
			}

			funcName := runtime.FuncForPC(pc).Name()

			for i := range ctxErr.CallerInfos {
				ci := &ctxErr.CallerInfos[i]
				if ci.Function == funcName && ci.File == file && ci.Line == line {
					ci.Msg += " " + msg
					break
				}
			}
		}
	}

	return ctxErr
}

func New(code Code, msgs ...string) *CtxErr {
	ctxErr := &CtxErr{
		CallerInfos: getCallStack(2),
		Code:        code,
	}

	firstCaller := ctxErr.CallerInfos[0]
	if len(msgs) > 0 {
		firstCaller.Msg += " " + msgs[0]
	}

	return ctxErr
}

func getCallStack(callerSkip ...int) (callerInfos []CallerInfo) {
	pc := make([]uintptr, MaxCallStackSize)

	skip := 2
	if len(callerSkip) > 0 {
		skip = callerSkip[0]
	}

	n := runtime.Callers(skip, pc)

	frames := runtime.CallersFrames(pc[:n])

	for {
		frame, more := frames.Next()
		callerInfos = append(callerInfos, CallerInfo{
			Function: frame.Function,
			File:     frame.File,
			Line:     frame.Line,
		})

		if !more {
			break
		}
	}

	return callerInfos
}
