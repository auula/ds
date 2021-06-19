package quick_test

import (
	"github.com/higker/ds/sort/quick"
	"testing"
)

func TestQuickSort(t *testing.T) {
	qs := quick.New()
	arrays := []float64{12, 322, 122, 1, 3, 4}
	qs.Sort(arrays)
	t.Log(arrays)
}
