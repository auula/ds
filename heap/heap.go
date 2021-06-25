package heap

type Heap struct {
	tree []int
	size int
}

// 构建一个大顶堆
func Build(arr []int) *Heap {
	heap := &Heap{
		tree: arr,
		size: len(arr),
	}
	heap.Build()
	return heap
}

func Insert(values ...int) *Heap {
	heap := new(Heap)
	heap.tree = values
	heap.size = len(heap.tree)
	return heap
}

func heapify(tree []int, n, i int) {
	if i >= n {
		return
	}
	// 拿到左右孩子节点的下标
	leftChild, rightChild := 2*i+1, 2*i+2
	max := i
	if leftChild < n && tree[leftChild] > tree[max] {
		max = leftChild
	}
	if rightChild < n && tree[rightChild] > tree[max] {
		max = rightChild
	}
	if max != i {
		tree[i], tree[max] = tree[max], tree[i]
		heapify(tree, n, max)
	}
}

func (h *Heap) Sort() {
	for i := h.size - 1; i >= 0; i-- {
		h.tree[0], h.tree[i] = h.tree[i], h.tree[0]
		heapify(h.tree, i, 0)
	}
}

func (h *Heap) Build() {
	lastIndex := h.size - 1
	parentIndex := (lastIndex - 1) / 2
	for i := parentIndex; i >= 0; i-- {
		h.Heapify(i)
	}
}

func (h *Heap) MoveMin() int {

	defer func() {
		h.tree = h.tree[1:]
		h.size--
		h.Build()
		h.Sort()
	}()

	return h.tree[0]
}

func (h *Heap) Heapify(index int) {
	heapify(h.tree, h.size, index)
}

func (h *Heap) Insert(v int) {
	h.tree = append(h.tree, v)
	h.size = len(h.tree)
	h.Build()
}

func (h *Heap) Size() int {
	return len(h)
}
