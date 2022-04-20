package ch5

import "testing"

func TestIf(t *testing.T) {
	if a := 1; a == 1 {
		t.Log("a==1")
	} else {
		t.Log("a!=1")
	}
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 6; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCase(t *testing.T) {
	for i := 0; i < 6; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("what?")
		}
	}
}
