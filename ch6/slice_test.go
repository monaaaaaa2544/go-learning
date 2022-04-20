package ch6

import "testing"

func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(s0, len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0))

	s2 := make([]int, 0, 5)
	t.Log(s2, len(s2), cap(s2))
	// t.Log(s2[0], s2[1])  //error
}

// 切片中的容量
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "2", "3", "4", "5"}

	Q1 := year[1:3]
	t.Log(Q1, len(Q1), cap(Q1))

	Q2 := year[1:2]
	t.Log(Q2, len(Q2), cap(Q2))

	Q2[0] = "Unknow"
	t.Log(year)
}
