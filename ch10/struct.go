package ch10

import "fmt"

// 定义一个结构体
type Stu struct {
	id   int
	name string
	age  int
}

func main() {

	e1 := Stu{1, "Tom", 20}
	e2 := new(Stu) //e2是一个指针
	fmt.Println(e1, e2)

	printStu(e1)
	printStus(e2)
}

func printStu(s Stu) {
	fmt.Println(s.id, s.name, s.age)
}

// 推荐用指针
func printStus(s *Stu) {
	fmt.Println(s.id, s.name, s.age)
}
