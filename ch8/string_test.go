package ch8

import "testing"
import "strings"
import "strconv"

// 常用字符串函数：strings、strconv
// String
func TestMap(t *testing.T) {
	s := "sdsfd,afasd"
	t.Log(len(s), s[4])

	// s[1] = "9" //error

	s1 := "中"
	c := []rune(s1)

	t.Logf("UTF8: %x", s1)
	t.Logf("unicode: %x", c[0])

	// 字符串分割
	parts := strings.Split(s, ",")
	t.Log(parts)

	// 字符串拼接
	t.Log(strings.Join(parts, "-"))

	// 字符串类型转换
	n, _ := strconv.Atoi("123") // 字符串转整型
	t.Log(n + 10)
	a := strconv.Itoa(123) // 整型转字符串
	t.Log("str" + a)
}
