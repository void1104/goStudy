package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

/**
- 这里在少量goroutine的时候可能没问题
- 但并发多了就会报fatal error: concurrent map writes错误
*/
func main() {
	wg3 := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg3.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			set(key, n)
			fmt.Printf("k=:%v, v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

/**
Go的sync包提供了一个开箱即用的并发安全版map-sync.Map。
开箱即用表示不用像内置的map一样使用make函数初始化就能直接使用。
*/

var smap = sync.Map{}

func main1() {
	wg4 := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg4.Add(1)
		go func(n int){
			key := strconv.Itoa(n)
			smap.Store(key, n)
			value, _ := smap.Load(key)
			fmt.Printf("k=%k, v:=%v\n", key, value)
			wg4.Done()
		}(i)
	}
}