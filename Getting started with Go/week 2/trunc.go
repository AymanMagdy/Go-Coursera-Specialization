package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Printf("Enter a float number: \n")
	var user_input float64
	
	_, err := fmt.Scanf("%f", &user_input)
	
	if err != nil{
		fmt.Print(err)
	} else {
		fmt.Println(int(math.Trunc(user_input)))

	}
	
}