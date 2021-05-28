// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/5/25 - 11:49 下午 - UTC/GMT+08:00

package hash

import (
	"math/rand"
	"testing"
	"time"
)

func TestMap(t *testing.T) {

}

func TestRead(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		t.Log(rand.Intn(10) + 1)
	}
}
