package main

import (
	"strings"
	//"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ans := make(map[string]int)
	for _, e := range strings.Fields(s) {
		ans[e]++
	}
	return ans
	//return map[string]int{"x": 1}
}

/*func main() {
	wc.Test(WordCount)
}*/
