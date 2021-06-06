package main

type BinaryHeap struct {
	Heap
}

type Heap interface {
	Build(values ...int)
	Insert(v int)
	TakeTop() int
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
	if mh.arr == nil {
		mh.arr = make([]int, cap(values))
	}
	mh.floatUp()
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
