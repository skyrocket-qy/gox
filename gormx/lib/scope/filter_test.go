package scope

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID   uint
	Name string
	Age  int
}

func setup(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(t, err)
	sqlDB, err := db.DB()
	assert.NoError(t, err)
	sqlDB.SetMaxOpenConns(1)

	err = db.AutoMigrate(&User{})
	assert.NoError(t, err)

	users := []User{
		{Name: "Alice", Age: 20},
		{Name: "Bob", Age: 21},
		{Name: "Charlie", Age: 22},
	}
	for _, user := range users {
		err := db.Create(&user).Error
		assert.NoError(t, err)
	}

	return db
}

func TestApplyFilter(t *testing.T) {
	db := setup(t)

	t.Run("exact match", func(t *testing.T) {
		var users []User

		filters := []Filter{
			{Field: "name", Fuzzy: false, Value: "Alice"},
		}
		err := db.Scopes(ApplyFilter(db, filters)).Find(&users).Error
		assert.NoError(t, err)
		assert.Len(t, users, 1)
		assert.Equal(t, "Alice", users[0].Name)
	})

	t.Run("fuzzy match", func(t *testing.T) {
		var users []User

		filters := []Filter{
			{Field: "name", Fuzzy: true, Value: "li"},
		}
		err := db.Scopes(ApplyFilter(db, filters)).Find(&users).Error
		assert.NoError(t, err)
		assert.Len(t, users, 2)
		assert.Contains(t, []string{users[0].Name, users[1].Name}, "Alice")
		assert.Contains(t, []string{users[0].Name, users[1].Name}, "Charlie")
	})

	t.Run("multiple filters", func(t *testing.T) {
		var users []User

		filters := []Filter{
			{Field: "name", Fuzzy: true, Value: "li"},
			{Field: "age", Fuzzy: false, Value: "20"},
		}
		err := db.Scopes(ApplyFilter(db, filters)).Find(&users).Error
		assert.NoError(t, err)
		assert.Len(t, users, 1)
		assert.Equal(t, "Alice", users[0].Name)
	})
}
