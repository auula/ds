// linkedList data structure

package linkedlist

// Node of LinkedList
type Node struct {
	Element interface{}
	Next    *Node
}

// LinkedList data structure
type LinkedList struct {
	Head *Node
	Last *Node
	Size int
}

func New() *LinkedList {
	return &LinkedList{
		Head: new(Node),
		Last: nil,
		Size: 0,
	}
}

func (ldl *LinkedList) Add(value interface{}) {
	node := &Node{
		Element: value,
		Next:    nil,
	}
	// 移动指针
	pointer := ldl.Head
	for pointer.Next != nil {
		pointer = pointer.Next
	}
	pointer.Next = node
	ldl.Last = node
	ldl.Size++
}

func (ldl *LinkedList) Insert(index int, value interface{}) {

	node := &Node{
		Element: value,
		Next:    nil,
	}

	// 插入位置和大小一致说明是尾巴的
	if index == ldl.Size {
		ldl.Last = node
		ldl.Size++
		return
	}

	// 头指针 尾指针
	PervPointer := ldl.Head
	var NextPointer *Node
	for i := 0; i < index-1; i++ {
		PervPointer = PervPointer.Next
	}
	// 新节点 指向之前那个节点下个节点
	NextPointer = PervPointer.Next
	PervPointer.Next = node
	node.Next = NextPointer
	ldl.Size++
}

func (ldl *LinkedList) Remove(index int) {
	node := ldl.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	ldl.Last = node
	node.Next = node.Next.Next
	ldl.Size--
}

func (ldl *LinkedList) Channel(channel chan *Node) {
	node := ldl.Head.Next
	go func() {
		for node != nil {
			channel <- node
			node = node.Next
		}
		close(channel)
	}()

}

func (ldl *LinkedList) Get(index int) *Node {
	node := ldl.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}
