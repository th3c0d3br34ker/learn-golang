package main

import "fmt"

func main() {

	result := proAdder(1, 2, 3, 4, 5)

	fmt.Printf("result: %v\n", result)

	jainam := User{"Jainam", true, 20}

	fmt.Printf("jainam: %+v\n", jainam)
	jainam.getStatus()
	jainam.NewAge(22)
	fmt.Printf("jainam: %+v\n", jainam)

}

func proAdder(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}

	return total
}

type User struct {
	Name   string
	Online bool
	Age    int
}

func (user User) getStatus() {
	fmt.Println("Is user online:", user.Online)
}

func (user User) NewAge(age int) {
	user.Age = age
	fmt.Printf("user.Age: %v\n", user.Age)
}
