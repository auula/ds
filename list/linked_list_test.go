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
	ctx, cancel := context.WithCancel(context.Background())

	list := New()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	//list.Insert(3, 2)

	list.Range(ctx, channel)

	// list.Remove(33)

	t.Log("linkedList size :", list.Size())
	t.Log("linkedList  node  index 3  :", list.Get(2))

	if list.Err() != nil {
		t.Error(list.Err())
	}

	for node := range channel {
		if node.Val() == 7 {
			cancel()
		}
		t.Log(node.Val())
	}
}
