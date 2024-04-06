package factorial

func Of(n int) int {
	if n == 1 {
		return 1
	}
	return n * Of(n-1)
}
