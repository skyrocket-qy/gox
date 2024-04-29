package logger

import (
	"errors"
	"fmt"
	"runtime"
)

// GetFileAndLine default callerSkip is 1
func GetFileAndLine(callerSkip int) (string, error) {
	_, file, line, ok := runtime.Caller(callerSkip)
	if !ok {
		return "", errors.New("find failed")
	}

	return fmt.Sprintf("%s %d", file, line), nil
}

// func main() {
// 	fmt.Println(GetFileAndLine(1))
// }
