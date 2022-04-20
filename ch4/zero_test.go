package ch4

import "testing"

// &^按位置零

// iota初始值为0，每新增一行常量声明将使 iota 计数一次
const (
	// << 左移
	Readable   = 1 << iota //01
	Writable               //10
	Executable             //100
)

func TestConst(t *testing.T) {
	a := 7            //0111
	a = a &^ Readable //0110

	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
