// Open Source: MIT License
// Author: Jaco Ding <ding@ibyte.me>
// Date: 2021-05-26 16:10 - UTC/GMT+08:00

package hash

import (
	"context"
	"hash/crc32"
	"math/rand"
	"time"

	"github.com/higker/ds"
	"github.com/higker/ds/list"
)

const (
	loadFactor = 0.75 // HashMap Load factor
	size       = 10   // HashMap entry element size
)

// String hashes a string to a unique hashcode.
//
// crc32 returns a uint32, but for our use we need
// a non negative integer. Here we cast to an integer
// and invert it if the result is negative.
// https://github.com/hashicorp/terraform/blob/v0.14.11/helper/hashcode/hashcode.go
// test: https://play.golang.org/p/fp5B1ZtbyO2
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

type MapItem struct {
	k, v interface{}
}

func (m *MapItem) Val() interface{} {
	// impl element interface
	return m.v
}

// 怎么设计一个map才合理
// HashMap hash table data structure
type Map struct {
	//capacity int
	size  int
	entry []list.List
	index map[interface{}]int // 记录key录入的下标位置
}

func NewMap() *Map {
	return &Map{
		size:  size,
		entry: make([]list.List, size),
	}
}

// Code handler hash function
func Code(key interface{}) int {
	var code int = -1
	switch key.(type) {
	case string:
		code = _stringToCode(key.(string))
	case int, int64:
		rand.Seed(time.Now().UnixNano())
		code = rand.Intn(size) + 1
	}
	return code
}

// FindIndex Query from Table Entry index
// index calculate formula: index = hash func % entry.len % size
func (hash *Map) FindIndex(key interface{}) int {
	return Code(key) & cap(hash.entry) % hash.size
}

// FindEntry Query from Table Entry
func (hash *Map) FindEntry(key interface{}) list.List {
	return hash.entry[Code(key)&cap(hash.entry)%hash.size]
}

func (hash *Map) Put(key, value interface{}) {
	index := hash.FindIndex(key)

	if hash.entry[index] == nil {
		hash.entry[index] = list.New()
	}
	// 去重 拿到node 检测k ！！ 干脆提示用户是重复的！ 取的时候要注意，！用的时候只能自己确定key是否唯一
	// 干脆这个叫Table吧，重新一个map 做hashmap 和 go map 一样
	hash.entry[index].Add(&MapItem{k: key, v: value})
}

func (hash *Map) Get(k interface{}) interface{} {

	entry := hash.FindEntry(k)

	elements, cancelFunc := entry.Range(context.Background())

	defer cancelFunc()
	for element := range elements {
		ele := element.(*ds.Node).Value.(*MapItem)
		if ele.k == k {
			cancelFunc()
			return ele.Val()
		}
	}

	return nil
}

func (hash *Map) Remove(k interface{}) {
	entry := hash.FindEntry(k)
	elements, cancelFunc := entry.Range(context.Background())
	defer cancelFunc()
	for element := range elements {
		ele := element.(*ds.Node).Value.(*MapItem)
		if ele.k == k {
			cancelFunc()
			ele.v = nil
		}
	}
}

//// checkSize https://play.golang.org/p/chJrLC8qHrZ
//func (hash *Map) checkSize() bool {
//	size := float64(hash.size)
//	capacity := float64(hash.capacity) * loadFactor
//	return int(size) >= int(capacity)
//}
