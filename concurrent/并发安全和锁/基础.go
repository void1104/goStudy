package main

import (
	"fmt"
	"sync"
)

/**
有时候GO代码会存在多个goroutine同时操作一个资源（临界区）
这种情况会发生竞态问题（数据竞争）。
*/

var x int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}


/**
这里开启了两个go程去累加变量x的值，这两个goroutine在访问和修改
x变量的时候就会存在数据竞争，导致最后的结果与期待的不符
 */
func main() {
	wg.Add(2) //
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
