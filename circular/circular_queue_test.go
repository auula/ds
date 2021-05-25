// Open Source: MIT License
// Author: Jaco Ding <ding@ibyte.me>
// Date: 2021/5/24 - 9:43 下午 - UTC/GMT+08:00

package circular

import (
	"testing"
)

func TestCycleQueue(t *testing.T) {
	queue, err := NewQueue(20)
	t.Error(err)

	for i := 0; i < 22; i++ {
		err := queue.Push(i)
		if err != nil {
			t.Error(err)
			break
		}
	}
	t.Log(queue.data)
	for i := 0; i < 22; i++ {
		t.Log(queue.Pop())
	}
	t.Log(queue.data)
}
