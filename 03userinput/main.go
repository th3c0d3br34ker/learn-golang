package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a rating")
	input, _, err := reader.ReadLine()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	numberRating, err := strconv.ParseInt(strings.TrimSpace(string(input)), 10, 64)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	numberRating = numberRating + 1

	fmt.Printf("Rating: %d", numberRating)
}
