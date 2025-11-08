package merge

import (
	"cmp"
)

func Sort[T cmp.Ordered](a []T) {
	sorted := mergeSort(a)
	for i, v := range sorted {
		a[i] = v
	}
}

func mergeSort[T cmp.Ordered](a []T) []T {
	if len(a) < 2 {
		return a
	}
	left := mergeSort(a[:len(a)/2])
	right := mergeSort(a[len(a)/2:])
	return merge(left, right)
}

func merge[T cmp.Ordered](a []T, b []T) []T {
	res := []T{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		res = append(res, a[i])
	}
	for ; j < len(b); j++ {
		res = append(res, b[j])
	}
	return res
}
