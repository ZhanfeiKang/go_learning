package cal

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func getsub(n1 int, n2 int) int {
	return n1 - n2
}
