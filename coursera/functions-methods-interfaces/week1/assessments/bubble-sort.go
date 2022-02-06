package main

import "fmt"

func main() {
	/* Test one */
	testOne := []int{3, 32, 4, 2, 8, 4, 6, 2, 1, 16}
	BubbleSort(testOne)
	fmt.Println(testOne)
	/* Test two (negative numbers) */
	testTwo := []int{3, -32, 4, 2, -8, 4, 6, 2, 1, 16}
	BubbleSort(testTwo)
	fmt.Println(testTwo)
}

func Swap(sli []int, index int) {
	aux := sli[index]
	sli[index] = sli[index+1]
	sli[index+1] = aux
}

func BubbleSort(sli []int) {
	n := len(sli)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}
