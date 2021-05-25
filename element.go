// Open Source: MIT License
// Author: Jaco Ding <ding@ibyte.me>
// Date: 2021/5/21 - 7:56 下午 - UTC/GMT+08:00

package ds

type Element interface {
	Val() interface{}
}

// Node 是链表的节点
type Node struct {
	Value interface{} `json:"value"`
	Next  *Node       `json:"next_node"`
}

func (n *Node) Val() interface{} {
	// impl element interface
	return n.Value
}

type DulNode struct {
	Perv  *DulNode    `json:"perv_node"`
	Next  *DulNode    `json:"next_node"`
	Value interface{} `json:"value"`
}

func (n *DulNode) Val() interface{} {
	// impl element interface
	return n.Value
}
