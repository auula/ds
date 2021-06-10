package heap

import (
	"fmt"
	"math"
)

type MinHeap struct {
	Element []int
}

func New() *MinHeap {
	return &MinHeap{Element: []int{math.MinInt64}}
}

func (m *MinHeap) Build(values ...int) {
	for _, v := range values {
		m.Insert(v)
	}
}

// 插入数字,插入数字需要保证堆的性质
func (m *MinHeap) Insert(v int) {
	m.Element = append(m.Element, v)
	i := len(m.Element) - 1
	// 上浮 找到前一个元素的下标拿到元素进行比较
	for m.Element[i/2] > v {
		m.Element[i] = m.Element[i/2]
		i /= 2
	}
	m.Element[i] = v
}

// 删除并返回最小值
func (m *MinHeap) DeleteMin() (int, error) {
	if len(m.Element) <= 1 {
		return 0, fmt.Errorf("MinHeap is empty")
	}
	minElement := m.Element[1]
	lastElement := m.Element[len(m.Element)-1]
	var i, child int
	for i = 1; i*2 < len(m.Element); i = child {
		child = i * 2
		if child < len(m.Element)-1 && m.Element[child+1] < m.Element[child] {
			child++
		}
		// 下滤一层
		if lastElement > m.Element[child] {
			m.Element[i] = m.Element[child]
		} else {
			break
		}
	}
	m.Element[i] = lastElement
	m.Element = m.Element[:len(m.Element)-1]
	return minElement, nil
}

// 堆的大小
func (m *MinHeap) Length() int {
	return len(m.Element) - 1
}

// 获取最小堆的最小值
func (m *MinHeap) Min() (int, error) {
	if len(m.Element) > 1 {
		return m.Element[1], nil
	}
	return 0, fmt.Errorf("heap is empty")
}

// MinHeap格式化输出
func (m *MinHeap) String() string {
	return fmt.Sprintf("%v", m.Element[1:])
}
