// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/21 - 4:19 下午 - UTC/GMT+08:00

package list

import (
	"context"
	"testing"
	"time"
)

func TestNewDoublyLinkedList(t *testing.T) {

	list := NewDoubly()

	for i := 0; i < 10; i++ {
		list.Add(i)
	}

	//list.Insert(3, 2)

	elements, cancelFunc := list.Range(context.Background())

	// list.Remove(33)

	t.Log("linkedList size :", list.Size())
	t.Log("linkedList  node  index 3  :", list.Get(2))

	if list.Err() != nil {
		t.Error(list.Err())
	}

	for element := range elements {
		go func() {
			time.Sleep(6 * time.Second)
			cancelFunc()
			return
		}()
		time.Sleep(1 * time.Second)
		t.Log(element.Val())
	}
}
