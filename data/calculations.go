package data

import (
	"math/big"
)

func Fibonacci(x, y int64) string {
	a := big.NewInt(0)
	b := big.NewInt(1)

	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(y), nil)
	for a.Cmp(&limit) < 0 {
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
	}
	return a.String()
}
