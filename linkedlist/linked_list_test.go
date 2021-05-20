// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2021/5/20 - 4:46 下午 - UTC/GMT+08:00

package linkedlist

import (
	"testing"
)

func TestNew(t *testing.T) {

	channel := make(chan *Node)

	linkedList := New()
	linkedList.Add(1)
	linkedList.Insert(2, 2)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(5)

	linkedList.Remove(5)

	linkedList.Channel(channel)

	t.Log("linkedList last node value :", linkedList.Last.Element)
	t.Log("linkedList  node size :", linkedList.Size)
	t.Log("linkedList  node  index 3  :", linkedList.Get(3))
	for node := range channel {
		t.Log(node.Element)
	}
}
