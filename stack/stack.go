// Open Source: MIT License
// Author: Jaco Ding <ding@ibyte.me>
// Date: 2021/5/24 - 4:27 下午 - UTC/GMT+08:00

// stack data structure
package stack

import (
	"github.com/higker/ds"
)

// Stack FILO first in last out
type Stack struct {
	size int
	top  *ds.DulNode // top 就是最顶上的那个元素
}

func New() *Stack {
	return &Stack{
		size: 0,
		top:  nil,
	}
}

func (s *Stack) Push(value interface{}) {
	node := &ds.DulNode{Value: value}
	// 检测是否为空
	if s.size == 0 {
		s.top = node
		s.size++
		return
	}
	// 把插入的节点指向前面一个节点
	node.Perv = s.top
	// 设置当前顶部节点是插入的节点
	s.top = node
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size == 0 {
		return nil
	}
	// 拿到顶部节点 然后指向前面一个
	node := s.top
	s.top = node.Perv
	s.size--
	return node.Val()
}
