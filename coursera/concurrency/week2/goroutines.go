package main

import "fmt"

/*
	The race condition itself is the shared resources between the two functions
	that are executed in the "main", in this case the x & y variables. Initially
	they are being initialized with the values 10 and 2 respectively, but within
	the goroutines other values are being assigned. This affects the result of the
	execution, different values are printed.
*/
func main() {
	for i := 0; i < 100; i++ {
		x, y := 10, 2

		go func(v *int) {
			*v = 60
		}(&x)

		go func(v *int) {
			*v = 3
		}(&y)

		func(v1 int, v2 int) {
			fmt.Println(v1 / v2)
		}(x, y)
	}
}
