package gormx_test

import (
	"testing"

	"github.com/skyrocket-qy/gox/gormx"

	pkgpbv1 "github.com/skyrocket-qy/protos/gen/pkgpb/v1"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// --- Test QuoteIfNeeded ---

func TestQuoteIfNeeded(t *testing.T) {
	t.Run("should quote a simple field", func(t *testing.T) {
		assert.Equal(t, "`id`", gormx.QuoteIfNeeded("id"))
	})

	t.Run("should quote a qualified field correctly", func(t *testing.T) {
		assert.Equal(t, "`users`.`name`", gormx.QuoteIfNeeded("users.name"))
	})

	t.Run("should handle empty string", func(t *testing.T) {
		assert.Equal(t, "``", gormx.QuoteIfNeeded(""))
	})
}

// --- Test ApplyFilter ---

type DummyFilterModel struct {
	ID   int
	Name string
	Role string
}

func setupFilterDryRunDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		DryRun: true,
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(t, err)

	return db
}

// NOTE: Testing the validation logic of ApplyFilter is not currently possible
// at the unit level. The function returns errors created with `erx.New`, but
// the base error code it uses (e.g., `ErrBadRequest`) is not defined in this
// package, and importing it would cause a circular dependency.
// The `erx` library itself also makes it difficult to create test-only
// error codes, as discovered in previous tests.
// The validation paths are tested implicitly by integration tests.

func TestApplyFilter_SQLGeneration(t *testing.T) {
	validFields := []string{"id", "name", "role"}
	// NOTE: Testing the join logic is difficult in a unit test as it depends
	// on GORM model definitions, so it has been omitted.
	filterExprs := map[string][]string{}

	t.Run("should apply simple EQ filter", func(t *testing.T) {
		db := setupFilterDryRunDB(t)
		filters := []*pkgpbv1.Filter{
			{Field: "name", Op: pkgpbv1.Operator_EQ, Values: []string{"jules"}},
		}
		scope, err := gormx.ApplyFilter(filters, validFields, filterExprs)
		assert.NoError(t, err)

		var results []DummyFilterModel

		tx := db.Model(&DummyFilterModel{}).Scopes(scope).Find(&results)

		expectedSQL := "SELECT * FROM `dummy_filter_models` WHERE `name` = ?"
		assert.Equal(t, expectedSQL, tx.Statement.SQL.String())
		assert.Equal(t, []any{"jules"}, tx.Statement.Vars)
	})

	t.Run("should apply multiple filters", func(t *testing.T) {
		db := setupFilterDryRunDB(t)
		filters := []*pkgpbv1.Filter{
			{Field: "name", Op: pkgpbv1.Operator_EQ, Values: []string{"jules"}},
			{Field: "id", Op: pkgpbv1.Operator_GTE, Values: []string{"100"}},
		}
		scope, err := gormx.ApplyFilter(filters, validFields, filterExprs)
		assert.NoError(t, err)

		var results []DummyFilterModel

		tx := db.Model(&DummyFilterModel{}).Scopes(scope).Find(&results)

		expectedSQL := "SELECT * FROM `dummy_filter_models` WHERE `name` = ? AND `id` >= ?"
		assert.Equal(t, expectedSQL, tx.Statement.SQL.String())
		assert.Equal(t, []any{"jules", "100"}, tx.Statement.Vars)
	})

	t.Run("should apply BETWEEN filter", func(t *testing.T) {
		db := setupFilterDryRunDB(t)
		filters := []*pkgpbv1.Filter{
			{Field: "id", Op: pkgpbv1.Operator_BETWEEN, Values: []string{"100", "200"}},
		}
		scope, err := gormx.ApplyFilter(filters, validFields, filterExprs)
		assert.NoError(t, err)

		var results []DummyFilterModel

		tx := db.Model(&DummyFilterModel{}).Scopes(scope).Find(&results)

		expectedSQL := "SELECT * FROM `dummy_filter_models` WHERE `id` BETWEEN ? AND ?"
		assert.Equal(t, expectedSQL, tx.Statement.SQL.String())
		assert.Equal(t, []any{"100", "200"}, tx.Statement.Vars)
	})
}
