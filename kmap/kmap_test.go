// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/27 - 1:03 上午 - UTC/GMT+08:00

// todo...

package kmap

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestKMap(t *testing.T) {
	k := fmt.Sprintf("k%d", _randomInt(1024*8))
	kMap := New()
	//kMap.Put(k, "bar")
	//for i := 0; i < 1000; i++ {
	//	kMap.Put(fmt.Sprintf("k%d", _randomInt(1024)), i)
	//}

	//kMap.Remove("foo")
	kMap.Put("foo", "bar")

	t.Log(kMap.Get(k))
}

func TestMap(t *testing.T) {
	maps := make(map[interface{}]interface{}, 1000)
	maps["foo"] = "bar"
	for i := 0; i < 1000000; i++ {
		maps[fmt.Sprintf("k%d", _randomInt(1024))] = i
	}
	t.Log(maps["foo"])
}

func TestRandInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(_randomInt(1024*2) % 100)
	}
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
