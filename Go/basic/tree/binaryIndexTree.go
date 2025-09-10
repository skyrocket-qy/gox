package binaryIndexTree

type bIT []int

func Build(array []int) bIT {
	n := len(array)
	b := make(bIT, len(array), len(array))
	for i, val := range array {
		b[i] = val
		j := i + 1 + lowbit(i+1)
		if j <= n {
			array[j-1] += array[i]
		}
	}
	return b
}

func (bit bIT) Update(i, value int) {
	n := len(bit)
	for i++; i <= n; i += lowbit(i) {
		bit[i-1] += value
	}
}

func (bit bIT) Query(l, r int) int {
	return bit.query(r) - bit.query(l-1)
}

func (bit bIT) query(r int) int {
	sum := 0
	for r++; r > 0; r -= lowbit(r) {
		sum += bit[r-1]
	}
	return sum
}

func lowbit(value int) int {
	return value & -value
}
