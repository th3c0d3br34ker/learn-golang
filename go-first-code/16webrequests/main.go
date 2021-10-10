package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const API_URI = "http://localhost:3000"

func main() {
	GetRequest()
	PostRequest()
	PostFormRequest()
}

func GetRequest() {
	const URL = API_URI + "/get"
	res, err := http.Get(URL)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Printf("Status code:  %d\n", res.StatusCode)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PostRequest() {
	const URL string = API_URI + "/post"

	// payload
	requestBody := strings.NewReader(`
			{
				"country": "India",
				"city": "Kolkata",
				"pincode": "700015"
			}
		`)

	res, err := http.Post(URL, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Printf("Status code:  %d\n", res.StatusCode)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}

func PostFormRequest() {
	const uri = API_URI + "/postform"

	// payload
	data := url.Values{}
	data.Add("country", "India")
	data.Add("city", "Kolkata")
	data.Add("pincode", "700015")

	res, err := http.PostForm(uri, data)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Printf("Status code:  %d\n", res.StatusCode)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
