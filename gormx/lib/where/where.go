package wh

import (
	"fmt"

	ope "github.com/skyrocketOoO/gox/gormx/lib/operator"
)

func B(field string, oper string) string {
	if oper == ope.Bt || oper == ope.NBt {
		return fmt.Sprintf("%s %s ? AND ?", field, oper)
	}
	return fmt.Sprintf("%s %s ?", field, oper)
}

// used for subquery
func BSub(field string, oper string) string {
	if oper == ope.Bt || oper == ope.NBt {
		return fmt.Sprintf("%s %s (?) AND (?)", field, oper)
	}
	return fmt.Sprintf("%s %s (?)", field, oper)
}

func Desc(column string) string {
	return column + " DESC"
}

func Asc(column string) string {
	return column
}
