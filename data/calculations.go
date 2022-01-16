package data

import (
	"fmt"
	"math/big"
)

func Fibonacci(x, y int64) (r string) {
	a := big.NewInt(0)
	b := big.NewInt(1)

	var i int64
	for i = 0; i < y; i++ {
		if i == x {
			r = a.String()
		} else if i > x {
			r = fmt.Sprintf("%v, %v", r, a.String())
		}
		// Compute the next Fibonacci number, storing it in a.
		a.Add(a, b)
		// Swap a and b so that b is the next number in the sequence.
		a, b = b, a
	}
	return r
}
