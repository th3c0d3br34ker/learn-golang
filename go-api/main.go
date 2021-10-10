package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice float64 `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// create DB
var Courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == "" || c.CoursePrice == 0
}

func main() {
	fmt.Println("GO API")
	r := mux.NewRouter()

	//seeding
	Courses = append(Courses, Course{
		CourseId:    "2",
		CourseName:  "ReactJS",
		CoursePrice: 299,
		Author: &Author{
			FullName: "Ramesh Choudhary",
			Website:  "lco.dev",
		},
	})
	Courses = append(Courses, Course{
		CourseId:    "4",
		CourseName:  "MERN Stack",
		CoursePrice: 199,
		Author: &Author{
			FullName: "Ramesh Choudhary",
			Website:  "go.dev",
		},
	})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/course", createPost).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourseById).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourseById).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controller - file
func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome!</h1>"))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Course by Id")

	w.Header().Set("Content-Type", "application/json")

	// get id from request
	params := mux.Vars(r)

	// GET the Course from DB
	for _, course := range Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course Found")
}

func createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create a Course")

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid Body")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Invalid Data in Body")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	Courses = append(Courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course by ID")
	w.Header().Set("Content-Type", "applicatioan/json")

	// first - grab id from req
	params := mux.Vars(r)

	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			Courses = append(Courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Invalid Course ID")
}

func deleteCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete  Course by ID")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted")
			return
		}
	}

	json.NewEncoder(w).Encode("Invalid Course ID")
}
