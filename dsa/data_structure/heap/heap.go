package heap

/* @tags: heap,priority queue */

// The difference between built-in heap is:
// You only need to implement the Element interface
// and the Less method which is used to compare elements
// And, there is some performance optimization than built-in at Pop()
// because built-in sacrifice some performance for better flexibility
type Element interface {
	Less(e Element) bool
}

type Heap []Element

func Init(eles []Element) *Heap {
	// offset
	h := &Heap{eles[0]}
	for _, e := range eles {
		h.Push(e)
	}
	return h
}

func (h Heap) Len() int { return len(h) - 1 }

func (h Heap) swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(e Element) {
	*h = append(*h, e)
	h.up(len(*h) - 1)
}

func (h *Heap) Pop() Element {
	n := len(*h) - 1
	res := (*h)[1]
	(*h)[1] = (*h)[n]
	*h = (*h)[:n]
	h.down(1)
	return res
}

func (h Heap) up(i int) {
	for j := i >> 1; i > 1 && h[j].Less(h[i]); i, j = j, j>>1 {
		h.swap(i, j)
	}
}

func (h Heap) down(i int) {
	for j := i << 1; j < len(h); i, j = j, j<<1 {
		if j+1 < len(h) && h[j].Less(h[j+1]) {
			j++
		}
		if h[j].Less(h[i]) {
			break
		}
		h.swap(i, j)
	}
}
