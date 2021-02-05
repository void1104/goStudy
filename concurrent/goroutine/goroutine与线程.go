package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
1.可增长的栈
	- OS线程（操作系统线程）一般都有固定的栈内存（通常为2MB）,一个goroutine的栈在其
	  生命周期开始时只有很小的栈（典型情况下2kb）,goroutine的栈不是固定的，他可以按
	  需增大和减小，goroutine的栈大小限制可以达到1GB，虽然极少会用到这个大小。
2.goroutine调度
	- OS线程是由OS内核来调度的，goroutine则是由Go运行时（runtime）自己的调度器调度的，
	  这个调度器使用一个称为m:n调度的技术（复用/调度m个goroutine到n个OS线程）。
      goroutine的调度不需要切换内核语境，所以调用一个goroutine比调度一个线程成本低很多。
3.GOMAXPROCS
	- Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个OS线程来同时执行GO代码。
	  默认值是机器上的CPU核心数。例如在一个8核心的机器上，调度器会把Go代码同时调度到8个OS线程上
	  （GOMAXPROCS是m:n调度中的n）。
*/

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go a()
	time.Sleep(time.Second)
	go b()
	time.Sleep(time.Second)
}

/**
总结：Go语言中的操作系统线程和goroutine的关系
	1. 一个操作系统线程对应用户态多个goroutine
	2. go程序可以同时使用多个操作系统线程
	3. goroutine和OS线程是多对多的关系，即m:n
*/
