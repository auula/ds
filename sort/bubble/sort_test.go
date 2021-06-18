package bubble_test

import (
	"testing"

	"github.com/higker/ds/sort/bubble"
)

func TestSort(t *testing.T) {
	arrays := []float64{12, 322, 123.1, 1, 3, 4}
	bubble := bubble.New('>')
	bubble.Sort(arrays)
	t.Log(arrays)
}
