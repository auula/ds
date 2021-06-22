package shell_test

import (
	"testing"

	"github.com/higker/ds/sort/shell"
)

func TestShellSort(t *testing.T) {
	arrays := []float64{1, 22, 33.31, 223, 41, 123, 111, 888}
	ss := shell.New()
	ss.Sort(arrays)
	t.Log(arrays)
}
