package ch2

import "testing"

// 斐波那契数列
func TestFib(t *testing.T) {
	a := 1 //类型推断的写法
	b := 1

	t.Log(a)
	for i := 0; i < 6; i++ {
		t.Log(" ", b)
		a, b = b, a+b
	}
}
