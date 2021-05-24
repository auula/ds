// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/24 - 4:27 下午 - UTC/GMT+08:00

package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := New()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
	t.Log(queue.DeQueue())
}
