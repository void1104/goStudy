package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

/**
GPM可视化调试
*/

/*
trace的创建过程
 1 创建文件
 2 启动
 3 停止
*/

func main() {
	//1. 创建一个trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//2. 启动trace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	// 正常要调试的业务
	fmt.Println("hello world!")

	//3. 停止trace
	trace.Stop()
}
