// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/25 - 11:49 下午 - UTC/GMT+08:00

package hash

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	tab := NewMap()
	tab.Put("k1", 1)
	tab.Put("k2", 1)
	tab.Put("k3", 1)
	tab.Put("k4", "四")
	tab.Put("k5", 1)
	fmt.Println(tab.Get("k4"))
}
