package insertion_test

import (
	"testing"

	"github.com/higker/ds/sort/insertion"
)

func TestInsertionSort(t *testing.T) {
	arrays := []float64{12, 322, 123.1, 1, 3, 4}
	insertion := insertion.New()
	insertion.Sort(arrays)
	t.Log(arrays)
}
