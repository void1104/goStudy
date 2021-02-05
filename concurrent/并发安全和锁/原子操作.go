package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/**
- 代码中的加锁操作因为涉及内核态的上下文切换会比较耗时，代价比较高。
- 针对基本数据类型我们还可以用原子操作来保证并发安全。
- 因为原子操作是GO语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好。
*/

// 示例：互斥锁和原子操作的比较

var x5 int64
var l sync.Mutex
var wg5 sync.WaitGroup

// 普通版加函数
func add5() {
	x5++
	wg5.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l.Lock()
	x5++
	l.Unlock()
	wg5.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x5, 1)
	wg5.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 90000009; i++ {
		wg5.Add(1)
		//go add5()
		go mutexAdd()		// 差别不大，是否内部有锁消除或优化（逃逸分析）
		//go atomicAdd()
	}
	wg5.Wait()
	end := time.Now()
	fmt.Println(x5)
	fmt.Println(end.Sub(start))
}
