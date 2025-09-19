package prefixsumarray

/* @tags: array,prefix sum */

func NewPreFixSumArray(in []int) []int {
	preFixSumArray := make([]int, len(in))

	preFixSumArray[0] = in[0]
	for i := 1; i < len(in); i++ {
		preFixSumArray[i] = in[i] + preFixSumArray[i-1]
	}

	return preFixSumArray
}
