package sort

/* @tags: sort */

import (
	"time"
)

// hack method, just representing
func SleepSort(nums []int) []int {
	res := make([]int, len(nums))
	i := 0
	sleep := func(num int, res []int, i *int) {
		time.Sleep(time.Duration(num) * time.Second)
		res[*i] = num
		*i++
	}

	for _, num := range nums {
		go sleep(num, res, &i)
	}

	time.Sleep(10 * time.Second)
	return res
}
