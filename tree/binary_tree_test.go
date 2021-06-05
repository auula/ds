package main

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := New()
	// tree.Insert(1).
	// 	Insert(2).
	// 	Insert(3).
	// 	Insert(4).
	// 	Insert(5)
	//DepthTraverse(tree.root)
	// traverseChan := make(chan interface{})
	// go DepthTraverse(tree.root, traverseChan)
	// for v := range traverseChan {
	// 	t.Log(v)
	// }

	traverseChan := make(chan interface{})
	tree.Inserts(1, 1, 1, 2, 2, 3, 4)
	go BreadthTraverse(tree.root, traverseChan)
	for v := range traverseChan {
		t.Log(v)
	}
}
