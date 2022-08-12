package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

/*
error handling

os.open school pdf

if err != nil
return N/A to webserver
*/

func Course(input string) {

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
}
