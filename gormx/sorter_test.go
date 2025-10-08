package gormx

import (
	"testing"

	pkgpbv1 "github.com/skyrocket-qy/protos/gen/pkgpb/v1"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestToPascalCase(t *testing.T) {
	assert.Equal(t, "Id", ToPascalCase("id"))
	assert.Equal(t, "Name", ToPascalCase("name"))
	assert.Equal(t, "UserID", ToPascalCase("userID"))
	assert.Empty(t, ToPascalCase(""))
	assert.Equal(t, "A", ToPascalCase("a"))
}

func setupSorterDryRunDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		DryRun: true,
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(t, err)

	return db
}

type DummySorterModel struct {
	ID   int
	Name string
}

func TestApplySorter(t *testing.T) {
	t.Run("should do nothing with no sorters", func(t *testing.T) {
		db := setupSorterDryRunDB(t)
		// Call with only one argument, so the variadic part is empty
		scope := ApplySorter(nil)
		tx := db.Model(&DummySorterModel{}).Scopes(scope).Find(&[]DummySorterModel{})
		assert.NotContains(t, tx.Statement.SQL.String(), "ORDER BY")
	})

	t.Run("should do nothing with nil default sorter", func(t *testing.T) {
		db := setupSorterDryRunDB(t)
		// This passes a slice containing a single nil element to the variadic parameter
		scope := ApplySorter(nil, nil)
		tx := db.Model(&DummySorterModel{}).Scopes(scope).Find(&[]DummySorterModel{})
		// The fixed function should handle this gracefully.
		assert.NotContains(t, tx.Statement.SQL.String(), "ORDER BY")
	})

	t.Run("should apply default sorter when primary is empty", func(t *testing.T) {
		db := setupSorterDryRunDB(t)
		defaultSorter := &pkgpbv1.Sorter{Field: "id", Asc: false}
		scope := ApplySorter(nil, defaultSorter)
		tx := db.Model(&DummySorterModel{}).Scopes(scope).Find(&[]DummySorterModel{})
		expectedSQL := "SELECT * FROM `dummy_sorter_models` ORDER BY id DESC"
		assert.Equal(t, expectedSQL, tx.Statement.SQL.String())
	})

	t.Run("should apply a single primary sorter", func(t *testing.T) {
		db := setupSorterDryRunDB(t)
		sorters := []*pkgpbv1.Sorter{{Field: "name", Asc: true}}
		scope := ApplySorter(sorters)
		tx := db.Model(&DummySorterModel{}).Scopes(scope).Find(&[]DummySorterModel{})
		// Note the PascalCase conversion on the field
		expectedSQL := "SELECT * FROM `dummy_sorter_models` ORDER BY Name"
		assert.Equal(t, expectedSQL, tx.Statement.SQL.String())
	})

	t.Run("should apply multiple primary sorters", func(t *testing.T) {
		db := setupSorterDryRunDB(t)
		sorters := []*pkgpbv1.Sorter{
			{Field: "name", Asc: false},
			{Field: "id", Asc: true},
		}
		scope := ApplySorter(sorters)
		tx := db.Model(&DummySorterModel{}).Scopes(scope).Find(&[]DummySorterModel{})
		expectedSQL := "SELECT * FROM `dummy_sorter_models` ORDER BY Name DESC,Id"
		assert.Equal(t, expectedSQL, tx.Statement.SQL.String())
	})

	t.Run("should ignore default sorter if primary sorters exist", func(t *testing.T) {
		db := setupSorterDryRunDB(t)
		primarySorters := []*pkgpbv1.Sorter{{Field: "name", Asc: true}}
		defaultSorter := &pkgpbv1.Sorter{Field: "id", Asc: false}
		scope := ApplySorter(primarySorters, defaultSorter)
		tx := db.Model(&DummySorterModel{}).Scopes(scope).Find(&[]DummySorterModel{})
		expectedSQL := "SELECT * FROM `dummy_sorter_models` ORDER BY Name"
		assert.Equal(t, expectedSQL, tx.Statement.SQL.String())
	})
}
