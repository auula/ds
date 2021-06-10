package heap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := New()
	heap.Build(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1)
	t.Log(heap.DeleteMin())
	t.Log(heap.DeleteMin())
	t.Log(heap.DeleteMin())
}
