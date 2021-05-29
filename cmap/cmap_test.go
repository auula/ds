// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/29 - 3:07 下午 - UTC/GMT+08:00

// todo...

package cmap

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"testing"
)

var N = 10

var simpleMap map[string]interface{}
var mutexMap MutexMap
var syncMap sync.Map
var concurrentMap *ConcurrentMap

type MutexMap struct {
	sync.RWMutex
	m map[string]interface{}
}

func (m *MutexMap) Set(k string, v interface{}) {
	m.Lock()
	m.m[k] = v
	m.Unlock()
}

func (m *MutexMap) Get(k string) (v interface{}) {
	m.RLock()
	v = m.m[k]
	m.RUnlock()
	return
}

func init() {
	simpleMap = make(map[string]interface{}, N)
	mutexMap = MutexMap{m: make(map[string]interface{}, N)}
	syncMap = sync.Map{}
	for i := 0; i < N; i++ {
		syncMap.Store(strconv.Itoa(i), i)
	}
	concurrentMap = New(32)
}

func BenchmarkMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			idx := rand.Intn(N)
			simpleMap[strconv.Itoa(idx)] = j + 1
			wg.Done()
		}
		wg.Wait()
	}
}

func BenchmarkMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			idx := rand.Intn(N)
			_ = simpleMap[strconv.Itoa(idx)]
			wg.Done()
		}
		wg.Wait()
	}
}

func BenchmarkMutexMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				mutexMap.Set(strconv.Itoa(idx), j+1)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkMutexMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				mutexMap.Get(strconv.Itoa(idx))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				syncMap.Store(strconv.Itoa(idx), j+1)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkSyncMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				syncMap.Load(strconv.Itoa(idx))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkCMap_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				concurrentMap.Set(strconv.Itoa(idx), j+1)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkCMap_Get(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		wg.Add(N)
		for j := 0; j < N; j++ {
			go func() {
				idx := rand.Intn(N)
				concurrentMap.Get(strconv.Itoa(idx))
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentMap(b *testing.B) {
	concurrentMap := New(32 * 2)
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
					concurrentMap.Set(key, val)
					wg.Done()
				}(fmt.Sprintf("k%d", _randomInt(nums)), _randomInt(nums))

				// // 模拟并发随机读
				go func(key string) {
					concurrentMap.Get(key)
					wg.Done()
				}(fmt.Sprintf("k%d", _randomInt(nums)))

			}
			wg.Wait()
		})
	}
}

func BenchmarkSyncMap(b *testing.B) {
	m := sync.Map{}
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
					m.Store(key, val)
					wg.Done()
				}(fmt.Sprintf("k%d", _randomInt(nums)), _randomInt(nums))

				// // 模拟并发随机读
				go func(key string) {
					m.Load(key)
					wg.Done()
				}(fmt.Sprintf("k%d", _randomInt(nums)))

			}
			wg.Wait()
		})
	}
}

func _randomInt(max int) int {
	var n uint16
	binary.Read(crand.Reader, binary.LittleEndian, &n)
	return int(n) % max

}
