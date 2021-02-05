package main

import (
	"fmt"
	"time"
)

/**
概念：
	1. goroutine类似于线程，属于用户态的线程
	2. goroutine是由Go语言的运行时（runtime）调度完成，而线程是由操作系统调度完成。
	   Go语言还提供channel在多个goroutine间进行通信。goroutine和channel
	   是Go语言秉承的CSP(Communicating Sequential Process)并发模式的重要实现基础。
	3. Go语言会智能地将goroutine中的任务合理分配给每个CPU，其在语言层面内置了调度和上下文切换
*/

/**
一，
	hello()函数很有可能不会执行，在程序启动时Go程序就会为main()函数创建一个
	默认的goroutine。当main()函数返回的时候该goroutine就结束了，所有在main()
	函数中启动的goroutine会一同结束。
*/

func hello1() {
	fmt.Println("Hello Goroutine!")
}

func main() {
	go hello1()
	time.Sleep(time.Second)
	fmt.Println("main goroutine done!")
}
