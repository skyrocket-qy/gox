package gormx

import (
	"github.com/skyrocket-qy/erx"
	"gorm.io/gorm"
)

type Err string

func (e Err) Str() string {
	return string(e)
}

const (
	ErrDuplicate Err = "409.0000"
)

func CheckExist(db *gorm.DB, model any, query string, args ...any) error {
	var count int64
	if err := db.Model(model).Where(query, args...).Count(&count).Error; err != nil {
		return erx.W(err)
	}

	if count > 0 {
		return erx.New(ErrDuplicate)
	}

	return nil
}

// func ApplyFilter(filters []*pkgpbv1.Filter, validFields map[string]struct{}) (
// 	func(db *gorm.DB) *gorm.DB, error,
// ) {
// 	for _, ft := range filters {
// 		if ft.Field == "" {
// 			return nil, errors.New("field is empty")
// 		}

// 		if _, ok := validFields[ft.Field]; !ok {
// 			return nil, fmt.Errorf("invalid field: %v", ft.Field)
// 		}

// 		switch ft.Op {
// 		case pkgpbv1.FilterOperator_FILTER_OPERATOR_EQ,
// 			pkgpbv1.FilterOperator_FILTER_OPERATOR_GT,
// 			pkgpbv1.FilterOperator_FILTER_OPERATOR_GTE,
// 			pkgpbv1.FilterOperator_FILTER_OPERATOR_LT,
// 			pkgpbv1.FilterOperator_FILTER_OPERATOR_LTE:
// 			if len(ft.Values) != 1 {
// 				return nil, fmt.Errorf("%v filter requires one value", ft.Op)
// 			}
// 		case pkgpbv1.FilterOperator_FILTER_OPERATOR_BETWEEN:
// 			if len(ft.Values) != 2 {
// 				return nil, errors.New("between filter requires two values")
// 			}
// 			// Optional: check range
// 			if ft.Values[1] <= ft.Values[0] {
// 				return nil, errors.New("between: second value must be >= first")
// 			}
// 		default:
// 			return nil, fmt.Errorf("unsupported operator: %v", ft.Op)
// 		}
// 	}

// 	return func(db *gorm.DB) *gorm.DB {
// 		for _, ft := range filters {
// 			column := ft.Field // ⚠️ Ensure it's a valid column name

// 			switch ft.Op {
// 			case pkgpbv1.FilterOperator_FILTER_OPERATOR_EQ:
// 				db = db.Where(fmt.Sprintf("%s = ?", column), ft.Values[0])
// 			case pkgpbv1.FilterOperator_FILTER_OPERATOR_GT:
// 				db = db.Where(fmt.Sprintf("%s > ?", column), ft.Values[0])
// 			case pkgpbv1.FilterOperator_FILTER_OPERATOR_GTE:
// 				db = db.Where(fmt.Sprintf("%s >= ?", column), ft.Values[0])
// 			case pkgpbv1.FilterOperator_FILTER_OPERATOR_LT:
// 				db = db.Where(fmt.Sprintf("%s < ?", column), ft.Values[0])
// 			case pkgpbv1.FilterOperator_FILTER_OPERATOR_LTE:
// 				db = db.Where(fmt.Sprintf("%s <= ?", column), ft.Values[0])
// 			case pkgpbv1.FilterOperator_FILTER_OPERATOR_BETWEEN:
// 				db = db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", column), ft.Values[0], ft.Values[1])
// 			}
// 		}
// 		return db
// 	}, nil
// }
