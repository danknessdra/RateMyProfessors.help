package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
error handling

os.open school json

if err != nil
return N/A to webserver
*/

// return ranking of courses based off school, and course name

// first return courses that contain course name

type Course struct {
	Course     string  `json:"course"`
	Department string  `json:"department"`
	Prof       string  `json:"prof"`
	Rating     float64 `json:"rating"`
}

type Courses struct {
	Courses []Course `json:"courses"`
}

type ByRanking []Course

func (a ByRanking) Len() int           { return len(a) }
func (a ByRanking) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRanking) Less(i, j int) bool { return a[i].Rating < a[j].Rating }

func GetCourse(input string) {

	jsonFile, err := os.Open("../data/courses.json")

	if err != nil {
		fmt.Println(err)
	}

	// research what ioutil.ReadALL does
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// courses var that is the struct Courses( which is a list of course structs )
	var courses Courses

	//var courses map[string]Course

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'course' which we defined above
	json.Unmarshal(byteValue, &courses)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example

	count := 0
	sort.Sort(ByRanking(courses.Courses))

	for i := 0; i < len(courses.Courses); i++ {

		count += 1
		//fmt.Println(count)

		if strings.Compare(courses.Courses[i].Course, "MATH 1A") == 0 {
			fmt.Println("Department: " + courses.Courses[i].Department)
			fmt.Println("Course: " + courses.Courses[i].Course)
			fmt.Println("Professor: " + courses.Courses[i].Prof)
			fmt.Println("Rating: " + strconv.FormatFloat(courses.Courses[i].Rating, 'f', 1, 32))
		}

		// kinda/kinda dont fully understand format float, but hopefully it means no exp, and truncate to only 1 place

	}

}
