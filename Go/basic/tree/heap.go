package advanced_data_structure

import "fmt"


/*array*/
type MinHeap []int

func (mh *MinHeap) Build() { *mh = append(*mh, 0) }

func (mh *MinHeap) Put(value int) {
	*mh = append(*mh, value)
	mh.climb(len(*mh) - 1)
}

func (mh *MinHeap) Get() int {
	re := (*mh)[1]
	change(&(*mh)[1], &(*mh)[len(*mh)-1])
	*mh = (*mh)[:len(*mh)-1]
	mh.goDown(1)
	return re
}

func (mh *MinHeap) climb(k int) {
	for k > 1 && compare((*mh)[k>>1], (*mh)[k]) {
		change(&(*mh)[k>>1], &(*mh)[k])
		k >>= 1
	}
}

func (mh *MinHeap) goDown(k int) {
	for k<<1 <= len(*mh)-1 {
		j := k << 1
		if j < len(*mh)-1 && compare((*mh)[j], (*mh)[j+1]) {
			j++
		}
		if compare((*mh)[k], (*mh)[j]) {
			change(&(*mh)[k], &(*mh)[j])
			k = j
		} else {
			break
		}
	}
}

func compare(a, b int) bool { return a > b }

func change(a, b *int) { *a, *b = *b, *a }

func main() {
	list := [...]int{11, 2, 10, 4, 5, 7, 9, 8, 6, 1}

	mh := MinHeap{}
	mh.Build()
	for _, val := range list {
		mh.Put(val)
	}
	fmt.Println(mh)
	fmt.Println(mh.Get())
	fmt.Println(mh)
	fmt.Println(mh.Get())
	fmt.Println(mh)

}
