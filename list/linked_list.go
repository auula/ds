// Open Source: MIT License
// Author: Jaco Ding <ding@ibyte.me>
// Date: 2021/5/20 - 4:46 下午 - UTC/GMT+08:00

// LinkedList data structure

package list

import (
	"context"
	"fmt"

	"github.com/higker/ds"
)

// List 链表的抽象接口
type List interface {
	Get(index int) ds.Element                                          // 通过下标获取node
	Remove(index int)                                                  // 通过下标移除
	Insert(index int, value interface{})                               // 通过下标插入
	Range(ctx context.Context) (<-chan ds.Element, context.CancelFunc) // 遍历
	Add(value interface{})                                             // 添加元素
	Size() int                                                         // List大小
	Err() error                                                        // 是否发生Error
}

type LinkedList struct {
	head, last *ds.Node
	size       int
	err        error
}

// New create a LinkedList
func New() List {
	return &LinkedList{
		head: nil,
		last: nil,
		size: 0,
	}
}

func (list *LinkedList) Insert(index int, value interface{}) {

	// 检查err 如果有err了这个程序不往下执行
	if list.err != nil {
		return
	}

	if index < 0 || index > list.size {
		list.err = fmt.Errorf("index out of bounds, by Insertion(%d,%s)", index, value)
		return
	}

	node := &ds.Node{Value: value, Next: nil}
	if list.size == 0 {
		// 空链表
		list.head = node
		list.last = node
	} else if index == 0 {
		// 头部插入
		node.Next = list.head
		list.head = node
	} else if list.size == index {
		list.last.Next = node
		list.last = node
	} else {
		// 找到3个元素中前面的那个节点
		perv := list.Get(index - 1).(*ds.Node)
		// 把插入节点的next指向原理位置的下一个那块next
		node.Next = perv.Next
		// 前面next指向插入的
		perv.Next = node
	}
	list.size++

}

func (list *LinkedList) Get(index int) ds.Element {
	if list.err != nil {
		return nil
	}
	if index < 0 || index > list.size {
		list.err = fmt.Errorf("index out of bounds, by Get(%d)", index)
		return nil
	}
	node := list.head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node
}

func (list *LinkedList) Remove(index int) {

	if list.err != nil {
		return
	}
	if index < 0 || index > list.size {
		list.err = fmt.Errorf("index out of bounds, by Remove(%d)", index)
		return
	}

	if index == 0 {
		next := list.head.Next
		list.head = next
	} else if index == list.size {
		perv := list.Get(list.size - 1).(*ds.Node)
		list.last = perv
	} else {
		perv := list.Get(index - 1).(*ds.Node)
		// 移除中间的那个
		perv.Next = perv.Next.Next
	}
	list.size--

}

func (list *LinkedList) Range(ctx context.Context) (<-chan ds.Element, context.CancelFunc) {
	node := list.head
	cancelCtx, cancelFunc := context.WithCancel(ctx)
	// 不需要二次关闭 close(channel) 否则 panic
	// 这个通道应该是在上层关闭 而不是在这里关闭
	// 重复关闭则 panic
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
					//if channel != nil {
					//	close(channel)
					//	return
					//}
					return
				}
			}
		}
	}()
	return channel, cancelFunc
}

func (list *LinkedList) Add(value interface{}) {
	node := &ds.Node{Value: value}
	if list.size == 0 {
		list.head = node
		list.last = node
	} else {
		list.last.Next = node
		list.last = node
	}
	list.size++
}

func (list *LinkedList) Err() error {
	return list.err
}

func (list *LinkedList) Size() int {
	return list.size
}
