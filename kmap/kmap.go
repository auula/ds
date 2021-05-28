// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/26 - 10:52 下午 - UTC/GMT+08:00

package kmap

import (
	"hash/crc32"
	"sort"
	"sync"
)

var (
	// 为了快速查找建立外部索引 k:1,34 就能快速查找到位置
	_index      map[interface{}][2]int
	_dirtyIndex []int
)

// 创建的时候计算
func init() {
	_index = make(map[interface{}][2]int, 100)
	_dirtyIndex = make([]int, 0, 1024)
}

type KMap interface {
	Put(k, v interface{}) bool
	Replace(k, v interface{})
	Get(k interface{}) interface{}
	Remove(k interface{})
	Debug()
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
	//capacity int
	size  int
	entry []*Bucket
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
	// 检测是否存在
	if _, ok := _index[k]; !ok {
		return nil
	}
	return m.entry[_index[k][0]].data[_index[k][1]].v
}

func (m *Map) Put(k, v interface{}) bool {
	if _, ok := _index[k]; !ok {
		return false
	}

	sort.Ints(_dirtyIndex)

	// 拿到所在的组
	bucketIndex := m.index(k)

	// 通过Bucket的索引检测是否满了，如果找到了说明已经满了，让Bucket扩容 竖向
	if binarySearch(bucketIndex, 0, len(_dirtyIndex), _dirtyIndex) == -1 {
		bucket := m.GetBucket(bucketIndex)
		bucket.data[bucket.tailPointer] = &MapItem{k: k, v: v}
		_index[k] = [2]int{bucketIndex, bucket.tailPointer}
		bucket.tailPointer++
	} else {

	}

	return true
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

func binarySearch(v, left, right int, arr []int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if arr[mid] == v {
		return mid // index
	}
	if arr[mid] < v {
		return binarySearch(v, mid+1, right, arr)
	}
	if arr[mid] > v {
		return binarySearch(v, left, mid-1, arr)
	}
	return -1
}
