package gormx

import (
	"github.com/skyrocketOoO/erx"
	"gorm.io/gorm"
)

type Err string

func (e Err) Code() string {
	return string(e)
}

func (e Err) Msg() string {
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
