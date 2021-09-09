package D2021_09

/*
斐波那契
 */

func Lc904Fib(n int) int {
	if n == 0 {
		return 0
	}
	a, b := 0, 1
	for i := 0; i < n - 1; i++ {
		a, b = b, (a + b) % (1e9 + 7)
	}
	return b
}
