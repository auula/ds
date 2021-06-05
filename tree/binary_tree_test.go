package main

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := New()
	tree.insert(1).
		insert(2).
		insert(3).
		insert(4).
		insert(5)
	//DepthTraverse(tree.root)
	traverseChan := make(chan interface{})
	go DepthTraverse(tree.root, traverseChan)
	for v := range traverseChan {
		t.Log(v)
	}
}
