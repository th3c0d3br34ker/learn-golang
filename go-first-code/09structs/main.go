package main

import "fmt"

func main() {
	jainam := User{"Jainam", true, 20}

	fmt.Printf("jainam: %+v\n", jainam)

}

type User struct {
	Name   string
	Online bool
	Age    int
}
