// Open Source: MIT License
// Author: Jaco Ding <ding@ibyte.me>
// Date: 2021/5/21 - 8:30 下午 - UTC/GMT+08:00

package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())
	t.Log(stack.Pop())

}
