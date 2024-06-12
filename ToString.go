package GoUtils

import (
	"fmt"
	"strconv"
)

func ToString[T any](v T) string {
	switch val := any(v).(type) {
	case string:
		return val
	case bool:
		return strconv.FormatBool(val)
	case int:
		return strconv.Itoa(val)
	case int8, int16, int32, int64:
		return strconv.FormatInt(val.(int64), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(val.(uint64), 10)
	case float32, float64:
		return strconv.FormatFloat(val.(float64), 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", val)
	}
}
