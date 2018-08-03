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
	func main() {
		m := make(map[string]int)
	
		m["Answer"] = 42
		fmt.Println("The value:", m["Answer"])
	
		m["Answer"] = 48
		fmt.Println("The value:", m["Answer"])
	
		delete(m, "Answer")
		fmt.Println("The value:", m["Answer"])
	
		v, ok := m["Answer"]
		fmt.Println("The value:", v, "Present?", ok)
	}
	//fmt.Println(foo1(1, 10))
	//fmt.Println(foo2(1, 2))
}
