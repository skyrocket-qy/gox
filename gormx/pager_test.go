package gormx

import (
	"testing"

	pkgpbv1 "github.com/skyrocket-qy/protos/gen/pkgpb/v1"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupPagerDryRunDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		DryRun: true,
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(t, err)

	return db
}

type DummyPagerModel struct {
	ID int
}

func TestApplyPager(t *testing.T) {
	testCases := []struct {
		name         string
		pager        *pkgpbv1.Pager
		expectedSQL  string
		expectedVars []any
	}{
		{
			name:         "should do nothing for nil pager",
			pager:        nil,
			expectedSQL:  "SELECT * FROM `dummy_pager_models`",
			expectedVars: []any{},
		},
		{
			name: "should apply limit and offset for first page",
			pager: &pkgpbv1.Pager{
				Size:   10,
				Number: 1,
			},
			expectedSQL:  "SELECT * FROM `dummy_pager_models` LIMIT 10",
			expectedVars: []any{},
		},
		{
			name: "should apply limit and offset for second page",
			pager: &pkgpbv1.Pager{
				Size:   25,
				Number: 2,
			},
			expectedSQL:  "SELECT * FROM `dummy_pager_models` LIMIT 25 OFFSET 25",
			expectedVars: []any{},
		},
		{
			name: "should handle zero size",
			pager: &pkgpbv1.Pager{
				Size:   0,
				Number: 5,
			},
			// GORM correctly generates LIMIT 0
			expectedSQL:  "SELECT * FROM `dummy_pager_models` LIMIT 0",
			expectedVars: []any{},
		},
		{
			name: "should handle zero page number",
			pager: &pkgpbv1.Pager{
				Size:   10,
				Number: 0,
			},
			// number-1 becomes -1. Offset(-1) is ignored by GORM.
			expectedSQL:  "SELECT * FROM `dummy_pager_models` LIMIT 10",
			expectedVars: []any{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := setupPagerDryRunDB(t)

			var results []DummyPagerModel

			tx := db.Model(&DummyPagerModel{}).Scopes(ApplyPager(tc.pager)).Find(&results)

			assert.Equal(t, tc.expectedSQL, tx.Statement.SQL.String())
			// Note: GORM puts limit and offset directly into the SQL string for DryRun,
			// so the Vars slice will be empty for these tests.
		})
	}
}
