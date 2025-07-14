package lib

import "fmt"

func As(exp string, as string) string {
	return fmt.Sprintf("%s AS %s", exp, as)
}

func Ord(column string, asc bool) string {
	if !asc {
		return column + " DESC"
	}
	return column
}
