package main

import "fmt"

/**
- 3种可行方法：
	1 channel通信
	2 select
	3 等待组
*/

/**
1 channel通信实现
*/

var sync1 = make(chan int)

// 实际任务
func cal(no int) {
	fmt.Printf("%d 计算中\n", no)
	sync1 <- no
}

func main() {
	count := 5
	for i := 0; i < 5; i++ {
		go cal(i)
	}
	// 在channel未关闭之前，主线程不会结束
	for range sync1 {
		count--
		if count == 0 {
			close(sync1)
		}
	}
	fmt.Println("cal end...")
}
