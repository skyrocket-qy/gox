package columnname_test

import (
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/skyrocket-qy/gox/gormx/columnname"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestToCamel(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "hello_world",
			want:  "HelloWorld",
		},
		{
			input: "foo_bar_baz",
			want:  "FooBarBaz",
		},
		{
			input: "single",
			want:  "Single",
		},
		{
			input: "",
			want:  "",
		},
		{
			input: "a_b_c",
			want:  "ABC",
		},
		{
			input: "_a",
			want:  "A",
		},
		{
			input: "a_",
			want:  "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := columnname.ToCamel(tt.input); got != tt.want {
				t.Errorf("ToCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetColumns(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	require.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT SCHEMA_NAME from Information_schema.SCHEMATA")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"SCHEMA_NAME"}).AddRow("test"))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` LIMIT ?")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(nil))

	rows := sqlmock.NewRows([]string{
		"column_name", "column_default", "is_nullable", "data_type",
		"character_maximum_length", "column_type", "column_key", "extra",
		"column_comment", "numeric_precision", "numeric_scale", "datetime_precision",
	}).
		AddRow("id", nil, false, "int", nil, "int(11)", "PRI", "auto_increment", "", nil, nil, nil).
		AddRow("name", nil, true, "varchar", 255, "varchar(255)", "", "", "", nil, nil, nil)

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT column_name, column_default, is_nullable = 'YES', data_type, "+
			"character_maximum_length, column_type, column_key, extra, column_comment, "+
			"numeric_precision, numeric_scale , datetime_precision "+
			"FROM information_schema.columns WHERE table_schema = ? AND table_name = ? "+
			"ORDER BY ORDINAL_POSITION")).
		WithArgs("test", "users").
		WillReturnRows(rows)

	columns, err := columnname.GetColumns(gormDB, "users")
	require.NoError(t, err)
	assert.Equal(t, []string{"id", "name"}, columns)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGenTableColumnNamesCode(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	require.NoError(t, err)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT SCHEMA_NAME from Information_schema.SCHEMATA")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"SCHEMA_NAME"}).AddRow("test"))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` LIMIT ?")).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows(nil))
	rows := sqlmock.NewRows([]string{
		"column_name", "column_default", "is_nullable", "data_type",
		"character_maximum_length", "column_type", "column_key", "extra",
		"column_comment", "numeric_precision", "numeric_scale", "datetime_precision",
	}).
		AddRow("id", nil, false, "int", nil, "int(11)", "PRI", "auto_increment", "", nil, nil, nil).
		AddRow("user_name", nil, true, "varchar", 255, "varchar(255)", "", "", "", nil, nil, nil)

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT column_name, column_default, is_nullable = 'YES', data_type, "+
			"character_maximum_length, column_type, column_key, extra, column_comment, "+
			"numeric_precision, numeric_scale , datetime_precision "+
			"FROM information_schema.columns WHERE table_schema = ? AND table_name = ? "+
			"ORDER BY ORDINAL_POSITION")).
		WithArgs("test", "users").
		WillReturnRows(rows)

	path := "./col/users.go"
	err = columnname.GenTableColumnNamesCode(gormDB, []string{"users"}, path)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())

	// cleanup
	err = os.RemoveAll("./col")
	require.NoError(t, err)
}

func TestGenTableColumnNamesCode_ToCamel(t *testing.T) {
	output := columnname.ToCamel("hello_world")
	assert.Equal(t, "HelloWorld", output)
}
