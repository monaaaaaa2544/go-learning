package ch7

import "testing"

// Map
func TestMap(t *testing.T) {
	m1 := map[string]int{"asdf": 1, "df": 3}
	t.Log(m1["df"])

	m2 := make(map[string]int, 10)
	t.Logf("%T, %v", m2, m2)
	t.Log(m2, len(m2))

	if _, ok := m1["3"]; ok {
		t.Log("key 3 is exist")
	} else {
		t.Log("key 3 is not exist")
	}
}

// Map遍历
func TestMapTrave(t *testing.T) {
	m1 := map[string]int{"sasdfasdfasdf": 3, "asdfad": 56}

	for k, v := range m1 {
		t.Log(k, v)
	}
}
