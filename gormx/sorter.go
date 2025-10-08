package gormx

import (
	"strings"

	pkgpbv1 "github.com/skyrocket-qy/protos/gen/pkgpb/v1"
	"gorm.io/gorm"
)

func ApplySorter(
	seqSorters []*pkgpbv1.Sorter,
	dfSort ...*pkgpbv1.Sorter,
) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(seqSorters) == 0 {
			// Use default sorter only if it exists and is not nil
			if len(dfSort) == 0 || dfSort[0] == nil {
				return db
			}

			df := dfSort[0]

			expr := df.GetField()
			if !df.GetAsc() {
				expr += " DESC"
			}

			return db.Order(expr)
		}

		for _, sorter := range seqSorters {
			expr := ToPascalCase(sorter.GetField())
			if !sorter.GetAsc() {
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
