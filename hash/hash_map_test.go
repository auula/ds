// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/25 - 11:49 下午 - UTC/GMT+08:00

package hash

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	tab := NewMap()
	tab.Put("k1", 1)
	tab.Put("k2", 1)
	tab.Put("k3", 1)
	tab.Put("k4", "四")
	tab.Put("k5", 1)
	tab.Remove("k3")
	fmt.Println(tab.Get("k3"))
	fmt.Println(tab.Get("k4"))
}

func TestRead(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		t.Log(rand.Intn(10) + 1)
	}
}
