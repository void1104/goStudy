package main

func test3() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		s = append(s, func() { // 将多个匿名函数添加到列表
			println(&i, i) // 当main执行这些函数时，它们读取的是环境变量i最后一次循环时的值。
		})
	}
	return s
}

// 解决方法就是每次用不同的环境变量或传参复制，让各自闭包环境各不相同
func test4() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		x := i
		s = append(s, func() {
			println(&x, x)
		})
	}
	return s
}

func main() {
	for _, f := range test3() {
		f()
	}
}
