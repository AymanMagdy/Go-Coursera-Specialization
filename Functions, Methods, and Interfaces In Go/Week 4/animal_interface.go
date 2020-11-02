package main

import (
	"fmt"
)

// The Animal struct interface with 3 string fields
type Animal struct {
	food 		string
	moveType	string
	sound 		string
}

// The animal functions
func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.moveType)
}

func (animal Animal) Speak() {
	fmt.Println(animal.sound)
}


func main() {
	Cow := Animal{"Grass", "Walk", "Moo"}
	Bird := Animal{"Worms", "Fly", "Peep"}
	Snake := Animal{"Mice", "Slither", "Hss"}

	Cow.Eat()
	Bird.Move()
	Snake.Speak()
}