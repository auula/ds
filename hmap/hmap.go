
// // Open Source: MIT License
// // Author: Jaco Ding <deen.job@qq.com>
// // Date: 2021/5/28 - 10:52 下午 - UTC/GMT+08:00

// package hmap

// import (
// 	"hash/crc32"
// 	"sync"
// )

// // 深挖map读写性能
// var (
// 	// 为了快速查找建立外部索引 k:1,34 就能快速查找到位置
// 	// _index map[interface{}][2]int
// 	_index sync.Map
// )

// type Map interface {
// 	Put(k, v interface{}) bool
// 	Replace(k, v interface{})
// 	Get(k interface{}) interface{}
// 	Remove(k interface{})
// 	Capacity() int
// }

// type MapItem struct {
// 	k, v interface{}
// }

// type Bucket struct {
// 	tailPointer  int
// 	data         []*MapItem
// 	size         int // 这个确定的
// 	sync.RWMutex     // 各个分片Map各自的锁 缩小锁的资源颗粒度
// }

// type HMap struct {
// 	size  int // size 是entry的个数
// 	entry []*Bucket
// }

// // 1. for 初始化
// func New() Map {
// 	m := new(HMap)
// 	m.entry = make([]*Bucket, 32)
// 	// 初始化索引
// 	for i := range m.entry {
// 		bk := new(Bucket)
// 		mapItems := make([]*MapItem, 1024<<8)
// 		bk.data = mapItems
// 		bk.size = cap(mapItems)
// 		m.entry[i] = bk
// 	}
// 	m.size = cap(m.entry)
// 	return m
// }

// func (m *HMap) Hash(key interface{}) (code int) {
// 	switch key.(type) {
// 	case string:
// 		code = _stringToCode(key.(string))
// 	case int:
// 		code = key.(int)
// 	default:
// 		panic("unsupported type")
// 	}
// 	return
// }

// // 通过哈希计算 得到root节点下标
// func (m *HMap) index(k interface{}) int {
// 	return m.Hash(k) % m.size
// }

// // 通过索引拿到数据桶
// func (m *HMap) GetBucket(index int) *Bucket {
// 	return m.entry[index]
// }

// func (m *HMap) Get(k interface{}) interface{} {
// 	// 检测是否存在
// 	if _, ok := _index.Load(k); !ok {
// 		return nil
// 	}
// 	index, _ := _index.Load(k)

// 	// 直接通过缓存的索引拿数据 因为是切片时间复杂度 O(1)
// 	return m.entry[index.([2]int)[0]].data[index.([2]int)[1]].v
// }

// func (m *HMap) Put(k, v interface{}) bool {

// 	if _, ok := _index.Load(k); ok {
// 		return false
// 	}

// 	// 通过计算拿到所在的数据桶下标
// 	bucketIndex := m.index(k)

// 	// 通过Bucket的索引拿到桶
// 	bucket := m.GetBucket(bucketIndex)

// 	bucket.Lock()
// 	// 如果找到了说明已经满了，让Bucket扩容 横向水平扩容，此处锁的的颗粒度非常大，严重影响读写性能
// 	if bucket.tailPointer == bucket.size {
// 		bucket.resize()
// 	}
// 	bucket.data[bucket.tailPointer] = &MapItem{k: k, v: v}
// 	bucket.tailPointer++
// 	bucket.Unlock()
// 	_index.Store(k, [2]int{bucketIndex, bucket.tailPointer - 1})

// 	return true
// }

// func (m *HMap) Remove(k interface{}) {
// 	if _, ok := _index.Load(k); !ok {
// 		return
// 	}
// 	coordinate, _ := _index.Load(k)
// 	m.GetBucket(coordinate.([2]int)[0]).Del(coordinate.([2]int)[1])
// 	// 移除索引
// 	_index.Delete(k)
// }

// func (m *HMap) Replace(k, v interface{}) {
// 	m.Remove(k)
// 	m.Put(k, v)
// }

// // 会实时返回当前KMap的容量
// func (m *HMap) Capacity() int {
// 	sum := 0
// 	for i := range m.entry {
// 		sum += m.entry[i].size
// 	}
// 	return sum
// }

// func (b *Bucket) resize() {
// 	newData := make([]*MapItem, cap(b.data)*2)
// 	i := 0
// 	for i = range newData {
// 		// 扩容还原老的数据
// 		if i < cap(b.data) {
// 			newData[i] = b.data[i]
// 		} else {
// 			break
// 		}
// 	}
// 	// 尾指针永远指向空位
// 	b.size = cap(newData)
// 	b.data = newData
// 	b.tailPointer = i
// }

// func (b *Bucket) Del(x int) {
// 	b.Lock()
// 	b.data[x] = nil
// 	b.Unlock()
// }

// func _stringToCode(s string) int {
// 	v := int(crc32.ChecksumIEEE([]byte(s)))
// 	if v >= 0 {
// 		return v
// 	}
// 	if -v >= 0 {
// 		return -v
// 	}
// 	// v == MinInt
// 	return 0
// }

