package factorial_test

import (
	"fmt"
	"testing"

	"github.com/catdevman/dsa/factorial"
)

func TestFactorial(t *testing.T) {
	testcases := []struct {
		name     string
		fact     int
		expected int
	}{
		{
			"factorial of 5",
			5,
			120,
		},
		{
			"factorial of 2",
			2,
			2,
		},
		{
			"factorial of 1",
			1,
			1,
		},
		{
			"factorial of 10",
			10,
			3628800,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res := factorial.Of(testcase.fact)

			if res != testcase.expected {
				t.Fatal(fmt.Sprintf("got %d but expected %d", res, testcase.expected))
			}
		})
	}

}
