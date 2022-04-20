package ch9

import "math/rand"

// 函数多个返回值
func MultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(30)
}

func sum(nums ...int) {
	for _, num := range nums {
		println(num)
	}
}

func main() {
	a, b := MultiValues()
	println(a, b)

	sum(1, 2, 3, 4, 5)
}
