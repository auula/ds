// Open Source: MIT License
// Author: Jaco Ding <ding@ibyte.me>
// Date: 2021/5/21 - 3:13 下午 - UTC/GMT+08:00

package list

import (
	"context"
	"fmt"
	"github.com/higker/ds"
)

type DoublyLinkedList struct {
	head, last *ds.DulNode
	size       int
	err        error
}

func (dl *DoublyLinkedList) Insert(index int, value interface{}) {
	// 检查err 如果有err了这个程序不往下执行
	if dl.err != nil {
		return
	}

	if index < 0 || index > dl.size {
		dl.err = fmt.Errorf("index out of bounds, by Insertion(%d,%s)", index, value)
		return
	}
	node := &ds.DulNode{Value: value}
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
		pervNode := dl.Get(index - 1).(*ds.DulNode)
		node.Perv = pervNode
		node.Next = pervNode.Next
		pervNode.Next = node
	}
	dl.size++
}

func (dl *DoublyLinkedList) Get(index int) ds.Element {
	if dl.err != nil {
		return nil
	}
	if index < 0 || index > dl.size {
		dl.err = fmt.Errorf("index out of bounds, by Get(%d)", index)
		return nil
	}
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
		pervNode := dl.Get(index - 1).(*ds.DulNode)
		pervNode.Next = pervNode.Next.Next
	}
	dl.size--
}

func (dl *DoublyLinkedList) Range(ctx context.Context) (<-chan ds.Element, context.CancelFunc) {
	node := dl.head

	cancelCtx, cancelFunc := context.WithCancel(ctx)
	channel := make(chan ds.Element)

	go func() {
		for {
			select {
			case <-cancelCtx.Done():
				close(channel)
				return
			default:
				if node != nil && channel != nil {
					channel <- node
					node = node.Next
				} else {
					return
				}
			}
		}
	}()
	return channel, cancelFunc
}

func (dl *DoublyLinkedList) Add(value interface{}) {
	node := &ds.DulNode{Value: value}
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

func (dl *DoublyLinkedList) Err() error {
	return dl.err
}

func (dl *DoublyLinkedList) Size() int {
	return dl.size
}

// New create a DoublyLinkedList
func NewDoubly() List {
	return &DoublyLinkedList{
		head: nil,
		last: nil,
		size: 0,
	}
}
