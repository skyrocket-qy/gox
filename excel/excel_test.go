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
