package main

import (
	"fmt"
)

var arr [5]int
var res int

func main() {
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
		res = res + arr[i]
	}
	fmt.Print(res)
}
