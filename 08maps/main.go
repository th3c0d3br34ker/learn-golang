package main

import "fmt"

func main() {
	fruits := make(map[string]string, 5)

	fruits["🍎"] = "Apple"
	fruits["🍇"] = "Grape"
	fruits["🥭"] = "Mango"
	fruits["🍌"] = "Banana"

	fmt.Printf("fruits: %v\n", fruits)

	fmt.Printf("fruits[\"🥭\"]: %v\n", fruits["🥭"])
}
