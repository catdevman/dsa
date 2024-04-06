package merge

func Sort(a []int) {
	sorted := mergeSort(a)
	for i, v := range sorted {
		a[i] = v
	}
}

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left := mergeSort(a[:len(a)/2])
	right := mergeSort(a[len(a)/2:])
	return merge(left, right)
}

func merge(a []int, b []int) []int {
	res := []int{}
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
