package logical

func and(a, b int) int {
	return a & b
}

func not(x int) int {
	return ^x
}

func andNot(a, b int) int {
	return a &^ b
}

func andNotNot(a, b int) int {
	return a &^ (^b)
}
