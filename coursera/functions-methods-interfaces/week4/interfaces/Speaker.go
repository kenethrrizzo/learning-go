package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

func (d *Dog) Speak() {
	if d == nil {
		fmt.Println("<speak>")
	} else {
		fmt.Println(d.name)
	}
}

func main() {
	var speaker Speaker
	var firstDog = &Dog{name: "woof"}
	speaker = firstDog
	speaker.Speak()
}