package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "html/index.html")
}

func main() {

	jsonFile, err := os.Open("../data/courses.json")

	if err != nil {
		fmt.Println(err)
	}

	type Course struct {
		Course     string  `json:"course"`
		Department string  `json:"department"`
		Prof       string  `json:"prof"`
		Rating     float64 `json:"rating"`
	}

	type Courses struct {
		Courses []Course `json:"courses"`
	}

	// research what ioutil.ReadALL does
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// courses var that is the struct Courses( which is a list of course structs )
	var courses Courses

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'course' which we defined above
	json.Unmarshal(byteValue, &courses)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example

	count := 0
	for i := 0; i < len(courses.Courses); i++ {
		count += 1
		fmt.Println(count)
		fmt.Println("Department: " + courses.Courses[i].Department)
		fmt.Println("Course: " + courses.Courses[i].Course)
		fmt.Println("Professor: " + courses.Courses[i].Prof)

		// kinda/kinda dont fully understand format float, but hopefully it means no exp, and truncate to only 1 place
		fmt.Println("Rating: " + strconv.FormatFloat(courses.Courses[i].Rating, 'f', 1, 32))
	}

	// returns distance, [src, destination]
	best, err := Graph().Longest(0, 2)
	fmt.Println(best)
	if err != nil {
		log.Fatal(err)
	}

	router := httprouter.New()
	// Index needs to be a handler, idk if ServeFile is a handler
	router.GET("/", Index)
	//router.GET("/hello/:name", Hello)
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/css/*filepath", http.Dir("css"))

	log.Fatal(http.ListenAndServe(":8080", router))
}

/*
type node struct {
	Course    string  `json:"course"`
	Professor string  `json:"prof"`
	Rating    float32 `json:"rating"`
}

var nodes = []node{
	{Course : "Math", Professor : "jeff", Rating : "3.2"},
	{Course : "Physics", Professor : "saul", Rating : "3.4"},
	{Course : "English", Professor : "white", Rating : "4.6"},
	{Course : "Chemistry", Professor : "jesse", Rating : "4.2"},
	{Course : "History", Professor : "raymond", Rating : "5"},
	{Course : "PE", Professor : "bhat", Rating : "0.0"},
}

*/

/* generates sample nodes
for i in range(0,6):
	courses = ["Math", "Physics", "English", "Chemistry", "History", "PE" ]
	professors = ["jeff", "saul", "white", "jesse", "raymond", "bhat"]
	rating = [3.2, 3.4, 4.6, 4.2, 5, 0.0]

	print("\t{{Course : \"{0}\", Professor : \"{1}\", Rating : \"{2}\"}},".format(courses[i], professors[i], rating[i]))
*/
