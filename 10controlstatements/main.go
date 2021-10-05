package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ifing()
	swtiching()
}

func ifing() {

	loginCount := 23
	result := ""

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount < 20 {
		result = "Premium User"
	} else {
		result = "God User"
	}

	fmt.Printf("result: %v\n", result)

	if num := 3; num < 10 {
		fmt.Println("Less than 3.")
	} else {
		fmt.Println("More than 3.")
	}
}

func swtiching() {

	rand.Seed(int64(time.Now().Nanosecond()))
	diceNumber := rand.Intn(6) + 1

	fmt.Printf("diceNumber: %v\n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	case 3:
		fmt.Println("Value is 3")
	case 4:
		fmt.Println("Value is 4")
	case 5:
		fmt.Println("Value is 5")
	case 6:
		fmt.Println("Value is 6")

	}
}
