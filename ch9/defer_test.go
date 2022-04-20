package ch9

import "testing"

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clear resources")
	}()

	t.Log("Start")
	panic("error") // 这里会触发defer
}
