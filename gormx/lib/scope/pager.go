package scope

import "gorm.io/gorm"

type Pager struct {
	Number int `json:"number" validate:"required"`
	Size   int `json:"size"   validate:"required"`
}

func ApplyPager(pager *Pager) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pager == nil {
			return db
		}
		return db.
			Offset(pager.Size * (pager.Number - 1)).
			Limit(pager.Size)
	}
}
