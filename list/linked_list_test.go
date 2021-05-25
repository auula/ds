// Open Source: MIT License
// Author: Jaco Ding <ding@ibyte.me>
// Date: 2021/5/20 - 4:46 下午 - UTC/GMT+08:00

package list

import (
	"context"
	"testing"

	"github.com/higker/ds"
)

// go test -v --run=TestNew
func TestNew(t *testing.T) {

	channel := make(chan ds.Element)

	linkedList := New()
	linkedList.Add(1)

	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)
	linkedList.Add(4)
	linkedList.Add(6)
	linkedList.Add(7)
	linkedList.Add(7)
	linkedList.Add(7)
	linkedList.Add(7)

	linkedList.Insertion(3, 2)

	ctx, cancel := context.WithCancel(context.Background())

	linkedList.Range(ctx, channel)

	//linkedList.Remove(3)

	t.Log("linkedList last node value :", linkedList.last)
	t.Log("linkedList size :", linkedList.size)
	t.Log("linkedList  node  index 3  :", linkedList.Get(2))

	if linkedList.Try() {
		t.Error(linkedList.Error())
	}

	for node := range channel {
		if node.Val() == 6 {
			cancel()
		}
		t.Log(node.Val())
	}
}
