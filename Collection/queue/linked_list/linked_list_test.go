package queue

import (
	"testing"
)

func BenchmarkLinkedListQueue(b *testing.B) {
	q := NewQueue[int]()
	for i := 0; i < 1000000000; i++ {
		q.Push(i)
	}
	for i := 0; i < 1000000000; i++ {
		_, _ = q.Pop()
	}
}
