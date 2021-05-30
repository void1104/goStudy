package main

import (
	"fmt"
	"sync"
)

type SingleTon struct{}

var singleton *SingleTon
var once sync.Once

func getSingleTon() *SingleTon {
	once.Do(func() {
		singleton = new(SingleTon)
	})
	return singleton
}

func main() {
	var wg sync.WaitGroup
	var t sync.Map
	//t[1] = "s"
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			obj := getSingleTon()
			fmt.Println(obj)
			wg.Done()
		}()
	}
	wg.Wait()
}
