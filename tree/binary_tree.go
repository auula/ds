package main

import "fmt"

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

func (bt *BinaryTree) pervOrder() {
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

func New() *BinaryTree {
	return &BinaryTree{}
}

func main() {
	tree := New()
	tree.insert(1).
		insert(2).
		insert(3).
		insert(-1).
		insert(4)
	tree.pervOrder()
}
