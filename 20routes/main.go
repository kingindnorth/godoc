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

// model for course file
type course struct {
	courseId    string
	couseName   string
	coursePrize int
	author      *author
}

type author struct {
	fullname string
	lastname string
}

// fake db
var courses []course

func main() {
	r := mux.NewRouter()
	//seeding
	courses = append(courses, course{
		courseId: "1", couseName: "reactjs", coursePrize: 299, author: &author{
			fullname: "prajjawal",
			lastname: "tiwari",
		}})
	courses = append(courses, course{
		courseId: "2", couseName: "nodejs", coursePrize: 299, author: &author{
			fullname: "uday",
			lastname: "tiwari",
		}})
	courses = append(courses, course{
		courseId: "3", couseName: "golang", coursePrize: 299, author: &author{
			fullname: "brad",
			lastname: "traversy",
		}})
	courses = append(courses, course{
		courseId: "4", couseName: "javaScript", coursePrize: 299, author: &author{
			fullname: "hitesh",
			lastname: "choudhary",
		}})
	fmt.Println(courses)
	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/couser/{id}", deleteOneCourse).Methods("DELETE")

	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// ------------------------------------ignore-----------------------------------//
// apis
// helper
func (c *course) isEmpty() bool {
	return c.courseId == "" && c.couseName == ""
}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	//grab id from request
	params := mux.Vars(r)
	//loop through courses, find id and send rrequired result
	for _, course := range courses {
		if course.courseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
		json.NewEncoder(w).Encode("no course found with this id")
	}
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	//what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
		return
	}
	//what about : {}
	var course course
	json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("no data")
		return
	}

	//generate unique id and convert the id to string
	rand.Seed(time.Now().UnixNano())
	course.courseId = strconv.Itoa(rand.Intn(100)) //converted to string and stored in courseId
	//append new course to courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	//grab id from request
	params := mux.Vars(r)
	//loop to find the id in the courses
	for index, val := range courses {
		if val.courseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course course
			json.NewDecoder(r.Body).Decode(&course)
			course.courseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("course not found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	//grab id from request
	params := mux.Vars(r)
	//loop to find the id in the courses
	for index, val := range courses {
		if val.courseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course deleted")
			return
		}
	}
	json.NewEncoder(w).Encode("course not found")
}
