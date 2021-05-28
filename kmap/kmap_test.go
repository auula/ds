// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/27 - 1:03 上午 - UTC/GMT+08:00

// todo...

package kmap

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"sync"
	"testing"
)

func TestKMap(t *testing.T) {

	kMap := New()

	for i := 0; i < 100000; i++ {
		kMap.Put(i, i)
	}

	t.Log(kMap.Get(999))
	t.Log(kMap.Capacity())
}

func TestMap(t *testing.T) {

	var maps sync.Map

	for i := 0; i < 100000; i++ {
		maps.Store(i, i)
	}

}

func TestKMapCRUD(t *testing.T) {
	kMap := New()
	kMap.Put("foo", "bar")
	kMap.Remove("foo")
	t.Log(kMap.Get("foo"))
	kMap.Put("foo", "bar")
	t.Log(kMap.Get("foo"))
}

func BenchmarkWriteKMap(b *testing.B) {
	maps := make(map[interface{}]interface{}, b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		maps[fmt.Sprintf("k%d", _randomInt(1024))] = n
	}
}

func BenchmarkWriteMap(b *testing.B) {
	kMap := New()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		kMap.Put(fmt.Sprintf("k%d", _randomInt(1024)), n)
	}
}

func _randomInt(max int) int {
	var n uint16
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return int(n) % max

}
