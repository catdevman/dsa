package merge_test

import (
	"slices"
	"testing"

	"github.com/catdevman/dsa/merge"
)

func TestMergeSort(t *testing.T) {
	testcases := []struct {
		name string
		arr  []int
	}{
		{
			"sort small array",
			[]int{0, 2, 1, 3, 5, 4},
		},
		{
			"slightly larger than small array",
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			"already sorted",
			[]int{0, 1, 2, 3, 4, 5},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			merge.Sort(testcase.arr)
			if !slices.IsSorted(testcase.arr) {
				t.Fail()
			}

		})
	}
}
