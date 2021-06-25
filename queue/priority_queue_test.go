package queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := Priority(5)
	pq.EnQueue(1)
	pq.EnQueue(12)
	pq.EnQueue(412)
	pq.EnQueue(65)
	pq.EnQueue(212)
	t.Log(pq.DeQueue())
	t.Log(pq.DeQueue())
	t.Log(pq.DeQueue())
	t.Log(pq.DeQueue())
	t.Log(pq.DeQueue())
}
