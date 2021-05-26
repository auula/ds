// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/26 - 10:52 下午 - UTC/GMT+08:00

// 自己设计的map 超大型单机并发

package kmap

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"time"
)

// 为了快速查找建立外部索引 k:1,34 就能快速查找到位置
var _index map[interface{}][2]int

type KMap interface {
	Put(k interface{}, v interface{})
	Get(k interface{}) interface{}
	Debug()
}

type Root struct {
	lastIndex int
	data      []*MapItem
	size      int // 这个确定的
}

type MapItem struct {
	k, v interface{}
}

type Map struct {
	capacity int
	size     int
	index    []*Root
}

// 1. for 初始化
// 2.
func New() KMap {
	m := new(Map)
	m.index = make([]*Root, 10, 10)
	for i := range m.index {
		root := new(Root)
		mapItems := make([]*MapItem, 100, 100)
		root.data = mapItems
		root.size = 0
		m.index[i] = root
	}
	m.size = 10
	return m
}

func (m *Map) Hash(key interface{}) int {
	var code int = -1
	switch key.(type) {
	case string:
		code = _stringToCode(key.(string))
	case int, int64:
		rand.Seed(time.Now().UnixNano())
		code = rand.Intn(10) + 1
	}
	return code
}

func (m *Map) Index(k interface{}) int {
	return (m.Hash(k) % cap(m.index)) % m.size
}

func (m *Map) Put(k interface{}, v interface{}) {
	// 拿到所在的组，满了重新做一次记录
	root := m.index[m.Index(k)]
	if root.lastIndex == root.size {
		// 容量已经满了
	}
	// 通过尾部指针找到数组当前在哪个位置是空的，把元素插入
	root.data[root.lastIndex] = &MapItem{k: k, v: v}
	root.lastIndex++

}

func (m *Map) Debug() {
	fmt.Println(m.index[1].data[3].k)
	fmt.Println(m.index[1].data[3].v)
}

func (m *Map) Get(k interface{}) interface{} {
	root := m.index[m.Index(k)]
	for _, ele := range root.data {
		if ele.k == k {
			return ele.v
		}
	}
	return nil
}

func (m *Map) Remove(k interface{}) {
	root := m.index[m.Index(k)]
	for _, ele := range root.data {
		if ele.k == k {
			ele = nil
			return
		}
	}
}

func _stringToCode(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}
