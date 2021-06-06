package main

import "fmt"

type BinaryHeap struct {
	Heap
}

type Heap interface {
	Build(values ...int)
	// Insert(v int)
	// TakeTop() int
}

type MaxHeap struct {
	arr []int
}

func New(heap Heap) *BinaryHeap {
	return &BinaryHeap{
		Heap: heap,
	}
}

func (mh *MaxHeap) Build(values ...int) {
	mh.floatUp()
	for _, v := range values {
		mh.arr = append(mh.arr, v)
	}
	for i := (len(mh.arr) - 2) / 2; i >= 0; i-- {
		mh.sink(i, len(mh.arr))
	}
	fmt.Println(mh.arr)
}

// 上浮操作
func (mh *MaxHeap) floatUp() {

	// 找到最后一个需要上浮的元素下标拿到元素
	childIndex := len(mh.arr) - 1
	// 拿到他父节点 方便后面比较
	parentIndex := (childIndex - 1) / 2

	element := mh.arr[childIndex]
	// 开始比较如果大于他的父亲就进行往上冒泡
	for childIndex > 0 && element > mh.arr[parentIndex] {
		// 交换位置
		mh.arr[childIndex] = mh.arr[parentIndex]
		// 更新坐标
		childIndex = parentIndex
		// 进行更新父亲坐标
		parentIndex = (childIndex - 1) / 2
	}
	// childIndex 肯定是0
	mh.arr[childIndex] = element
}

func (mh *MaxHeap) sink(parentIndex, length int) {
	element := mh.arr[parentIndex]
	childIndex := 2*parentIndex + 1
	for childIndex < length {
		if childIndex+1 < length && mh.arr[childIndex+1] < mh.arr[childIndex] {
			childIndex++
		}
		if element <= mh.arr[childIndex] {
			break
		}
		mh.arr[parentIndex] = mh.arr[childIndex]
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}
	mh.arr[parentIndex] = element
}

func main() {
	heap := New(&MaxHeap{})
	heap.Build(1, 2, 3, 4, 5, 6)
}
