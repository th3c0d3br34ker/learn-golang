package main

import "fmt"

func main() {

	fruits := []string{"Apple - 🍎", "Banana - 🍌", "Mango - 🥭", "Grape - 🍇"}

	for _, fruit := range fruits {

		if fruit == "Apple - 🍎" {
			goto eatApple
		}

		fmt.Printf("fruit: %v\n", fruit)
	}

eatApple:
	fmt.Println("Eating Apple!")
}
