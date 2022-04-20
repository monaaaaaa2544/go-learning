package ch2

// 必须大写才能被包外访问
func GenerateFib(n int) []int {
	a := 1
	b := 1
	arr := []int{}

	arr = append(arr, a)
	for i := 0; i < n; i++ {
		arr = append(arr, b)
		a, b = b, a+b
	}

	return arr

}
