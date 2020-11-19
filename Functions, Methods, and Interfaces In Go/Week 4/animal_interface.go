package main

import (
	"fmt"
)

// The animal type with the needed strings needed
type animal struct {
	food 	  string
	locmotion string
	noise     string	
}

// The animal interface with the needed methods.
type animalInterface interface {
	Eat()
	Move()
	Speak()
}

// The Eat, Move and Speak methods
func (ani animal) Eat() {
	fmt.Println(ani.food)
	return
}

func (ani animal) Move() {
	fmt.Println(ani.locmotion)
	return
}

func (ani animal) Speak() {
	fmt.Println(ani.noise)
	return
}


func main() {
	animalMap := make(map[string]animal)

	animalMap["cow"] = animal{"grass", "walk", "moo"}
	animalMap["bird"] = animal{"worm", "fly", "peep"}
	animalMap["snake"] = animal{"mice", "slither", "hss"}

	var generalAnimal animalInterface

	for {
		var command, requestAnimal, requestType string
		fmt.Print(">")
		fmt.Scan(&command, &requestAnimal, &requestType)

		if command == "query" {
			generalAnimal = animalMap[requestAnimal]

			switch requestType {
			case "eat":
				generalAnimal.Eat()
			case "move":
				generalAnimal.Move()
			case "speak":
				generalAnimal.Speak()
			}
		}

		if command == "newanimal" {
			animalMap[requestAnimal] = animalMap[requestType]
			fmt.Println("Created a new animal")
		}
	}
}
