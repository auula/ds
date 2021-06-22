package merge_test

import (
	"testing"

	"github.com/higker/ds/sort/merge"
)

func TestMergeSort(t *testing.T) {
	arrays := []float64{12, 322, 122, 1, 3, 4}
	ms := merge.New()
	ms.Sort(arrays)
	t.Log(arrays)
}
