// Open Source: MIT License
// Author: Jaco Ding <ding@ibyte.me>
// Date: 2021/5/29 - 1:03 上午 - UTC/GMT+08:00

// todo...

package kmap

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func TestKMap(t *testing.T) {

	kMap := New()

	for i := 0; i < 1000000; i++ {
		kMap.Put(i, i)
	}

	for i := 0; i < 1000000-1; i++ {
		kMap.Get(i)
	}
}

func TestMap(t *testing.T) {

	var maps sync.Map

	for i := 0; i < 1000000; i++ {
		maps.Store(i, i)
	}

	for i := 0; i < 100000; i++ {
		maps.Load(i)
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

func BenchmarkKMapShared(b *testing.B) {

	m := New()
	nums := 10000
	b.ResetTimer()

	for i := 0; i < 5; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {

			b.N = 1000000

			wg := sync.WaitGroup{}
			wg.Add(b.N * 2)
			for i := 0; i < b.N; i++ {

				// 模拟并发随机写
				go func(key string, val interface{}) {
					m.Put(key, val)
					wg.Done()
				}(fmt.Sprintf("k%d", _randomInt(nums)), _randomInt(nums))

				// // 模拟并发随机读
				go func(key string) {
					m.Get(key)
					wg.Done()
				}(fmt.Sprintf("k%d", _randomInt(nums)))

			}
			wg.Wait()
		})
	}
}

func BenchmarkMapShared(b *testing.B) {

	nums := 10000
	mux := sync.RWMutex{}
	maps := make(map[string]interface{}, nums)

	b.ResetTimer()

	for i := 0; i < 5; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {

			b.N = 10000000

			wg := sync.WaitGroup{}
			wg.Add(b.N * 2)
			for i := 0; i < b.N; i++ {

				// 模拟并发随机写
				go func(key string, val interface{}) {
					mux.Lock()
					maps[key] = val
					mux.Unlock()
					wg.Done()
				}(fmt.Sprintf("k%d", _randomInt(nums)), _randomInt(nums))

				// // 模拟并发随机读
				go func(key string) {
					mux.Lock()
					_ = maps[key]
					mux.Unlock()
					wg.Done()
				}(fmt.Sprintf("k%d", _randomInt(nums)))

			}
			wg.Wait()
		})
	}
}
