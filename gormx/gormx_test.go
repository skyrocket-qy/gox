package gormx_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/skyrocket-qy/gox/gormx"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Name string
}

func (User) TableName() string {
	return "users"
}

func TestCheckExist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	t.Run("record exists", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `users` WHERE name = ?")).
			WithArgs("test").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		err := gormx.CheckExist(gormDB, &User{}, "name = ?", "test")
		assert.Error(t, err)
	})

	t.Run("record does not exist", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `users` WHERE name = ?")).
			WithArgs("test").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

		err := gormx.CheckExist(gormDB, &User{}, "name = ?", "test")
		assert.NoError(t, err)
	})

	t.Run("db error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `users` WHERE name = ?")).
			WithArgs("test").
			WillReturnError(errors.New("db error"))

		err := gormx.CheckExist(gormDB, &User{}, "name = ?", "test")
		assert.Error(t, err)
	})
}

func TestErr_Str(t *testing.T) {
	if gormx.ErrDuplicate.Str() != "409.0000" {
		t.Errorf("ErrDuplicate.Str() failed, got %s", gormx.ErrDuplicate.Str())
	}
}
