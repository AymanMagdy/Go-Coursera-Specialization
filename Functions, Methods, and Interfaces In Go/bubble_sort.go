package main

import "fmt"

func bubble_sort(elements []int) []int {
	// The buble sort algorithm
	end := len(elements) - 1

	for {
		if end == 0 {
			break
		}

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}
		end -= 1
	}

	return elements
}

func main(){
	slice_numbers := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	sorted_slice := bubble_sort(slice_numbers)

	fmt.Println(sorted_slice)
}