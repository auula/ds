// Copyright (c) 2020 sDing
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2021/5/21 - 3:13 下午 - UTC/GMT+08:00

package List

import "fmt"

type DulNode struct {
	Perv  *DulNode    `json:"perv_node"`
	Next  *DulNode    `json:"next_node"`
	Value interface{} `json:"value"`
}

type DoublyLinkedList struct {
	head, last *DulNode
	size       int
	err        error
}

func (dl *DoublyLinkedList) Insertion(index int, value interface{}) {
	// 检查err 如果有err了这个程序不往下执行
	if dl.err != nil {
		return
	}

	if index < 0 || index > dl.size {
		dl.err = fmt.Errorf("index out of bounds, by Insertion(%d,%s)", index, value)
		return
	}
	node := &DulNode{Value: value}
	if dl.size == 0 {
		dl.head = node
		dl.last = node
	} else if index == 0 {
		// 在头部插入
		node.Next = dl.head
		dl.last = node
	} else if index == dl.size {
		// 在尾巴插入
		node.Perv = dl.last
		dl.last.Next = node
		dl.last = node
	} else {
		pervNode := dl.Get(index - 1).(*DulNode)
		node.Perv = pervNode
		node.Next = pervNode.Next
		pervNode.Next = node
	}
	dl.size++
}

func (dl *DoublyLinkedList) Get(index int) Element {
	node := dl.head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}

func (dl *DoublyLinkedList) Remove(index int) {
	if dl.err != nil {
		return
	}
	if index < 0 || index > dl.size {
		dl.err = fmt.Errorf("index out of bounds, by Remove(%d)", index)
		return
	}

	if index == 0 {
		next := dl.head.Next
		dl.head = next
		dl.head.Perv = nil
	} else if index == dl.size {
		perv := dl.last.Perv
		perv.Next = nil
		dl.last = perv
	} else {
		pervNode := dl.Get(index - 1).(*DulNode)
		pervNode.Next = pervNode.Next.Next
	}
	dl.size--
}

func (dl *DoublyLinkedList) Range(channel chan Element) {
	node := dl.head
	go func() {
		for node != nil {
			channel <- node
			node = node.Next
		}
		close(channel)
	}()
}

func (dl *DoublyLinkedList) Add(value interface{}) {
	node := &DulNode{Value: value}
	if dl.size == 0 {
		dl.head = node
		dl.last = node
	} else {
		node.Perv = dl.last
		dl.last.Next = node
		dl.last = node
	}
	dl.size++
}

func (n *DulNode) Val() interface{} {
	// impl element interface
	return n.Value
}

func (dl *DoublyLinkedList) Try() bool {
	return dl.err != nil
}

func (dl *DoublyLinkedList) Error() string {
	return dl.err.Error()
}

// New create a DoublyLinkedList
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{
		head: nil,
		last: nil,
		size: 0,
	}
}
