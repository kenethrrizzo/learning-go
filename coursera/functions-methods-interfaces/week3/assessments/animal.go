package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) InitData(food, locomotion, noise string) {
	a.food = food
	a.locomotion = locomotion
	a.noise = noise
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

// execution

func main() {
	var a Animal
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		request, _ := reader.ReadString('\n')
		requestSplited := strings.Split(request, " ")
		animal := strings.TrimSpace(requestSplited[0])
		action := strings.TrimSpace(requestSplited[1])
		initData(&a, animal)
		printAction(&a, action)
	}
}

func initData(a *Animal, animal string) {
	if animal == "cow" {
		a.InitData("grass", "walk", "moo")
	} else if animal == "bird" {
		a.InitData("worms", "fly", "peep")
	} else if animal == "snake" {
		a.InitData("mice", "slither", "hsss")
	} else {
		err := errors.New("Animal not defined")
		panic(err)
	}
}

func printAction(a *Animal, action string) {
	if action == "eat" {
		a.Eat()
	} else if action == "move" {
		a.Move()
	} else if action == "speak" {
		a.Speak()
	} else {
		err := errors.New("Action not defined")
		panic(err)
	}
}
