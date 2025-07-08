package excel

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

/*
The table format is as follows:
| keyName | valName |
| key1    | val1    |
| key2    | val2    |
*/
func ToExcel1D[K comparable, V any](
	table [][]string,
) (keyName, valName string, data map[K]V, err error) {
	if len(table) == 0 {
		return "", "", nil, errors.New("empty table")
	}

	if len(table[0]) != 2 {
		return "", "", nil, errors.New("table format error")
	}

	if len(table) < 2 {
		return "", "", nil, errors.New("no data")
	}

	keyName = table[0][0]
	valName = table[0][1]
	data = make(map[K]V)
	for _, row := range table[1:] {
		if len(row) != 2 {
			return "", "", nil, errors.New("table format error")
		}

		k, err := StrToType[K](row[0])
		if err != nil {
			return "", "", nil, err
		}
		v, err := StrToType[V](row[1])
		if err != nil {
			return "", "", nil, err
		}
		data[k] = v
	}

	return keyName, valName, data, nil
}

func StrToType[T any](s string) (T, error) {
	var zero T
	t := reflect.TypeOf(zero)

	switch t.Kind() {
	case reflect.String:
		return any(s).(T), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return zero, err
		}
		return any(int(i)).(T), nil // careful, might truncate for smaller ints
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return zero, err
		}
		return any(f).(T), nil
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return zero, err
		}
		return any(b).(T), nil
	default:
		return zero, fmt.Errorf("unsupported type %v", t)
	}
}

/*
The table format is as follows:
|  			 | colKey1 | colKey2 |
| rowKey1    | val1    | val3    |
| rowKey2    | val2    | val4    |
*/
func ToExcel2D[RowKey, ColKey comparable, V any](table [][]string) (
	rowKeys []RowKey, colKeys []ColKey, data map[RowKey]map[ColKey]V, err error,
) {
	if len(table) == 0 {
		return nil, nil, nil, errors.New("empty table")
	}

	if len(table[0]) < 2 {
		return nil, nil, nil, errors.New("table format error")
	}

	if len(table) < 2 {
		return nil, nil, nil, errors.New("no data")
	}

	rowKeys = make([]RowKey, 0, len(table)-1)
	colKeys = make([]ColKey, 0, len(table[0])-1)
	data = make(map[RowKey]map[ColKey]V)

	for _, col := range table[0][1:] {
		k, err := StrToType[ColKey](col)
		if err != nil {
			return nil, nil, nil, err
		}
		colKeys = append(colKeys, k)
	}

	for _, row := range table[1:] {
		if len(row) < len(table[0]) {
			return nil, nil, nil, errors.New("table format error")
		}

		rowKey, err := StrToType[RowKey](row[0])
		if err != nil {
			return nil, nil, nil, err
		}
		rowKeys = append(rowKeys, rowKey)

		colToVal := make(map[ColKey]V)
		for i, val := range row[1:] {
			colKey := colKeys[i]

			v, err := StrToType[V](val)
			if err != nil {
				return nil, nil, nil, err
			}
			colToVal[colKey] = v
		}
		data[rowKey] = colToVal
	}

	return rowKeys, colKeys, data, nil
}

/*
The table format is as follows:
| gorup1 |         | colKey1 | colKey2 | group2 |		  | colKey1 |
|        | rowKey1 | val1    | val3    |        | rowKey1 | val1    |
|        | rowKey2 | val2    | val4    |        | rowKey2 | val2    |
*/
func ToExcelGroup[GroupKey comparable](
	table [][]string, groupKeyPattern string,
) (groupKeys []GroupKey, data map[GroupKey][][]string, err error) {
	if len(table) == 0 {
		return nil, nil, errors.New("empty table")
	}

	if len(table[0]) < 2 {
		return nil, nil, errors.New("table format error")
	}

	_, err = regexp.MatchString(groupKeyPattern, table[0][0])
	if err != nil {
		return nil, nil, err
	}
	_, err = regexp.MatchString(groupKeyPattern, table[0][len(table[0])-1])
	if err != nil {
		return nil, nil, err
	}

	left := 0
	data = map[GroupKey][][]string{}
	var preGroupKey GroupKey
	for i, col := range table[0] {
		matched, err := regexp.MatchString(groupKeyPattern, col)
		if err != nil {
			return nil, nil, err
		}
		if matched {
			groupKey, err := StrToType[GroupKey](table[0][i])
			if err != nil {
				return nil, nil, err
			}

			if left != 0 {
				// means add to group
				data[preGroupKey] = GetGroup(table, left, i-1)
			}
			preGroupKey = groupKey
			groupKeys = append(groupKeys, groupKey)
			left = i + 1
		} else {
			if left == 0 {
				return nil, nil, errors.New("group key not found")
			}
		}

		if i == len(table[0])-1 {
			data[preGroupKey] = GetGroup(table, left, i)
		}
	}

	return groupKeys, data, nil
}

func GetGroup(table [][]string, left, right int) [][]string {
	result := make([][]string, len(table))
	for i, row := range table {
		if len(row) < right+1 {
			result[i] = row[left:]
			continue
		}
		result[i] = row[left : right+1]
	}
	return result
}

/*
The table format is as follows:
| colKey1 | colKey2 |
| val1    | val1    |
| val2    | val2    |
|         | val3    |
each col len separately
*/
func ToColsList[K comparable, V any](table [][]string) (data map[K][]V, err error) {
	if len(table) == 0 {
		return nil, errors.New("empty table")
	}

	data = map[K][]V{}
	for j, col := range table[0] {
		key, err := StrToType[K](col)
		if err != nil {
			return nil, err
		}

		vals := []V{}
		for i := 1; i < len(table); i++ {
			if len(table[i]) < j+1 {
				break
			}

			val, err := StrToType[V](table[i][j])
			if err != nil {
				return nil, err
			}
			vals = append(vals, val)
		}

		data[key] = vals
	}

	return data, nil
}
