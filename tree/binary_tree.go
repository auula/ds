package main

import (
	"fmt"

	"github.com/higker/ds/queue"
	"github.com/higker/ds/stack"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int
}

type BinaryTree struct {
	root *BinaryNode
}

func (bt *BinaryTree) Insert(v int) *BinaryTree {
	if bt.root == nil {
		bt.root = &BinaryNode{data: v}
	} else {
		bt.root.insert(v)
	}
	return bt
}

func (bt *BinaryTree) Inserts(values ...int) {
	for _, v := range values {
		bt.Insert(v)
	}
}

func (node *BinaryNode) insert(v int) {

	// 递归结束条件
	if node == nil {
		return
	}

	if v <= node.data {
		// 递归查到最后面一个左节点
		if node.left == nil {
			node.left = &BinaryNode{data: v}
		} else {
			node.left.insert(v)
		}
	} else {
		// 递归查到最后面一个右节点
		if node.right == nil {
			node.right = &BinaryNode{data: v}
		} else {
			node.right.insert(v)
		}
	}
}

func (bt *BinaryTree) PervOrder() {
	bt.root.pervOrder()
}

func (node *BinaryNode) pervOrder() {
	fmt.Println(node.data)
	if node.left != nil {
		node.left.pervOrder()
	}
	if node.right != nil {
		node.right.pervOrder()
	}
}

// 递归实现广度优先遍历
func BreadthTraverse(node *BinaryNode, channel chan interface{}) {
	q := queue.New()
	q.EnQueue(node)
	defer close(channel)
	for !q.IsEmpty() {
		tempNode := q.DeQueue().(*BinaryNode)
		channel <- tempNode.data
		if tempNode.left != nil {
			q.EnQueue(tempNode.left)
		}
		if tempNode.right != nil {
			q.EnQueue(tempNode.right)
		}
	}

}

// 深度优先 我简称w遍历
func DepthTraverse(node *BinaryNode, channel chan interface{}) {
	s := stack.New()
	defer close(channel)
	// 这个条件就是帮助回滚
	for node != nil || !s.IsEmpty() {
		// 每3个元素看做一个整体
		for node != nil {
			channel <- node.data
			s.Push(node)
			node = node.left
		}
		if !s.IsEmpty() {
			node = s.Pop().(*BinaryNode)
			node = node.right
		}
	}

}

func New() *BinaryTree {
	return &BinaryTree{}
}
