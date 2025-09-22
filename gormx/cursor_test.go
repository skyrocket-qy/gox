package gormx_test

import (
	"regexp"
	"testing"

	"github.com/skyrocket-qy/gox/gormx"
	pkgpbv1 "github.com/skyrocket-qy/protos/gen/pkgpb/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// setupDryRunDB creates a new GORM DB instance in DryRun mode for testing SQL generation.
func setupDryRunDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		DryRun: true,
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err)

	return db
}

// normalizeSQL is a helper to make SQL string comparison more reliable.
// It removes table name quotes and collapses whitespace.
func normalizeSQL(sql string) string {
	sql = regexp.MustCompile(`\s+`).ReplaceAllString(sql, " ")
	sql = regexp.MustCompile("`|\"").ReplaceAllString(sql, "")

	return sql
}

type DummyModel struct {
	ID        int
	CreatedAt int
	Score     int
}

func TestApplyCursor(t *testing.T) {
	testCases := []struct {
		name         string
		cursorData   *pkgpbv1.CursorData
		expectedSQL  string
		expectedVars []any
	}{
		{
			name:         "should not apply where clause for empty fields",
			cursorData:   &pkgpbv1.CursorData{},
			expectedSQL:  "SELECT * FROM dummy_models",
			expectedVars: []any{},
		},
		{
			name: "should apply single ascending field",
			cursorData: &pkgpbv1.CursorData{
				Fields: []*pkgpbv1.Field{
					{Col: "id", Val: "10", Asc: true},
				},
			},
			expectedSQL:  "SELECT * FROM dummy_models WHERE (id > ?)",
			expectedVars: []any{"10"},
		},
		{
			name: "should apply single descending field",
			cursorData: &pkgpbv1.CursorData{
				Fields: []*pkgpbv1.Field{
					{Col: "created_at", Val: "12345", Asc: false},
				},
			},
			expectedSQL:  "SELECT * FROM dummy_models WHERE (created_at < ?)",
			expectedVars: []any{"12345"},
		},
		{
			name: "should apply two ascending fields",
			cursorData: &pkgpbv1.CursorData{
				Fields: []*pkgpbv1.Field{
					{Col: "created_at", Val: "12345", Asc: true},
					{Col: "id", Val: "100", Asc: true},
				},
			},
			expectedSQL:  "SELECT * FROM dummy_models WHERE (created_at > ?) OR (created_at = ? AND id > ?)",
			expectedVars: []any{"12345", "12345", "100"},
		},
		{
			name: "should apply two fields with mixed order",
			cursorData: &pkgpbv1.CursorData{
				Fields: []*pkgpbv1.Field{
					{Col: "score", Val: "95", Asc: false},
					{Col: "id", Val: "50", Asc: true},
				},
			},
			expectedSQL:  "SELECT * FROM dummy_models WHERE (score < ?) OR (score = ? AND id > ?)",
			expectedVars: []any{"95", "95", "50"},
		},
		{
			name: "should handle three fields",
			cursorData: &pkgpbv1.CursorData{
				Fields: []*pkgpbv1.Field{
					{Col: "score", Val: "100", Asc: false},
					{Col: "created_at", Val: "12345", Asc: true},
					{Col: "id", Val: "200", Asc: true},
				},
			},
			expectedSQL: "SELECT * FROM dummy_models WHERE (score < ?) " +
				"OR (score = ? AND created_at > ?) " +
				"OR (score = ? AND created_at = ? AND id > ?)",
			expectedVars: []any{"100", "100", "12345", "100", "12345", "200"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := setupDryRunDB(t)

			var results []DummyModel

			tx := db.Scopes(gormx.ApplyCursor(tc.cursorData))
			tx.Find(&results)

			stmt := tx.Statement
			assert.Equal(t, normalizeSQL(tc.expectedSQL), normalizeSQL(stmt.SQL.String()))
			assert.Equal(t, tc.expectedVars, stmt.Vars)
		})
	}
}

func TestApplyCursor_WithNilCursor(t *testing.T) {
	t.Run("should not apply where clause or panic with nil cursor", func(t *testing.T) {
		db := setupDryRunDB(t)

		var results []DummyModel

		// The scope is applied, but the query isn't run until a finalizer method like Find() is
		// called.
		tx := db.Scopes(gormx.ApplyCursor(nil)).Model(&DummyModel{})

		// Now run the finalizer
		tx.Find(&results)

		stmt := tx.Statement
		assert.Equal(t, "SELECT * FROM dummy_models", normalizeSQL(stmt.SQL.String()))
		assert.Empty(t, stmt.Vars)
	})
}
