package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("golang time package")

	now := time.Now()

	fmt.Printf(now.Format("01-02-2006 Monday"))

}
