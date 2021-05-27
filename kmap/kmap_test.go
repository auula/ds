// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/27 - 1:03 上午 - UTC/GMT+08:00

// todo...

package kmap

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	kMap := New()
	for i := 0; i < 100; i++ {
		kMap.Put(fmt.Sprintf("k%d", _randomInt(1024)), i)
	}

	kMap.Debug()

}

func TestRandInt(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Log(_randomInt(1024*2) % 100)
	}
}
