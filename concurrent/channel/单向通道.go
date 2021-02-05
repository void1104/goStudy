package main

import "fmt"

/**
有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道
都会对其进行限制，比如只能发送或只能接收。Go语言提供了单向通道来处理这种情况
*/

func counter(out chan<- int) { // 只能发送的通道
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	// 在函数传参及任何赋值操作中将双向通道转换为单向通道是可以的，但反过来不行
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go square(ch2, ch1)
	printer(ch2)
}
