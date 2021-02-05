package main

import (
	"fmt"
	"sync"
)

var wg2 sync.WaitGroup

func hello3(i int) {
	defer wg2.Done()
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		// 会发现每次打印数字的顺序不一样，这是因为10个goroutine是并发执行的，而goroutine的调度是随机的。
		go hello3(i)
	}
	wg2.Wait()
}
