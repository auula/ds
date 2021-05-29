// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/29 - 2:27 下午 - UTC/GMT+08:00

// ConcurrentMap

package cmap

import (
	"hash/crc32"
	"sync"
)

// 分段map集合
type blockMaps []*blockMap

// 每块map里面存储数据段
type mapItem map[interface{}]interface{}

type ConcurrentMap struct {
	blockMaps     // 分段的同步map块
	size      int // 有多少个块
}

type blockMap struct {
	// 分段的map块
	items mapItem
	sync.RWMutex
}

func (m *blockMap) set(k, v interface{}) {
	m.Lock()
	m.items[k] = v
	m.Unlock()
}

func (m *blockMap) get(k interface{}) interface{} {
	m.RLock()
	v := m.items[k]
	m.RUnlock()
	return v
}

func New(size int) *ConcurrentMap {
	c := new(ConcurrentMap)
	c.size = size
	c.blockMaps = make(blockMaps, size)
	for i := 0; i < c.size; i++ {
		c.blockMaps[i] = &blockMap{
			items: make(mapItem),
		}
	}
	return c
}

func (m *ConcurrentMap) blockMap(key interface{}) *blockMap {
	// 通过哈希计算得到map的所在的位置
	return m.blockMaps[(m.HashCode(key) % m.size)]
}

func (m *ConcurrentMap) Set(k, v interface{}) {
	b := m.blockMap(k)
	b.set(k, v)
}

func (m *ConcurrentMap) Get(k interface{}) interface{} {
	b := m.blockMap(k)
	return b.get(k)
}

func (m *ConcurrentMap) Remove(k interface{}) {
	b := m.blockMap(k)
	b.Lock()
	delete(b.items, k)
	b.Unlock()
}

func (m *ConcurrentMap) HashCode(key interface{}) (code int) {
	switch key.(type) {
	case string:
		code = _stringToCode(key.(string))
	case int:
		code = key.(int)
	default:
		panic("unsupported type")
	}
	return
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
