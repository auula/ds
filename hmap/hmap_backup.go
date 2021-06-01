package hmap

import (
	"hash/crc32"
	"sync"
)

type mapItem struct {
	k, v interface{}
}

type Bucket struct {
	pointer int
	data    []*mapItem
	size    int
	sync.RWMutex
}

type HMap struct {
	size  int
	entry []*Bucket
}

func New(size int) *HMap {
	m := new(HMap)
	m.entry = make([]*Bucket, size)
	for i := range m.entry {
		bk := new(Bucket)
		mapItems := make([]*mapItem, 1024)
		bk.data = mapItems
		bk.size = cap(mapItems)
		m.entry[i] = bk
	}
	m.size = cap(m.entry)
	return m
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

func (m *HMap) Hash(key interface{}) (code int) {
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

func (m *HMap) Put(k, v interface{}) {
	bucket := m.GetBucket(m.HashCode(k))
	bucket.Lock()
	bucket.data[bucket.pointer] = &mapItem{k: k, v: v}
	bucket.pointer++
	bucket.Unlock()
}

func (m *HMap) Get(k interface{}) interface{} {
	var value interface{}
	bucket := m.GetBucket(m.HashCode(k))
	bucket.Lock()
	data := bucket.data
	for _, v := range data {
		if v.k.(string) == k {
			value = v.v
			break
		}
	}
	bucket.Unlock()
	return value
}

func (m *HMap) Remove(k interface{}) {
	bucket := m.GetBucket(m.HashCode(k))
	bucket.Lock()
	data := bucket.data
	// 通过下标移除元素
	for i, v := range data {
		if v.k.(string) == k {
			bucket.data = append(bucket.data[:i], bucket.data[i+1:]...)
		}
	}
	bucket.Unlock()
}

// 通过索引拿到数据桶
func (m *HMap) GetBucket(index int) *Bucket {
	return m.entry[index]
}

func (m *HMap) HashCode(k interface{}) int {
	return m.Hash(k) % m.size
}
