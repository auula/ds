package heap

import "fmt"

type MaxHeap struct {
	Element []int
}

func NewMax() *MaxHeap {
	return &MaxHeap{Element: make([]int, 0)}
}

func (m *MaxHeap) Insert(v int) {
	m.Element = append(m.Element, v)
}

func (m *MaxHeap) Build(values ...int) error {
	if len(values) < 3 {
		return fmt.Errorf("the number of elements does not meet the construction conditions")
	}
	for _, v := range values {
		m.Insert(v)
	}
	m.FloatUp()
	return nil
}

func (m *MaxHeap) FloatUp() {
	childIndex := len(m.Element) - 1
	fatherIndex := (childIndex - 1) / 2
	ele := m.Element[childIndex]
	for childIndex > 0 && ele < m.Element[fatherIndex] {
		m.Element[childIndex] = m.Element[fatherIndex]
		// 继续找父亲
		childIndex = fatherIndex
		fatherIndex = (childIndex - 1) / 2
	}
	// 说明位置找到了替换
	m.Element[childIndex] = ele
}

func main() {
	heap := NewMax()
	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(7)
	heap.Insert(8)
	heap.Insert(9)
	heap.Insert(10)
	heap.Insert(-1)
	heap.FloatUp()
	fmt.Println(heap.Element)
}
