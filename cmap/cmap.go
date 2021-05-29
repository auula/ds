// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/29 - 2:27 下午 - UTC/GMT+08:00

// ConcurrentMap

package cmap

import (
	"hash/crc32"
	"sync"
)

type blockMap []*sync.Map

type ConcurrentMap struct {
	blockMap     // 分段的同步map块
	size     int // 有多少个块
}

func New(size int) *ConcurrentMap {
	c := new(ConcurrentMap)
	c.size = size
	c.blockMap = make(blockMap, size)
	for i := 0; i < c.size; i++ {
		c.blockMap[i] = &sync.Map{}
	}
	return c
}

func (m *ConcurrentMap) GetMap(key interface{}) *sync.Map {
	// 通过哈希计算得到map的所在的位置
	return m.blockMap[(m.Hash(key) % m.size)]
}

func (m *ConcurrentMap) Set(k, v interface{}) {
	smap := m.GetMap(k)
	smap.Store(k, v)
}

func (m *ConcurrentMap) Get(k interface{}) interface{} {
	smap := m.GetMap(k)
	if v, ok := smap.Load(k); ok {
		return v
	}
	return nil
}

func (m *ConcurrentMap) Hash(key interface{}) (code int) {
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
