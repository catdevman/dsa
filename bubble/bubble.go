package bubble

func Sort(nums []int) {
	swap := true
	for swap {
		swap = false
		for i, n := range nums {
			if i != len(nums)-1 {
				if n > nums[i+1] {
					nums[i], nums[i+1] = nums[i+1], n
					swap = true
				}
			}
		}
	}
}
