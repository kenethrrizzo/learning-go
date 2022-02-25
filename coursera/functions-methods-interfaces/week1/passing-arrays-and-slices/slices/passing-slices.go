package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	foo(a)
	fmt.Println(a)
}

func foo(sli []int) {
	sli[0] = sli[0] + 1
}
