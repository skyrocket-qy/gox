package excel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToExcel1D(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name     string
		table    [][]string
		wantKey  string
		wantVal  string
		wantData map[K]V
		IsErr    bool
	}

	tests := []testCase[string, int]{
		{
			name:    "normal case",
			table:   [][]string{{"Word", "Count"}, {"hello", "1"}, {"world", "2"}},
			wantKey: "Word",
			wantVal: "Count",
			wantData: map[string]int{
				"hello": 1,
				"world": 2,
			},
			IsErr: false,
		},
		{
			name:  "empty table",
			table: [][]string{},
			IsErr: true,
		},
		{
			name:  "wrong column size",
			table: [][]string{{"Key"}, {"a", "1"}},
			IsErr: true,
		},
		{
			name:  "no data rows",
			table: [][]string{{"Key", "Val"}},
			IsErr: true,
		},
		{
			name:  "bad int conversion",
			table: [][]string{{"Key", "Val"}, {"a", "bad-int"}},
			IsErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			keyName, valName, data, err := ToExcel1D[string, int](tc.table)

			if !tc.IsErr {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
			assert.Equal(t, tc.wantKey, keyName)
			assert.Equal(t, tc.wantVal, valName)
			assert.Equal(t, tc.wantData, data)
		})
	}
}

func TestToColsList(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name  string
		table [][]string
		want  map[K][]V
		isErr bool
	}

	tests := []testCase[string, int]{
		{
			name: "normal case - 2 cols",
			table: [][]string{
				{"col1", "col2"},
				{"1", "2"},
				{"3", "4"},
			},
			want: map[string][]int{
				"col1": {1, 3},
				"col2": {2, 4},
			},
			isErr: false,
		},
		{
			name: "diff len cols",
			table: [][]string{
				{"col1", "col2"},
				{"1", "2"},
				{"3", "4"},
				{"", "5"},
			},
			want: map[string][]int{
				"col1": {1, 3},
				"col2": {2, 4, 5},
			},
			isErr: false,
		},
		{
			name:  "empty table",
			table: [][]string{},
			isErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ToColsList[string, int](tc.table)

			if tc.isErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.want, got)
			}
		})
	}
}

func TestToExcelGroup(t *testing.T) {
	type testCase[K comparable] struct {
		name     string
		table    [][]string
		pattern  string
		wantKeys []K
		wantData map[K][][]string
		wantErr  bool
	}

	tests := []testCase[string]{
		{
			name: "two groups, valid structure",
			table: [][]string{
				{"group1", "", "col1", "col2", "group2", "", "col3"},
				{"", "row1", "1", "3", "", "row1", "5"},
				{"", "row2", "2", "4", "", "row2", "6"},
			},
			pattern:  `^group\d+`,
			wantKeys: []string{"group1", "group2"},
			wantData: map[string][][]string{
				"group1": {
					{"", "col1", "col2"},
					{"row1", "1", "3"},
					{"row2", "2", "4"},
				},
				"group2": {
					{"", "col3"},
					{"row1", "5"},
					{"row2", "6"},
				},
			},
			wantErr: false,
		},
		{
			name:    "empty table",
			table:   [][]string{},
			pattern: `^group\d+`,
			wantErr: true,
		},
		{
			name: "invalid header (not enough cols)",
			table: [][]string{
				{"group1"},
			},
			pattern: `^group\d+`,
			wantErr: true,
		},
		{
			name: "group key not found at start",
			table: [][]string{
				{"", "col1", "col2"},
				{"row1", "1", "2"},
			},
			pattern: `^group\d+`,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			keys, data, err := ToExcelGroup[string](tc.table, tc.pattern)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.wantKeys, keys)
			assert.Equal(t, tc.wantData, data)
		})
	}
}

func TestToExcel2D(t *testing.T) {
	type testCase[RK, CK comparable, V any] struct {
		name     string
		table    [][]string
		wantRows []RK
		wantCols []CK
		wantData map[RK]map[CK]V
		wantErr  bool
	}

	tests := []testCase[string, string, int]{
		{
			name: "normal case",
			table: [][]string{
				{"", "col1", "col2"},
				{"row1", "1", "3"},
				{"row2", "2", "4"},
			},
			wantRows: []string{"row1", "row2"},
			wantCols: []string{"col1", "col2"},
			wantData: map[string]map[string]int{
				"row1": {"col1": 1, "col2": 3},
				"row2": {"col1": 2, "col2": 4},
			},
			wantErr: false,
		},
		{
			name:    "empty table",
			table:   [][]string{},
			wantErr: true,
		},
		{
			name: "not enough columns in header",
			table: [][]string{
				{""},
				{"row1", "1"},
			},
			wantErr: true,
		},
		{
			name: "no data rows",
			table: [][]string{
				{"", "col1", "col2"},
			},
			wantErr: true,
		},
		{
			name: "bad int conversion",
			table: [][]string{
				{"", "col1"},
				{"row1", "bad"},
			},
			wantErr: true,
		},
		{
			name: "row length mismatch",
			table: [][]string{
				{"", "col1", "col2"},
				{"row1", "1"},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rowKeys, colKeys, data, err := ToExcel2D[string, string, int](tc.table)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.wantRows, rowKeys)
			assert.Equal(t, tc.wantCols, colKeys)
			assert.Equal(t, tc.wantData, data)
		})
	}
}
