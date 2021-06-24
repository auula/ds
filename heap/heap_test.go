package heap

import (
	"testing"
)

func TestHeap_Sort(t *testing.T) {
	heap := Insert(12, 9, 33, 232, 11, 3, 1)
	heap.Build()
	heap.Sort()
	t.Log(heap.tree)
	heap.MoveMin()
	t.Log(heap.tree)
	heap.MoveMin()
	t.Log(heap.tree)
}
