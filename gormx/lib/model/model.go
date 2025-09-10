package model

import "gorm.io/gorm"

type Base struct {
	gorm.Model

	ID string `gorm:"type:char(36);primaryKey"`
}
