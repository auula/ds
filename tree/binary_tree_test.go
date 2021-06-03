package main

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := New()
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	tree.pervOrder()
}
