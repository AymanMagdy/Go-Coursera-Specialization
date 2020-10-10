package main

import (
	"fmt"
	"strings"
)

func getFirstChar(userInput string) bool {
	var firstElement = userInput[0]
	if firstElement == byte('i') || firstElement == byte('I') { 
		return true
	}
	return false
}

func getLastChar(userInput string) bool {
	var lastChar byte
	var length = len(userInput)
	lastChar = userInput[length-1]
	if lastChar == byte('n') || lastChar == byte('N') {
		return true
	}
	return false
}

func serachLetterA(userInput string) bool {
	if strings.Contains(userInput, "a") { 
		return true
	}
	return false
}

func main()  {
	var userInput string
	var result = "Not Found!"
	fmt.Println("Enter a string: ")
	fmt.Scanf("%s", &userInput)
	if getFirstChar(userInput) && getLastChar(userInput)  && serachLetterA(userInput) {
		result = "Found!"
	}
	fmt.Println(result)
}