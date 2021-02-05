package main

import "fmt"

/**
- 在某些场景下我们需要同步从多个通道接收数据。通道在接收数据时，如果没有数据可以接收将会发生阻塞。
*/

/**
- 这种方式也可以实现从多个通道接收值的需求，但是运行性能会差很多。
*/
func demo(ch1, ch2 chan string) {
	for {
		// 尝试从ch1接收值
		data1, ok1 := <-ch1
		// 尝试从ch2接收值
		data2, ok2 := <-ch2
		fmt.Println(data1, ok1)
		fmt.Println(data2, ok2)
	}
}

/**
- 为了应对这种场景，Go内置了select关键字，可以同时响应多个通道的操作。
- select的使用类似switch语句
- 使用select语句能提高代码的可读性。如果多个case同时满足，select会随机选择一个。对于没有case的select{}会一直等待
*/
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
