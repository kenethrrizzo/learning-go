package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	fmt.Println(foo(a))
}

func foo(x [3]int) int {
	return x[0]
}
