package queue

import (
	"fmt"

	"github.com/higker/ds/heap"
)

type PriorityQueue struct {
	heap *heap.Heap
	size int
}

func Priority(n int) *PriorityQueue {
	heap := heap.Build(make([]int, 0, n))
	return &PriorityQueue{
		heap: heap,
		size: n,
	}
}

func (pq *PriorityQueue) EnQueue(v int) error {
	if pq.heap.Size() > pq.size {
		return fmt.Errorf("priority queue full, priority size is %d", pq.size)
	}
	pq.heap.Insert(v)
	pq.size++
	pq.heap.Sort()
	return nil
}

func (pq *PriorityQueue) DeQueue() int {
	pq.size--
	return pq.heap.MoveMin()
}
