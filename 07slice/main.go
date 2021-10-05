package main

import "fmt"

func main() {
	var fruits = []string{}

	fruits = append(fruits, "Apple - 🍎", "Banana - 🍌")
	fruits = append(fruits, "Mango - 🥭", "Grape - 🍇")

	fmt.Printf("fruits: %v\n", fruits)

	index := 2
	fruits = append(fruits[:index], fruits[index+1:]...)

	fmt.Printf("fruits: %v\n", fruits)

}
