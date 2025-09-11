package topk

import (
	"container/heap"
	"hash/fnv"
	"math"
	"sort"
)

// Item represents an item and its estimated frequency.
type Item struct {
	Value string
	Count uint32
}

// A MinHeap implements heap.Interface and holds Items.
type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Count < h[j].Count }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// CountMinSketch (simplified for Top-K context)
type CountMinSketch struct {
	counts [][]uint32
	d, w   int
	seeds  []uint64
}

func newCountMinSketch(epsilon, delta float64) *CountMinSketch {
	d := int(math.Ceil(math.Log(1 / delta)))
	w := int(math.Ceil(math.E / epsilon))

	counts := make([][]uint32, d)
	for i := range counts {
		counts[i] = make([]uint32, w)
	}

	seeds := make([]uint64, d)
	for i := 0; i < d; i++ {
		seeds[i] = uint64(i + 1)
	}

	return &CountMinSketch{
		counts: counts,
		d:      d,
		w:      w,
		seeds:  seeds,
	}
}

func (cms *CountMinSketch) hash(data []byte, seed uint64) uint32 {
	h := fnv.New64a()
	h.Write(data)
	h.Write([]byte{byte(seed)})
	return uint32(h.Sum64())
}

func (cms *CountMinSketch) Add(data []byte) {
	for i := 0; i < cms.d; i++ {
		idx := cms.hash(data, cms.seeds[i]) % uint32(cms.w)
		cms.counts[i][idx]++
	}
}

func (cms *CountMinSketch) Estimate(data []byte) uint32 {
	minCount := uint32(math.MaxUint32)
	for i := 0; i < cms.d; i++ {
		idx := cms.hash(data, cms.seeds[i]) % uint32(cms.w)
		if cms.counts[i][idx] < minCount {
			minCount = cms.counts[i][idx]
		}
	}
	return minCount
}

// TopK represents a data structure to find the top K most frequent items.
type TopK struct {
	k         int
	cms       *CountMinSketch
	minHeap   *MinHeap
	itemCounts map[string]uint32 // To store exact counts for items in the heap
}

// New creates a new TopK instance.
// k: The number of top items to track.
// epsilon, delta: Parameters for the underlying Count-Min Sketch.
func New(k int, epsilon, delta float64) *TopK {
	h := &MinHeap{}
	heap.Init(h)
	return &TopK{
		k:         k,
		cms:       newCountMinSketch(epsilon, delta),
		minHeap:   h,
		itemCounts: make(map[string]uint32),
	}
}

// Add adds an item to the TopK tracker.
func (tk *TopK) Add(item string) {
	tk.cms.Add([]byte(item))
	estimatedCount := tk.cms.Estimate([]byte(item))

	// Update exact count for this item
	tk.itemCounts[item] = estimatedCount

	// Check if item is already in heap
	foundInHeap := false
	for i, hItem := range *tk.minHeap {
		if hItem.Value == item {
			(*tk.minHeap)[i].Count = estimatedCount
			heap.Fix(tk.minHeap, i) // Re-heapify
			foundInHeap = true
			break
		}
	}

	if !foundInHeap {
		if tk.minHeap.Len() < tk.k {
			heap.Push(tk.minHeap, Item{Value: item, Count: estimatedCount})
		} else if estimatedCount > (*tk.minHeap)[0].Count {
			heap.Pop(tk.minHeap)
			heap.Push(tk.minHeap, Item{Value: item, Count: estimatedCount})
		}
	}
}

// GetTopK returns the current top K items.
func (tk *TopK) GetTopK() []Item {
	// Create a copy and sort it in descending order
	result := make([]Item, tk.minHeap.Len())
	copy(result, *tk.minHeap)
	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})
	return result
}