package main

import (
	"fmt"
)

type I interface {
	test()
}

type V struct {
	x, y int
}

func (v *V) test() {
	fmt.Println((*v).y)
}

func foo1(x int, y int) (int, int) {
	return x + y, x
}

func foo2(x int, y int) int {
	return x + y
}

//var a = make([]int, 2, 2)

func ts(t []int) {
	t[0] = 100
}

var a = make([]int, 7, 8)

func Test(slice []int) {
	fmt.Println(len(slice))
	slice = slice[:8]
	fmt.Println(len(slice))
	slice = append(slice[:7], 100)
	fmt.Println(slice)
}

func main() {
	for i := 0; i < 7; i++ {
		a[i] = i
	}

	Test(a)
	fmt.Println(a)
}

//func main() {
//a := []byte{2, 2}
//a := make([]byte, 2)
//a[0] = 1
//fmt.Printf("%v", a)
//ts(a)
//fmt.Printf("%v", a)
/*st := V{1, 2}
var interf I
interf = &st
interf.test()*/

//var st *V
//st.x = 1
/*a := [8]byte{4, 4, 4, 4, 4, 4, 4, 4}
for i, v := range a {
	fmt.Println(i, v)
}*/
/*a := make([]byte, 0, 1)
fmt.Println(cap(a))
a[0] = 1
a.append(1)
a = a[:cap(a)]
fmt.Println(cap(a))*/

/*var i int = 3
if i == 3 {
	fmt.Println(2)
}*/

//}

//fmt.Println(foo1(1, 10))
//fmt.Println(foo2(1, 2))
