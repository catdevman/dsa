package radixsort_test

import (
	"slices"
	"testing"

	"github.com/catdevman/dsa/radixsort"
)

func TestRadixSort(t *testing.T) {

	testcases := []struct {
		name string
		test []int
	}{
		{
			"sort small array",
			[]int{20, 23, 53, 771, 156, 248},
		},
		{
			"sort slightly larger than small array",
			[]int{10, 234, 30, 10923, 56579, 42343, 87454, 7252, 6000, 999},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res := testcase.test
			radixsort.Sort(res)
			if !slices.IsSorted(res) {
				t.Fail()
			}
		})
	}
}
