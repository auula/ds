// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2021/5/20 - 4:46 下午 - UTC/GMT+08:00

package List

import (
	"testing"
)

func TestNew(t *testing.T) {

	channel := make(chan *Node)

	linkedList := New()
	linkedList.Add(&Element{value: 1})

	linkedList.Add(&Element{value: 2})
	linkedList.Add(&Element{value: 3})
	linkedList.Add(&Element{value: 4})

	linkedList.Insertion(3, &Element{value: 2})

	linkedList.Range(channel)

	linkedList.Remove(3)

	t.Log("linkedList last node value :", linkedList.last)
	t.Log("linkedList  node size :", linkedList.size)
	//t.Log("linkedList  node  index 3  :", linkedList.Get(3))
	for node := range channel {
		t.Log(node.ele)
	}
}
