package minheap_test

import (
	"cmp"
	"container/heap"
	"fmt"
	"reflect"
	"testing"

	"github.com/catdevman/dsa/minheap"
)

// Test case structure for generic tests
type testCase[T cmp.Ordered] struct {
	name      string
	initial   []T
	pushes    []T
	pops      int
	poppedSeq []T
}

func TestMinHeap(t *testing.T) {
	// --- Integer Test Cases ---
	intTests := []testCase[int]{
		{
			name:      "SimpleIntInitAndPop",
			initial:   []int{5, 2, 8, 1, 9},
			pushes:    []int{},
			pops:      2,
			poppedSeq: []int{1, 2},
		},
		{
			name:      "IntInitAndPush",
			initial:   []int{10, 20},
			pushes:    []int{5, 15},
			pops:      3,
			poppedSeq: []int{5, 10, 15},
		},
		{
			name:      "EmptyHeapPushPop",
			initial:   []int{},
			pushes:    []int{3, 1, 2},
			pops:      3,
			poppedSeq: []int{1, 2, 3},
		},
	}

	// --- Float64 Test Cases ---
	floatTests := []testCase[float64]{
		{
			name:      "FloatInitAndPushPop",
			initial:   []float64{3.14, 2.71, 0.5, 1.61},
			pushes:    []float64{0.1},
			pops:      5,
			poppedSeq: []float64{0.1, 0.5, 1.61, 2.71, 3.14},
		},
	}

	// --- String Test Cases (Lexicographical Min) ---
	stringTests := []testCase[string]{
		{
			name:      "StringInitAndPushPop",
			initial:   []string{"zebra", "apple", "fig", "carrot"},
			pushes:    []string{"ant"},
			pops:      3,
			poppedSeq: []string{"ant", "apple", "carrot"},
		},
	}

	// Execute the tests generically for all defined types
	runTests(t, intTests)
	runTests(t, floatTests)
	runTests(t, stringTests)
}

// runTests is a generic function that handles the test execution logic.
func runTests[T cmp.Ordered](t *testing.T, tests []testCase[T]) {
	for _, tt := range tests {
		// Use t.Run to isolate tests and name them clearly (e.g., "SimpleIntInitAndPop/int")
		t.Run(fmt.Sprintf("%s/%T", tt.name, tt.initial), func(t *testing.T) {

			// 1. Setup and Initialization
			allElements := append(append([]T{}, tt.initial...), tt.pushes...)
			h := minheap.MinHeap[T](allElements)
			heap.Init(&h)

			// 2. Popping and sequence check
			actualPoppedSeq := make([]T, 0, tt.pops)
			for i := 0; i < tt.pops; i++ {
				if h.Len() == 0 {
					t.Fatalf("Heap underflow: tried to pop %d elements but heap is empty.", tt.pops)
				}
				v := heap.Pop(&h).(T) // Pop the element and assert its type
				actualPoppedSeq = append(actualPoppedSeq, v)
			}

			// 3. Verify the sequence of elements popped
			if !reflect.DeepEqual(actualPoppedSeq, tt.poppedSeq) {
				t.Errorf("Popped sequence mismatch.\nExpected: %v\nActual:   %v", tt.poppedSeq, actualPoppedSeq)
			}
		})
	}
}
