package main

import "fmt"

func main() {
	defer fmt.Println("Bye!") //se ejecutan al final de la funcion
	fmt.Println("Hello!")
	maxValue1 := GetMax(1, 3, 4, 5, 3, 2, 3)
	fmt.Println(maxValue1)
	vslice := []int{1, 2, 3, 4, 5, 6, 4, 3, 3, 4, 6, 7, 8}
	maxValue2 := GetMax(vslice...)
	fmt.Println(maxValue2)
}

func GetMax(vals ...int) int { //para admitir varios argumentos
	maxV := -1
	for _, v := range vals {
		if v > maxV {
			maxV = v
		}
	}
	return maxV
}
