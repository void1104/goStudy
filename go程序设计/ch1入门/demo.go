package main

import "fmt"

func main() {
	i := 1
	//f := false
	//var true bool
	switch  {
	case i > -1:
		fmt.Println("-1")
		fallthrough
	case i > -2:
		fmt.Println("-2")
		fallthrough
	case i > -3:
		fmt.Println("-3")
	}
}
