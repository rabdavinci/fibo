package data

func FibonacciFor(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	b1 := 0
	b2 := 1
	for i := 0; i < n-1; i++ {
		b1, b2 = b2, b1+b2
	}
	return b2
}

func FibonacciRecursive(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-2)
}
