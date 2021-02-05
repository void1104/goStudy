package main

import "fmt"

/**
有缓冲的通道：
	1. 通道的容量表示通道中能存放元素的数量
	2. len函数获取通道内元素的数量，cap函数获取通道的容量
*/

// channel练习：优雅的从通道循环取值
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0-100的数量发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2接收值打印
	for i := range ch2 {
		fmt.Println(i)
	}
}
