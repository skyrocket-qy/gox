package gormx

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

	// Test case 1: record exists
	mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `users` WHERE name = ?")).
		WithArgs("test").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
	err = CheckExist(gormDB, &User{}, "name = ?", "test")
	if err == nil {
		t.Error("expected an error, but got nil")
	}

	// Test case 2: record does not exist
	mock.ExpectQuery(regexp.QuoteMeta("SELECT count(*) FROM `users` WHERE name = ?")).
		WithArgs("test").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
	err = CheckExist(gormDB, &User{}, "name = ?", "test")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestErr_Str(t *testing.T) {
	if ErrDuplicate.Str() != "409.0000" {
		t.Errorf("ErrDuplicate.Str() failed, got %s", ErrDuplicate.Str())
	}
}
