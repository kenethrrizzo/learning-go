package main

func main() {
	x := 5
	addOne(&x)
	println(x)

	sum, message := multipleReturns(4, 5)
	println(sum)
	println(message)
}

func addOne(y *int) {
	*y = *y + 1
}

func multipleReturns(x, y int) (int, string) {
	return x + y, "successfully!"
}
