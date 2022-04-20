package ch5

import "testing"

func TestLoop(t *testing.T) {
	for n := 0; n < 10; n++ {
		t.Log(n)
	}

	// 无限循环
	// for {
	// 	t.Log("loop")
	// }

}
