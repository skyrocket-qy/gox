package kthselect

/* @tags: sort */

import (
	"fmt"
	"sort"
)

// find kth larger element
// O(nlogn)
func FindAfterSort(nums []int, k int) (int, error) {
	if k > len(nums) {
		return 0, fmt.Errorf("%d is larger than length of nums", k)
	}

	sort.Ints(nums)
	return nums[len(nums)-k], nil
}
