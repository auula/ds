// Open Source: MIT License
// Author: Jaco Ding <ding@ibyte.me>
// Date: 2021/5/20 - 4:46 下午 - UTC/GMT+08:00

package list

import (
	"testing"
)

// go test -v --run=TestNew
func TestNew(t *testing.T) {

	channel := make(chan Element)

	linkedList := New()
	linkedList.Add(1)

	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)

	linkedList.Insertion(3, 2)

	linkedList.Range(channel)

	//linkedList.Remove(3)

	t.Log("linkedList last node value :", linkedList.last)
	t.Log("linkedList size :", linkedList.size)
	t.Log("linkedList  node  index 3  :", linkedList.Get(2))

	if linkedList.Try() {
		t.Error(linkedList.Error())
	}

	for node := range channel {
		t.Log(node.Val())
	}
}
