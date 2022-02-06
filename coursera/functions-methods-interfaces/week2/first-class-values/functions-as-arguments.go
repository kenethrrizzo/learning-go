package main

import "fmt"

func ApplyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func AddOne(x int) int { return x+1 }

func SubstractOne(x int) int { return x-1 }

func main() {
	fmt.Println(ApplyIt(AddOne, 3))
	// Declare variable as a func
	var substract func(int) int
	substract = SubstractOne
	fmt.Println(ApplyIt(substract, 6))
	// Anonymous functions
	fmt.Println(ApplyIt(func (x int) int { return x*2 }, 5))
}