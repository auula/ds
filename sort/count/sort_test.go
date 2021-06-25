package count

import "testing"

func TestCountSort(t *testing.T) {
	arrays := []int{12, 3342, 551, 342, 112, 4535, 11}
	Sort(arrays)
	t.Log(arrays)
}
