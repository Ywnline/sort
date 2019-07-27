package sort

import (
	"testing"
)


func TestSort(t *testing.T) {
	li := []int{1, 5, 32, 2, 3, 23, 12, 4, 22}

	t.Log(bubbleSort(li))
}
