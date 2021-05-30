package main

import "fmt"

func testStruct() {
	//1. 函数作为结构体字段
	type calc struct { // 定义结构类型
		mul func(x, y int) int // 函数类型字段
	}

	x := calc{
		mul: func(x, y int) int {
			return x * y
		},
	}
	fmt.Println(x)
}

// 2. 函数经通道传递
func testChannel() {
	c := make(chan func(int, int) int, 2)

	c <- func(x, y int) int {
		return x + y
	}

	println((<-c)(1, 2))
}
