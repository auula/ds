// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/21 - 4:19 下午 - UTC/GMT+08:00

package list

import (
	"testing"

	"github.com/higker/ds"
)

func TestNewDoublyLinkedList(t *testing.T) {

	channel := make(chan ds.Element)

	list := NewDoublyLinkedList()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(4)
	list.Add(5)

	list.Insertion(5, 6)

	//list.Remove(4)

	list.Range(channel)

	for node := range channel {
		t.Log(node.Val())
		if node.Val() == 4 {
			t.Log("4 perv is", node.(*ds.DulNode).Perv.Val())
		}
	}
}
