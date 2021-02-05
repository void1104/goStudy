package main

import (
	"fmt"
	"sync"
)

/**
二，
	直接用time.Sleep()肯定是不合适的，可以使用sync.WaitGroup来实现并发任务的同步。
*/

var wg1 sync.WaitGroup // 需要注意sync.WaitGroup是一个结构体，传递时要传递指针

func hello2() {
	defer wg1.Done() // 计数器-1
	fmt.Println("Hello Goroutine!")
}

func main() {
	wg1.Add(1)  // 计数器+1
	go hello2() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wg1.Wait() // 阻塞直到计数器变为0
}
