package scope

import (
	"strings"

	"gorm.io/gorm"
)

type Sorter struct {
	Field string `json:"field" validate:"required"`
	Asc   bool   `json:"asc"   validate:"required"`
}

func ApplySorter(seqSorters []Sorter, dfSort ...Sorter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(seqSorters) == 0 {
			if len(dfSort) == 0 {
				return db
			}

			df := dfSort[0]
			expr := df.Field
			if !df.Asc {
				expr += " DESC"
			}
			return db.Order(expr)
		}

		for _, sorter := range seqSorters {
			expr := ToPascalCase(sorter.Field)
			if !sorter.Asc {
				expr += " DESC"
			}
			db = db.Order(expr)
		}
		return db
	}
}

func ToPascalCase(input string) string {
	if len(input) == 0 {
		return ""
	}

	// Capitalize the first character
	result := strings.ToUpper(string(input[0])) + input[1:]

	return result
}
