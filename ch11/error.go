package main

import (
	"errors"
	"fmt"
)

func E(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n should be >= 0")
	}
	if n <= 1 {
		return nil, errors.New("n should be > 1")
	}
	return []int{1, 2, 3}, nil
}

func Panic(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("n should be >= 0")
	}
	if n <= 1 {
		return nil, errors.New("n should be > 1")
	}
	return []int{1, 2, 3}, nil
}

func main() {
	// normal error
	if res, err := E(1); err != nil {
		fmt.Println(err) // normal error
	} else {
		fmt.Println(res)
	}

	// panic error
	if res, err := E(1); err != nil {
		defer func() {
			fmt.Println("defer")
		}()
		panic(err) // panic error
	} else {
		fmt.Println(res)
	}

	// recover error
	// if res, err := E(1); err != nil {
	// 	defer func() {
	// 		// recover error
	// 		if err := recover(); err != nil {
	// 			fmt.Println("recover panic")
	// 		}
	// 	}()
	// 	panic(err) // panic error
	// } else {
	// 	fmt.Println(res)
	// }
}
