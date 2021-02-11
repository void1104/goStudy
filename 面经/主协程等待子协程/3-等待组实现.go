package main

import (
	"fmt"
	"sync"
)

/**
3 等待组实现
*/

var wg sync.WaitGroup

func cal2(no int) {
	fmt.Printf("%d 计算中\n", no)
	wg.Done() // 计时器-1
}

func main() {
	wg.Add(5) // 计时器+5
	for i := 0; i < 5; i++ {
		go cal2(i)
	}
	wg.Wait()
	fmt.Println("CAL END...")
}
