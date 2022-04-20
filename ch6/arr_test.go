package ch6

import "testing"

func TestArray(t *testing.T) {
	arr1 := [...]int{1, 3, 4, 5, 6, 7, 8, 9}
	arr1[0] = 0

	t.Log(arr1)

	// 数组遍历
	for idx, e := range arr1 {
		t.Log(idx, e)
	}
	// 数组切片
	t.Log(arr1[:3])
}
