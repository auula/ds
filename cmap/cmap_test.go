// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/29 - 3:07 下午 - UTC/GMT+08:00

// todo...

package cmap

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkConcurrentMap(b *testing.B) {
	concurrentMap := New(32)
	nums := 10000
	b.ResetTimer()

	for i := 0; i < 2; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {

			b.N = 10000000

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

func _randomInt(max int) int {
	var n uint16
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return int(n) % max

}
