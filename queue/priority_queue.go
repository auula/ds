package queue

import (
	"github.com/higker/ds/heap"
)

type PriorityQueue struct {
	heap *heap.Heap
	size int
}

func Priority() *PriorityQueue {
	heap := heap.Build(make([]int, 0))
	return &PriorityQueue{
		heap: heap,
		size: 0,
	}
}

func (pq *PriorityQueue) EnQueue(v int) {
	pq.heap.Insert(v)
	pq.size++
	pq.heap.Sort()
}

func (pq *PriorityQueue) DeQueue() int {
	pq.size--
	return pq.heap.MoveMin()
}

func (pq *PriorityQueue) Size() int {
	return pq.size
}
