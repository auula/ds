package selection_test

import (
	"testing"

	"github.com/higker/ds/sort/selection"
)

func TestSelectSort(t *testing.T) {
	arrays := []float64{1, 22, 33.31, 223, 411}
	ss := selection.New()
	ss.Sort(arrays)
	t.Log(arrays)
}
