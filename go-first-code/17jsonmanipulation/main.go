package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"username"`
	Age      int
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	encodeJSON()
	decodeJSON()
}

func encodeJSON() {
	users := []User{
		{"Ramesh", 25, "123456", []string{"frontend", "app"}},
		{"Rajesh", 22, "123456", []string{"backend", "app"}},
		{"Kalpesh", 26, "123456", nil},
	}

	// package as JSON
	usersJson, err := json.MarshalIndent(users, "", "    ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(usersJson))
}

func decodeJSON() {
	jsonData := []byte(`
		{
			"username": "Ramesh",
			"Age": 25,
			"tags": ["frontend","app"]
		}
	`)

	var userData User

	isValid := json.Valid(jsonData)

	if isValid {
		json.Unmarshal(jsonData, &userData)

		// fmt.Printf("%#v\n", userData)
	} else {
		fmt.Println("Invalid JSON!")
	}

	var ramesh map[string]interface{}

	json.Unmarshal(jsonData, &ramesh)
	fmt.Printf("%#v\n", ramesh)

	for key, value := range ramesh {
		fmt.Printf(" Key: %v\n Value: %v\n Type: %T\n\n", key, value, value)
	}
}
