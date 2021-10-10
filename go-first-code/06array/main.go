package main

import "fmt"

func main() {
	var fruitsList [4]string

	fruitsList[0] = "Apple - 🍎"
	fruitsList[1] = "Banana - 🍌"
	fruitsList[2] = "Mango - 🥭"
	fruitsList[3] = "Grape - 🍇"

	fmt.Printf("fruitsList: %v\n", fruitsList)
}
