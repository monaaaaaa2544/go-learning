package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		// 获取命令行参数
		fmt.Println("Hello", os.Args[1])
	}
	os.Exit(-1) //返回退出状态
}
