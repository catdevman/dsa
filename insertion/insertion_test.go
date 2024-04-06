package insertion_test

import (
	"slices"
	"testing"

	"github.com/catdevman/dsa/insertion"
)

func TestInsertionSort(t *testing.T) {

	arrs := [][]int{
		{0, 1, 2, 3, 4},
		{4, 3, 2, 1, 0},
		{3, 2, 0, 4, 1},
	}

	for _, a := range arrs {
		insertion.Sort(a)
		if !slices.IsSorted(a) {
			t.Fatal("Failed to sort")
		}
	}

}
