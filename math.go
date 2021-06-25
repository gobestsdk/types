package types

func Pow(x, n int) (y int) {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	return x * Pow(x, n-1)
}
