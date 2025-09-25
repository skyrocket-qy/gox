package sort

/* @tags:sort,two pointer,divide and conquer */

// pivot select
// in / out memory
// recursive / iterative.
func QSortOutMem(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivotIndex := len(nums) / 2
	pivot := nums[pivotIndex]

	var l, r []int

	for i, num := range nums {
		if i == pivotIndex {
			continue
		}

		if num < pivot {
			l = append(l, num)
		} else {
			r = append(r, num)
		}
	}

	l = QSortOutMem(l)
	r = QSortOutMem(r)

	return append(append(l, pivot), r...)
}

func QSortInMem(nums []int, l, r int) {
	if l >= r {
		return
	}

	pivot := nums[(l+r)>>1]

	i, j := l, r
	for i <= j {
		for nums[i] < pivot {
			i++
		}

		for nums[j] > pivot {
			j--
		}

		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	if l < j {
		QSortInMem(nums, l, j)
	}

	if i < r {
		QSortInMem(nums, i, r)
	}
}

func QSortInPartition(nums []int, l, r int) {
	if l >= r {
		return
	}

	partition := func(nums []int, l, r int) int {
		pivot := nums[r]
		i := l - 1

		for j := l; j < r; j++ {
			if nums[j] < pivot {
				i++
				nums[i], nums[j] = nums[j], nums[i]
			}
		}

		i++
		nums[i], nums[r] = nums[r], nums[i]

		return i
	}

	pi := partition(nums, l, r)
	QSortInPartition(nums, l, pi-1)
	QSortInPartition(nums, pi+1, r)
}
