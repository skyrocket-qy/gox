package gormx

import (
	"log"
	"strings"

	"github.com/skyrocket-qy/erx"
	pkgpbv1 "github.com/skyrocket-qy/protos/gen/pkgpb/v1"
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

func ApplyCursor(c *pkgpbv1.CursorData) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if c == nil || len(c.GetFields()) == 0 {
			return db
		}

		var (
			orConditions []string
			orArgs       [][]any
		)

		for i := range c.GetFields() {
			var (
				condParts []string
				args      []any
			)

			// Equal conditions for previous fields

			for j := range i {
				condParts = append(condParts, c.GetFields()[j].GetCol()+" = ?")
				args = append(args, c.GetFields()[j].GetVal())
			}

			// Comparison condition for current field
			op := ">"
			if !c.GetFields()[i].GetAsc() {
				op = "<"
			}

			condParts = append(condParts, c.GetFields()[i].GetCol()+" "+op+" ?")
			args = append(args, c.GetFields()[i].GetVal())

			orConditions = append(orConditions, "("+strings.Join(condParts, " AND ")+")")
			orArgs = append(orArgs, args)
		}

		// Combine all OR conditions
		fullWhere := strings.Join(orConditions, " OR ")

		var allArgs []any
		for _, args := range orArgs {
			allArgs = append(allArgs, args...)
		}

		log.Print(fullWhere)
		log.Print(allArgs)

		return db.Where(fullWhere, allArgs...)
	}
}
