package main

import (
	"fmt"
	"sync"
	"time"
)

/**
- 读写互斥锁
*/

var (
	x2     int64
	wg2    sync.WaitGroup
	rwlock sync.RWMutex
)

func write() {
	rwlock.Lock() // 加写锁
	x2 = x2 + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解互斥锁
	wg2.Done()
}

func read() {
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg2.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go read()
	}
	wg2.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
