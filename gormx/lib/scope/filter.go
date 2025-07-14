package scope

import (
	ope "github.com/skyrocketOoO/gox/gormx/lib/operator"
	wh "github.com/skyrocketOoO/gox/gormx/lib/where"
	"gorm.io/gorm"
)

type Filter struct {
	Field string `json:"field" validate:"required"`
	Fuzzy bool   `json:"fuzzy" validate:"required"`
	Value string `json:"value" validate:"required"`
}

func ApplyFilter(db *gorm.DB, filters []Filter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, ft := range filters {
			if ft.Fuzzy {
				db = db.Where(wh.B(ft.Field, ope.Like), "%"+ft.Value+"%")
			} else {
				db = db.Where(wh.B(ft.Field, ope.Eq), ft.Value)
			}
		}

		return db
	}
}
