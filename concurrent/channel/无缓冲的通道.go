package main

import "fmt"

/**
无缓冲的通道，又称为阻塞通道
*/

func main1() {
	ch := make(chan int)
	ch <- 10 // 形成死锁：无缓冲的通道必须有人接收才能发送
	fmt.Println("发送成功")
}

func recv(c chan int) {
	ret := <-c // 会阻塞直到channel中有一个值
	fmt.Println("接收成功", ret)
}

/**
使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。
因此，无缓冲通道也被称为同步通道。
*/
func main() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10    // 会阻塞直到goroutine取出一个值
	fmt.Println("发送成功")
}
