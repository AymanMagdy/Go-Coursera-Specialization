package main

import (
	"errors"
	"fmt"
)

// The Named interface describes any entity that can have a name
type Named interface {
	getName() string
}

// The animal interface, prints the animal eat, move and sound
type Animal interface {
	Named
	Eat()
	Move()
	Speak()
}

// The Cow type
type Cow struct {
	name string
}

// The Bird type
type Bird struct {
	name string
}

type Snake struct {
	name string
}

// The Named interface with the needed methods for the animals.

// The cow methods for returning the name
func (cow Cow) getName() string {
	return cow.getName()
}

// The bird method for returning the name
func (bird Bird) getName() string {
	return bird.getName()
}

// The snake method for returning the name
func (snake Snake) getName() string {
	return snake.getName()
}


// The Eat method for the different types of animals {cow, bird and snake}
func (cow Cow) Eat() {
	fmt.Println("Grass")
}

func (bird Bird) Eat() {
	fmt.Println("Worms")
}

func (snake Snake) Eat() {
	fmt.Println("Mice")
}

// The Move method for the different types of animals {cow, bird and snake}
func (cow Cow) Move() {
	fmt.Println("Walk")
}

func (bird Bird) Move() {
	fmt.Println("Fly")
}

func (snake Snake) Move() {
	fmt.Println("Slither")
}

// The Speak method for the different types of animals {cow, bird, snake}
func (cow Cow) Speak() {
	fmt.Println("Moo")
}

func (bird Bird) Speak() {
	fmt.Println("Peep")
}

func (snake Snake) Speak() {
	fmt.Println("Hsss")
}

// Checking the input from the user
func isValidCommand(animal string) bool {
	return command == "newanimal" || command == "query"
}

func isValidAnimal(animal string) bool {
	return animal == "cow" || animal == "bird" || animal == "snake"
}

func isValidAnimalProperty(property string) bool {
	return property == "eat" || property == "move" || property == "speak"
}

func getRequest() (string, string, string) {
	var command, name, option string
	for {
		fmt.Print(">")
		if _, err := fmt.Scanln(&command, &name, &option); err != nil {
			fmt.Println("Error:", err)
		} else {
			if isValidCommand(command) {
				if command == "newanimal" {
					if isValidAnimal(animal) {
						break
					} else {
						fmt.Println("Invalid animal type. Please use 'cow', 'bird' or 'snake'.")
					}
				} else if command == "query" {
					if isValidAnimalProperty(option) {
						break
					} else {
						fmt.Println("Invalid animal property. Please use 'cow', 'bird' or 'snake'.")
					}
				}
			} else {
				fmt.Println("Invalid command. Please use 'newanimal' or 'query'.")
			}
		}
	}

	return command, name, option
}

// Appening a new animal type
func appendAnimal(animals *[]Animal, animal Animal) {
	*animals = append(*animals, animal)
	fmt.Println("Created it!")
}

// Finding the animals with the given names.
func findAnimal(animals []Animal, name string) (Animal, error) {
	var foundAnimal Animal
	var err = errors.New("Animal with that given name is not found.")
	for _, animal := range animals {
		if Named(animal).getName() == name {
			foundAnimal = animal
			err = nil
			break
		}
	}

	return foundAnimal, err
}

// The main function
func main() {
	var animals []Animal
	for {
		command, name, option := getRequest()
		switch command {
		case "newAnimal":
			switch option {
			case "cow":
				appendAnimal(*animals, Cow{name})
			case "bird":
				appendAnimal(*animals, Bird{name})
			case "snake":
				appendAnimal(*animals, Snake{name})
			}
		case "query":
			var animal Animal
			var err error
			if Animal, err = findAnimal(animals, name); err != nil {
				fmt.Println(err)
				break
			}
			switch option {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
				
			}
		}
	}
}
