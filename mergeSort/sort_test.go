package sort

import (
	"testing"
)

func TestSort(t *testing.T) {
	li := []int{8, 9, 5, 7, 1, 2, 5, 7, 6, 3, 5, 4, 8, 1, 8, 5, 3, 5, 8, 4}
	t.Log(mergeSort(li))
}
