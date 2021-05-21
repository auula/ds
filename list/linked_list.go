// linkedList data structure
package List

import "errors"

// Node 是链表的节点
type Node struct {
	Value interface{} `json:"value"`
	Next  *Node       `json:"next_node"`
}

// List 链表的抽象接口
type List interface {
	Get(index int) *Node                          // 通过下标获取node
	Remove(index int)                             // 通过下标移除
	Insertion(index int, value interface{}) error // 通过下标插入
	Range(channel chan *Node)                     // 插入channel遍历
	Add(value interface{})                        // 添加元素
}

type LinkedList struct {
	head, last *Node
	size       int
}

// New create a LinkedList
func New() *LinkedList {
	return &LinkedList{
		head: nil,
		last: nil,
		size: 0,
	}
}

func (list *LinkedList) Insertion(index int, value interface{}) error {
	if index < 0 || index > list.size {
		return errors.New("index out of bounds")
	}
	node := &Node{Value: value, Next: nil}
	if list.size == 0 {
		// 空链表
		list.head = node
		list.last = node
	} else if index == 0 {
		// 追加插入
		node.Next = list.head
		list.head = node
	} else if list.size == index {
		list.last.Next = node
		list.last = node
	} else {
		// 找到3个元素中前面的那个节点
		perv := list.Get(index - 1)
		// 把插入节点的next指向原理位置的下一个那块next
		node.Next = perv.Next
		// 前面next指向插入的
		perv.Next = node
	}
	list.size++
	return nil
}

func (list *LinkedList) Get(index int) *Node {
	node := list.head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}

func (list *LinkedList) Remove(index int) {

	if index == 0 {
		next := list.head.Next
		list.head = next
	} else if index == list.size {
		perv := list.Get(list.size - 1)
		list.last = perv
	} else {
		perv := list.Get(index - 1)
		// 移除中间的那个
		next := perv.Next.Next
		perv.Next = next
	}
	list.size--

}

func (list *LinkedList) Range(channel chan *Node) {
	node := list.head
	go func() {
		for node != nil {
			channel <- node
			node = node.Next
		}
		close(channel)
	}()
}

func (list *LinkedList) Add(value interface{}) {
	node := &Node{Value: value}
	if list.size == 0 {
		list.head = node
		list.last = node
	} else {
		list.last.Next = node
		list.last = node
	}
	list.size++
}
