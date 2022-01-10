package tests

import (
	"testing"

	"github.com/rabdavinci/fibo/data"
)

func getCases() map[int]int {
	cases := map[int]int{
		-1: -1,
		0:  0,
		1:  1,
		2:  1,
		3:  2,
		10: 55,
		20: 6765,
	}
	return cases
}

func TestFor(t *testing.T) {

	for i, v := range getCases() {
		n := data.FibonacciFor(i)
		if n != v {
			t.Errorf("Expected %d, but got %d", v, n)
		}
	}

}
func TestRecursive(t *testing.T) {
	for i, v := range getCases() {
		n := data.FibonacciRecursive(i)
		if n != v {
			t.Errorf("Expected %d, but got %d", v, n)
		}
	}
}
