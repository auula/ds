package heap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := New()
	heap.Build(11, 2, 22, 3, -2, 43, 21, 18)
	t.Log(heap.Min())
}
