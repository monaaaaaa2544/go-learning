package ch2

import "testing"

// go语言常量

// iota初始值为0，每新增一行常量声明将使 iota 计数一次
const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota //左移
	Writable
	Executable
)

var (
	a = 1
	b = 2
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)

	a := 7 //0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
