package main

import (
	"fmt"
)

func foo1(x int, y int) (int, int) {
	return x + y, x
}

func foo2(x int, y int) int {
	return x + y
}

func main() {
	var i int = 3
	if i == 3 {
		fmt.Println(2)
	}
	//fmt.Println(foo1(1, 10))
	//fmt.Println(foo2(1, 2))
}
