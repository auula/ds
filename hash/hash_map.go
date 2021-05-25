// Open Source: MIT License
// Author: Jaco Ding <ding@ibyte.me>
// Date: 2021-05-25 16:40:29 - UTC/GMT+08:00

package hash

import (
	"context"
	"github.com/higker/ds"
	"github.com/higker/ds/list"
	"hash/crc32"
)

const (
	loadFactor = 0.75 // HashMap Load factor
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

// HashMap hash table data structure
type Map struct {
	capacity int
	size     int
	entry    []*list.LinkedList
}

func NewMap() *Map {
	return &Map{
		capacity: 0,
		size:     20,
		entry:    make([]*list.LinkedList, 20),
	}
}

// Code handler hash function
func Code(key interface{}) int {
	var code int = -1
	switch key.(type) {
	case string:
		code = _stringToCode(key.(string))
	case int, int64:
		code = 1
	}
	return code
}

// FindIndex Query from Table Entry index
// index calculate formula: index = hash func % entry.len % size
func (hash *Map) FindIndex(key interface{}) int {
	return Code(key) & cap(hash.entry) % hash.size
}

// FindEntry Query from Table Entry
func (hash *Map) FindEntry(key interface{}) *list.LinkedList {
	return hash.entry[Code(key)&cap(hash.entry)%hash.size]
}

func (hash *Map) Put(key, value interface{}) {
	index := hash.FindIndex(key)

	if hash.entry[index] == nil {
		hash.entry[index] = list.New()
	}
	// 去重 拿到node 检测k
	hash.entry[index].Add(&MapItem{k: key, v: value})
}

func (hash *Map) Get(k interface{}) interface{} {

	channel := make(chan ds.Element, 10)
	ctx, cancel := context.WithCancel(context.Background())

	entry := hash.FindEntry(k)

	entry.Range(ctx, channel)
	for element := range channel {
		ele := element.(*ds.Node).Value.(*MapItem)
		if ele.k == k {
			cancel()
			return ele.Val()
		}
	}
	return nil
}

// checkSize https://play.golang.org/p/chJrLC8qHrZ
func (hash *Map) checkSize() bool {
	size := float64(hash.size)
	capacity := float64(hash.capacity) * loadFactor
	return int(size) >= int(capacity)
}
