package data

func Fibonacci(x, y int) FiboList {
	var list = FiboList{}
	if x < 0 || y < 0 || x > y {
		return list
	}
	if x == 0 && y == 0 {
		return append(list, 0)
	}

	n1 := 0
	n2 := 1

	for i := 0; i < y; i++ {
		list = append(list, n1)
		n1, n2 = n2, n1+n2
	}

	return list[x:]
}
