package main

import "fmt"

const SomeInfo string = "some common info"

func main() {
	var username string = "jainamd"
	fmt.Printf("%s: %T", username, username)

	fmt.Println()

	// no var
	age := 16
	fmt.Printf("%d: %T", age, age)

	fmt.Println()

	// pulbic variable
	fmt.Printf("%s: %T", SomeInfo, SomeInfo)

	fmt.Println()

}
