package main

import "fmt"

type IShape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	x, y int
}

func (t *Triangle) Area() float64 {
	return float64(t.x)
}
func (t *Triangle) Perimeter() float64 {
	return float64(t.y)
}

type Square struct {
	x, y int
}

func (s *Square) Area() float64 {
	return float64(s.x)
}
func (s *Square) Perimeter() float64 {
	return float64(s.y)
}

/* Exec */

func main() {
	var triangle = &Triangle{5, 5}
	fmt.Println(calculate(triangle))
	var square = &Square{5, 15}
	fmt.Println(calculate(square))
	typeSwitch(triangle)
	typeSwitch(square)
}

func calculate(shape IShape2D) bool {
	// type assertions for disambiguation
	triangle, ok := shape.(*Triangle)
	if ok {
		fmt.Println("I'm a triangle", triangle)
	} else {
		fmt.Println("I'm not a triangle :(")
	}
	return shape.Area() < 10 && shape.Perimeter() < 10
}

func typeSwitch(shape IShape2D) {
	switch sh := shape.(type) {
	case *Triangle:
		fmt.Println("I'm a triangle :)", sh)
	case *Square:
		fmt.Println("I'm a square :)", sh)
	default:
		fmt.Println("Idk who I am :(")
	}
}
