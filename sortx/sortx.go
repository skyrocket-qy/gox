package sortx

import (
	"slices"
	"sort"
)

func Insert[T any](arr []T, val T, f func(int) bool) []T {
	return slices.Insert(arr, sort.Search(len(arr), f), val)
}
