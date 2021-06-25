package count

import "testing"

func TestCountSort(t *testing.T) {
	arrays := []int{9, 8, 8, 8, 1, 2, 5, 6, 7}
	Sort(arrays)
	t.Log(arrays)
}
