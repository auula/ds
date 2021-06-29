package bucket_test

import (
	"testing"

	"github.com/higker/ds/sort/bucket"
)

func TestSort(t *testing.T) {
	arrays := []float64{12, 322, 123.1, 1, 3, 4}
	bucket.Sort(arrays, 10)
	t.Log(arrays)
}
