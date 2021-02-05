package main

import (
	"fmt"
	"sync"
)

var x1 int64
var wg1 sync.WaitGroup
var lock1 sync.Mutex

/**
- 互斥锁是一种常用的控制共享资源访问的方法，它
- 能够保证同时只有一个goroutine可以访问共享资源。
*/
func add1() {
	for i := 0; i < 5000; i++ {
		lock1.Lock()
		x = x + 1
		lock1.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2) // 计数器+2
	go add1()
	go add1()
	wg.Wait()		// 阻塞直到计数器归零
	fmt.Println(x)
}
