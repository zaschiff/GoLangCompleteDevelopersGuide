package main

import "fmt"

func main() {
	var numbers []int

	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
	}

	for _, value := range numbers {
		if value%2 == 0 {
			fmt.Println(value, "is even")
		} else {
			fmt.Println(value, "is odd")
		}
	}
}
