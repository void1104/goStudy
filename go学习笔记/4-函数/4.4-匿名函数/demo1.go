package main

import "fmt"

func main()  {
	// 1. 直接运行
	func(s string){
		fmt.Println(s)
	}("hello world")

	// 2. 赋值给变量
	add := func(x, y int) int {
		return x + y
	}
	println(add(1,2))

	// 3. 作为参数
	test1(func() {
		println("作为参数 hello world")
	})

	// 4. 作为返回值
	add2 := test2()
	println(add2(1,2))
}


func test1(f func()){

}

func test2() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}


