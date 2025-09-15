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

	db.Migrator().DropTable(&User{})
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

func TestApplyPager(t *testing.T) {
	db := setup(t)

	t.Run("pager", func(t *testing.T) {
		var users []User

		pager := &Pager{Number: 2, Size: 1}
		err := db.Scopes(ApplyPager(pager)).Find(&users).Error
		assert.NoError(t, err)
		assert.Len(t, users, 1)
		assert.Equal(t, "Bob", users[0].Name)
	})

	t.Run("nil pager", func(t *testing.T) {
		var users []User

		err := db.Scopes(ApplyPager(nil)).Find(&users).Error
		assert.NoError(t, err)
		assert.Len(t, users, 3)
	})
}

func TestApplySorter(t *testing.T) {
	db := setup(t)

	t.Run("single sorter", func(t *testing.T) {
		var users []User

		sorters := []Sorter{
			{Field: "age", Asc: false},
		}
		err := db.Scopes(ApplySorter(sorters)).Find(&users).Error
		assert.NoError(t, err)
		assert.Len(t, users, 3)
		assert.Equal(t, "Charlie", users[0].Name)
	})

	t.Run("multiple sorters", func(t *testing.T) {
		var users []User

		sorters := []Sorter{
			{Field: "age", Asc: true},
			{Field: "name", Asc: false},
		}
		err := db.Scopes(ApplySorter(sorters)).Find(&users).Error
		assert.NoError(t, err)
		assert.Len(t, users, 3)
		assert.Equal(t, "Alice", users[0].Name)
	})

	t.Run("default sorter", func(t *testing.T) {
		var users []User

		defaultSorter := Sorter{Field: "name", Asc: true}
		err := db.Scopes(ApplySorter([]Sorter{}, defaultSorter)).Find(&users).Error
		assert.NoError(t, err)
		assert.Len(t, users, 3)
		assert.Equal(t, "Alice", users[0].Name)
	})

	t.Run("no sorter", func(t *testing.T) {
		var users []User

		err := db.Scopes(ApplySorter([]Sorter{})).Find(&users).Error
		assert.NoError(t, err)
		assert.Len(t, users, 3)
	})
}

func TestToPascalCase(t *testing.T) {
	assert.Equal(t, "HelloWorld", ToPascalCase("helloWorld"))
	assert.Equal(t, "H", ToPascalCase("h"))
	assert.Empty(t, ToPascalCase(""))
}
