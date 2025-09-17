package gormx

import (
	"fmt"
	"slices"
	"strings"

	"github.com/skyrocket-qy/erx"
	"github.com/skyrocket-qy/gox/errcode"
	pkgpbv1 "github.com/skyrocket-qy/protos/gen/pkgpb/v1"
	"gorm.io/gorm"
)

var opTemplate = map[pkgpbv1.Operator]string{
	pkgpbv1.Operator_EQ:      "%s = ?",
	pkgpbv1.Operator_GT:      "%s > ?",
	pkgpbv1.Operator_GTE:     "%s >= ?",
	pkgpbv1.Operator_LT:      "%s < ?",
	pkgpbv1.Operator_LTE:     "%s <= ?",
	pkgpbv1.Operator_BETWEEN: "%s BETWEEN ? AND ?",
}

func ApplyFilter(filters []*pkgpbv1.Filter, validFields []string, filterExprs map[string][]string) (
	func(db *gorm.DB) *gorm.DB, error,
) {
	visited := map[string]struct{}{}

	for _, ft := range filters {
		if ft.GetField() == "" {
			return nil, erx.New(errcode.ErrBadRequest, "field is empty")
		}

		if !slices.Contains(validFields, ft.GetField()) {
			return nil, erx.Newf(errcode.ErrBadRequest, "invalid field: %v", ft.GetField())
		}

		if _, ok := visited[ft.GetField()]; ok {
			return nil, erx.Newf(errcode.ErrBadRequest, "duplicate field: %v", ft.GetField())
		}

		visited[ft.GetField()] = struct{}{}

		switch ft.GetOp() {
		case pkgpbv1.Operator_EQ,
			pkgpbv1.Operator_GT,
			pkgpbv1.Operator_GTE,
			pkgpbv1.Operator_LT,
			pkgpbv1.Operator_LTE:
			if len(ft.GetValues()) != 1 {
				return nil, erx.Newf(
					errcode.ErrBadRequest,
					"%v filter requires one value",
					ft.GetOp(),
				)
			}
		case pkgpbv1.Operator_BETWEEN:
			if len(ft.GetValues()) != 2 {
				return nil, erx.New(errcode.ErrBadRequest, "between filter requires two values")
			}
			// Optional: check range
			if ft.GetValues()[1] <= ft.GetValues()[0] {
				return nil, erx.New(errcode.ErrBadRequest, "between: second value must be >= first")
			}
		default:
			return nil, erx.Newf(errcode.ErrBadRequest, "unsupported operator: %v", ft.GetOp())
		}
	}

	return func(db *gorm.DB) *gorm.DB {
		joinVisited := map[string]struct{}{}

		for _, ft := range filters {
			if expr, ok := filterExprs[ft.GetField()]; ok {
				for _, join := range expr {
					if _, seen := joinVisited[join]; !seen {
						db = db.Joins(join)
						joinVisited[join] = struct{}{}
					}
				}
			}

			column := QuoteIfNeeded(ft.GetField())
			tmpl := opTemplate[ft.GetOp()]

			switch ft.GetOp() {
			case pkgpbv1.Operator_BETWEEN:
				if len(ft.GetValues()) >= 2 {
					db = db.Where(fmt.Sprintf(tmpl, column), ft.GetValues()[0], ft.GetValues()[1])
				}
			default:
				if len(ft.GetValues()) >= 1 {
					db = db.Where(fmt.Sprintf(tmpl, column), ft.GetValues()[0])
				}
			}
		}

		return db
	}, nil
}

func QuoteIfNeeded(field string) string {
	if field == "" {
		return "``"
	}

	if strings.Contains(field, ".") {
		parts := strings.Split(field, ".")
		for i, p := range parts {
			parts[i] = "`" + p + "`"
		}

		return strings.Join(parts, ".")
	}

	return "`" + field + "`"
}
