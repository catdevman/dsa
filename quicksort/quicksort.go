package quicksort

func Sort(a []int) {
	sorted := quickSort(a)
	for i, v := range sorted {
		a[i] = v
	}
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	pivot := a[len(a)-1]
	left := []int{}
	right := []int{}
	for _, v := range a[:len(a)-1] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	pivArr := []int{pivot}
	fullRight := append(pivArr, right...)
	return append(left, fullRight...)
}
