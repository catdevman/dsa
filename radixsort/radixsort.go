package radixsort

import (
	"strconv"
)

type Sortable struct {
	val    int
	strVal string
	digits []int
}

func Sort(a []int) {
	sorted := radixSort(a)
	for i, v := range sorted {
		a[i] = v
	}
}

func radixSort(a []int) []int {
	buckets := [][]int{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}

	longest := getLongestNumber(a)

	for i := 0; i <= longest; i++ {
		for len(a) > 0 {
			var current int
			current, a = a[0], a[1:]
			idx := getDigit(current, i)
			buckets[idx] = append(buckets[idx], current)
		}

		for j := range buckets {
			for len(buckets[j]) > 0 {
				var val int
				val, buckets[j] = buckets[j][0], buckets[j][1:]
				a = append(a, val)
			}
		}
	}
	return a
}

func getDigits(i int) []int {
	s := strconv.Itoa(i)
	digits := []int{}

	for _, n := range s {
		digits = append(digits, int(n-'0'))
	}
	return digits
}

func getDigit(i int, place int) int {
	digits := getDigits(i)
	idx := len(digits) - (place + 1)
	if idx >= 0 && idx < len(digits) {
		return digits[idx]
	} else {
		return 0
	}

}

func getLongestNumber(a []int) int {
	longest := 0
	for _, v := range a {
		curr := len(getDigits(v))
		if curr > longest {
			longest = curr
		}
	}
	return longest
}
