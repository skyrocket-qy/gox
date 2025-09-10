package differencearray

/* @tags: array */

/*
Do interval udpate l times first, then query k times
For each interval update, we also need to update two values, O(n * l) -> O(1 * l)
After updating, we need to use prefix sum to rebuild the original array, O(n)
*/

type DiffArr []int

func NewDifferenceArray(in []int) DiffArr {
	// D[i] = in[i] - in[i-1]
	diffArr := make([]int, len(in))
	diffArr[0] = in[0]
	for i := 1; i < len(diffArr); i++ {
		diffArr[i] = in[i] - in[i-1]
	}

	return diffArr
}

func (a DiffArr) IntervalUpdate(l, r, val int) {
	a[l] += val
	if r == len(a)-1 {
		return
	}
	a[r+1] -= val
}

func (a DiffArr) Rebuild() {
	for i := 1; i < len(a); i++ {
		a[i] += a[i-1]
	}
}

func (a DiffArr) Query(i int) int {
	return a[i]
}
