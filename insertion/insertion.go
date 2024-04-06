package insertion

func Sort(a []int) {
	for i := 1; i < len(a); i++ {
		numToInsert := a[i]
		j := 0
		for j = i - 1; j >= 0 && a[j] > numToInsert; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = numToInsert
	}
}
