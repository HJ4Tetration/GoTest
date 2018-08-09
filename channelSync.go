package main

import (
	"fmt"
)

/*func foo(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		c <- i
	}
}*/

func test(b []int) {
	b[0] = 100
	b = append(b, 4)
}

type I interface {
	M() int
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() int {
	fmt.Println(t.S)
	return 1
}

//func main() {
//a := [][]int{{1, 23}, {2}, {3}}
//b := a[:]
//fmt.Printf("%T\n%T\n", a, b)
//test(a)
//fmt.Println(a)

//}

/*func main() {
	c := make(chan int, 30)
	var wg sync.WaitGroup
	wg.Add(2)
	go foo(c, &wg)
	go foo(c, &wg)
	wg.Wait()
	close(c)

	for value, ok := <-c; ok; value, ok = <-c {
		fmt.Println(value)
	}
}*/

func foo(c chan int) {
	for i := 1; i < 10; i++ {
		c <- i
	}
	close(c)
}

/*func main() {
	c := make(chan int, 30)
	go foo(c)
	//go foo(c, &wg)

	for value, ok := <-c; ok; value, ok = <-c {
		fmt.Println(value)
	}
}*/
