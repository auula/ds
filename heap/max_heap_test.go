package heap

import (
	"testing"
)

func TestMaxHeap(t *testing.T) {
	heap := NewMax()
	heap.Build(1, 2, 3, 4, 5, 6, 7, 8, 9)
	t.Log(heap.Element)
	t.Log(heap.DeleteMax())
	t.Log(heap.DeleteMax())
	t.Log(heap.DeleteMax())
	t.Log(heap.DeleteMax())
	t.Log(heap.DeleteMax())
}
