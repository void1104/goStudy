package main

import "fmt"

type User interface {

}

type Student struct {

}

func main() {
	var u User
	var a int
	u = a
	fmt.Println(u == nil)
}