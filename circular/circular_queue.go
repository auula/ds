// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/24 - 8:03 下午 - UTC/GMT+08:00

// Circular Queue data structure
package circular

import (
	"errors"
	"fmt"
)

// Queue cycle buffer
type CycleQueue struct {
	data                  []interface{} // 存放元素的数组，准确来说是切片
	frontIndex, rearIndex int           // frontIndex 头指针,rearIndex 尾指针
	size                  int           // circular 的大小
}

// NewQueue Circular Queue
func NewQueue(size int) (*CycleQueue, error) {
	if size <= 0 || size < 10 {
		return nil, fmt.Errorf("initialize circular queue size fail,%d not legal,size >= 10", size)
	}
	cq := new(CycleQueue)
	cq.data = make([]interface{}, size)
	cq.size = size
	return cq, nil
}

// Push  add data to queue
func (q *CycleQueue) Push(value interface{}) error {
	if (q.rearIndex+1)%cap(q.data) == q.frontIndex {
		return errors.New("circular queue full")
	}
	q.data[q.rearIndex] = value
	q.rearIndex = (q.rearIndex + 1) % cap(q.data)
	return nil
}

// Pop return queue a front element
func (q *CycleQueue) Pop() interface{} {
	if q.rearIndex == q.frontIndex {
		return nil
	}
	v := q.data[q.frontIndex]
	q.data[q.frontIndex] = nil // 拿除元素 位置就设置为空
	q.frontIndex = (q.frontIndex + 1) % cap(q.data)
	return v
}
