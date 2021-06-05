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

func (bt *BinaryTree) insert(v int) *BinaryTree {
	if bt.root == nil {
		bt.root = &BinaryNode{data: v}
	} else {
		bt.root.insert(v)
	}
	return bt
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
func BreadthTraverse(node *BinaryNode) {
	q := queue.New()
	q.EnQueue(node)
	for !q.IsEmpty() {
		node = q.DeQueue().(*BinaryNode)
		fmt.Println(node.data)
		if node.left != nil {
			q.EnQueue(node.left)
		}
		if node.right != nil {
			q.EnQueue(node.right)
		}
	}
}

// 深度优先 我简称w遍历
func DepthTraverse(node *BinaryNode) {
	s := stack.New()
	// 这个条件就是帮助回滚
	for node != nil || !s.IsEmpty() {
		for node != nil {
			fmt.Println(node.data)
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

func main() {
	tree := New()
	tree.insert(1).
		insert(2).
		insert(3).
		insert(4).
		insert(5)
	//DepthTraverse(tree.root)
	BreadthTraverse(tree.root)
}
