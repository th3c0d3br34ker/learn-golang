package main

import (
	"fmt"
	"net/url"
)

const URL = "https://httpbin.org?=name=jainam&age=22"

func main() {

	result, _ := url.Parse(URL)

	// fmt.Printf("result: %+v\n", result.Scheme)

	queryParams := result.Query()

	for _, val := range queryParams {
		fmt.Println(val)
	}

}
