package type_test

import "testing"

// go无法进行隐式类型转换 只能进行显式类型转换
func TestType(t *testing.T) {
	a := 3
	b := "3"
	// b=a //error

	b = "8"
	t.Log(a, b)
}

// go不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a //a指针（a的地址）
	// aPtr = aPtr + 1 //error

	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}
