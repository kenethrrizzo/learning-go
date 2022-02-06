package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(acceleration, velocity, initialDisplacement float64) func(float64) float64 {
	function := func(time float64) float64 {
		return 0.5*acceleration*math.Pow(time, 2) + velocity*time + initialDisplacement
	}
	return function
}

func main() {
	var acceleration, velocity, initialDisplacement, time float64
	fmt.Print("Type the acceleration: ")
	fmt.Scanln(&acceleration)
	fmt.Print("Type the velocity: ")
	fmt.Scanln(&velocity)
	fmt.Print("Type the initial displacement: ")
	fmt.Scanln(&initialDisplacement)
	fn := GenDisplaceFn(acceleration, velocity, initialDisplacement)
	fmt.Print("Type the time: ")
	fmt.Scanln(&time)
	fmt.Println("\nDisplacement:", fn(time))
}
