// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/26 - 10:52 下午 - UTC/GMT+08:00

package kmap

import (
	"fmt"
	"hash/crc32"
	"sync"
)

// 深挖map读写性能

var (
	// 为了快速查找建立外部索引 k:1,34 就能快速查找到位置
	_index map[interface{}][2]int
)

// 创建的时候计算
func init() {
	_index = make(map[interface{}][2]int, 100)
}

type KMap interface {
	Put(k, v interface{}) bool
	Replace(k, v interface{})
	Get(k interface{}) interface{}
	Remove(k interface{})
	Capacity() int
}

type MapItem struct {
	k, v interface{}
}

type Bucket struct {
	tailPointer  int
	data         []*MapItem
	sync.RWMutex     // 各个分片Map各自的锁 缩小锁的资源颗粒度
	size         int // 这个确定的
}

type Map struct {
	size  int // size 是entry的个数
	entry []*Bucket
	sync.Mutex
}

// 1. for 初始化
func New() KMap {
	m := new(Map)
	m.entry = make([]*Bucket, 8)
	// 初始化索引
	for i := range m.entry {
		bk := new(Bucket)
		mapItems := make([]*MapItem, 20)
		bk.data = mapItems
		bk.size = cap(mapItems)
		m.entry[i] = bk
	}
	m.size = cap(m.entry)
	return m
}

func (m *Map) Hash(key interface{}) (code int) {
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

// 通过哈希计算 得到root节点下标
func (m *Map) index(k interface{}) int {
	return m.Hash(k) % m.size
}

// 通过索引拿到数据桶
func (m *Map) GetBucket(index int) *Bucket {
	return m.entry[index]
}

func (m *Map) Get(k interface{}) interface{} {
	fmt.Println(_index[k])
	// 检测是否存在
	if _, ok := _index[k]; !ok {
		return nil
	}
	return m.entry[_index[k][0]].data[_index[k][1]].v
}

func (m *Map) Put(k, v interface{}) bool {
	if _, ok := _index[k]; ok {
		return false
	}

	// 通过计算拿到所在的数据桶下标
	bucketIndex := m.index(k)

	// 通过Bucket的索引拿到桶
	bucket := m.GetBucket(bucketIndex)
	// 如果找到了说明已经满了，让Bucket扩容 纵向水平扩容
	if bucket.tailPointer == bucket.size {
		bucket.resize()
	}

	bucket.data[bucket.tailPointer] = &MapItem{k: k, v: v}
	_index[k] = [2]int{bucketIndex, bucket.tailPointer}
	bucket.tailPointer++
	fmt.Println(_index[k])
	return true
}

func (m *Map) Remove(k interface{}) {
	if _, ok := _index[k]; !ok {
		return
	}
	m.GetBucket(_index[k][0]).data[_index[k][1]] = nil
	m.GetBucket(_index[k][0]).tailPointer--
	delete(_index, k) // 移除索引
}

func (m *Map) Replace(k, v interface{}) {
	m.Remove(k)
	m.Put(k, v)
}

// 会实时返回当前KMap的容量
func (m *Map) Capacity() int {
	sum := 0
	for i := range m.entry {
		sum += m.entry[i].size
	}
	return m.size * sum
}

func (b *Bucket) resize() {
	b.Lock()
	defer b.Unlock()
	newData := make([]*MapItem, cap(b.data)*2)
	i := 0
	for i = range newData {
		// 扩容还原老的数据
		if i < cap(b.data) {
			newData[i] = b.data[i]
		} else {
			break
		}
	}
	// 尾指针永远指向空位
	b.size = cap(newData)
	b.data = newData
	b.tailPointer = i
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
