package main

// 就这段代码而言，test返回的匿名函数会引用上下文环境变量x，
func test(x int) func() {
	return func() {
		println(x)
	}
}

// 当该函数在main中执行时，它依然可正确读取x的值，这种现象就称之为闭包
func main() {
	f := test(123)
	f()
}
