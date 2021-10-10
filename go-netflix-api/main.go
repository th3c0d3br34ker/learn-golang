package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/th3c0br34ker/go-netflix-api/router"
)

func main() {
	fmt.Println("MongoDB API")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
