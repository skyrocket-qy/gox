package binarysearch

/* @tags: binary search */

// find i where nums[i] < comp && nums[i+1] >= comp
func FindLastSmall(nums []int, comp int) int {
	l, r := 0, len(nums)-1
	var m int
	for l <= r {
		m = (l + r) >> 1
		if nums[m] < comp {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return r
}

func FindMinInRotatedArray(nums []int) int {
	i, j := 0, len(nums)-1

	for i <= j {
		mid := (i + j) >> 1
		if nums[mid] > nums[j] {
			i = mid + 1
		} else if nums[mid] < nums[j] {
			j = mid
		} else if nums[mid] == nums[j] {
			j--
		}
	}

	return nums[i]
}
