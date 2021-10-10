package main

import "fmt"

func main() {
	var num int = 5
	var ptr *int = &num

	fmt.Printf("Pointer: %p \nValue: %d", ptr, *ptr)
}
