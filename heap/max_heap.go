package heap

import (
	"fmt"
	"math"
)

type MaxHeap struct {
	Element []int
}

func NewMax() *MaxHeap {
	return &MaxHeap{Element: make([]int, 0)}
}

func (m *MaxHeap) Insert(v int) {
	m.Element = append(m.Element, v)
	m.floatUp()
}

func (m *MaxHeap) Build(values ...int) error {
	if len(values) < 3 {
		return fmt.Errorf("the number of elements does not meet the construction conditions")
	}
	for _, v := range values {
		m.Insert(v)
	}
	return nil
}

func (m *MaxHeap) floatUp() {
	childIndex := len(m.Element) - 1
	fatherIndex := (childIndex - 1) / 2
	ele := m.Element[childIndex]
	for childIndex > 0 && ele > m.Element[fatherIndex] {
		m.Element[childIndex] = m.Element[fatherIndex]
		// 继续找父亲
		childIndex = fatherIndex
		fatherIndex = fatherIndex / 2
	}
	// 说明位置找到了替换
	m.Element[childIndex] = ele
}

func (m *MaxHeap) Max() int {
	if len(m.Element) > 1 {
		return m.Element[0]
	}
	return math.MaxInt64
}

func (m *MaxHeap) sink() {
	fatherIndex := 0
	temp := m.Element[fatherIndex]
	childIndex := 1
	for childIndex < len(m.Element) {
		if childIndex+1 < len(m.Element) && m.Element[childIndex+1] > m.Element[childIndex] {
			childIndex++
		}
		if temp > m.Element[childIndex] {
			break
		}
		m.Element[fatherIndex] = m.Element[childIndex]
		fatherIndex = childIndex
		childIndex = 2*childIndex + 1
	}
	m.Element[fatherIndex] = temp
}

func (m *MaxHeap) DeleteMax() int {
	// 将最后一个节点的值挂到根节点上
	root := m.Element[0]
	m.Element[0] = m.Element[len(m.Element)-1]
	m.sink()
	return root
}

func main() {
	heap := NewMax()
	heap.Insert(3)
	heap.Insert(2)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(5)
	fmt.Println(heap.Element)
	fmt.Println(heap.DeleteMax())
	fmt.Println(heap.Element)
	fmt.Println(heap.DeleteMax())
	fmt.Println(heap.DeleteMax())
	fmt.Println(heap.DeleteMax())
	fmt.Println(heap.DeleteMax())
}
