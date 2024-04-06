package quicksort_test

import (
	"slices"
	"testing"

	"github.com/catdevman/dsa/quicksort"
)

func TestQuickSort(t *testing.T) {
	testcases := []struct {
		name string
		test []int
	}{
		{
			"sort small array",
			[]int{0, 2, 3, 1, 5, 4},
		},
		{
			"sort slightly larger than small array",
			[]int{0, 2, 3, 1, 5, 4, 8, 7, 6, 9},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			quicksort.Sort(testcase.test)
			if !slices.IsSorted(testcase.test) {
				t.Fail()
			}
		})
	}
}
