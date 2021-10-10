package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fruits := []string{"Apple - 🍎", "Banana - 🍌", "Mango - 🥭", "Grape - 🍇"}

	file, err := os.Create("./out.txt")

	checkNilErr(err)

	length, err := io.WriteString(file, strings.Join(fruits, "\n"))

	checkNilErr(err)

	fmt.Printf("length: %v\n", length)

	defer file.Close()

	readFile("out.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Printf("databyte: %s\n", databyte)
}

func checkNilErr(err error) {

	if err != nil {
		panic(err)
	}
}
