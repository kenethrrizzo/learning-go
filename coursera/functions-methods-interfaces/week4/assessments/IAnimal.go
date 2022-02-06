package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	GetName() string
}

type Cow struct {
	name, food, locomotion, noise string
}
func (c *Cow) Eat() {
	fmt.Println(c.food)
}
func (c *Cow) Move() {
	fmt.Println(c.locomotion)
}
func (c *Cow) Speak() {
	fmt.Println(c.noise)
}
func (c *Cow) GetName() string {
	return c.name
}

type Bird struct {
	name, food, locomotion, noise string
}
func (b *Bird) Eat() {
	fmt.Println(b.food)
}
func (b *Bird) Move() {
	fmt.Println(b.locomotion)
}
func (b *Bird) Speak() {
	fmt.Println(b.noise)
}
func (b *Bird) GetName() string {
	return b.name
}

type Snake struct {
	name, food, locomotion, noise string
}
func (s *Snake) Eat() {
	fmt.Println(s.food)
}
func (s *Snake) Move() {
	fmt.Println(s.locomotion)
}
func (s *Snake) Speak() {
	fmt.Println(s.noise)
}
func (s *Snake) GetName() string {
	return s.name
}

// exec

func main() {
	var animalsCreated []Animal
	for {
		fmt.Print("\n>")
		reader := bufio.NewReader(os.Stdin)
		request, _ := reader.ReadString('\n')
		requestSplited := strings.Split(request, " ")
		if strings.TrimSpace(requestSplited[0]) == "newanimal" {
			animalsCreated = append(animalsCreated, createNewAnimal(
				strings.TrimSpace(requestSplited[1]), strings.TrimSpace(requestSplited[2])))
		} else if (strings.TrimSpace(requestSplited[0]) == "query") {
			query(requestSplited[1], strings.TrimSpace(requestSplited[2]), animalsCreated)
		} else {
			fmt.Println("Command not defined!")
		}
	}
}

func createNewAnimal(name, animalType string) Animal {
	var newAnimal Animal
	switch animalType {
	case "cow":
		newAnimal = &Cow{name: name, food: "grass", locomotion: "walk", noise: "moo"}
	case "bird":
		newAnimal = &Bird{name: name, food: "worms", locomotion: "fly", noise: "peep"}
	case "snake":
		newAnimal = &Snake{name: name, food: "mice", locomotion: "slither", noise: "hsss"}
	}
	fmt.Println("Created it!")
	return newAnimal
}

func searchAnimal(animalName string, animals []Animal) Animal {
	for _, animal := range animals {
		if animal.GetName() == animalName {
			return animal
		}
	}
	return nil
}

func query(animalName, animalInfo string, animals []Animal) {
	animal := searchAnimal(animalName, animals)
	if animal != nil {
		switch animalInfo {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Action not defined!")
		}
	} else {
		fmt.Println("Animal doesn't exists.")
	}
}